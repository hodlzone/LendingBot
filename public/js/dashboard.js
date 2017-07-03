var app=angular.module("lendingApp",["ngRoute","ngMask", "ngCookies"]);

app.factory('redirectInterceptor', function($q,$location,$window){
    return  {
        response:function(response) {
        	console.log("RESPONSE: " + response);
        	return response
        },
        responseError:function(response) {
        	console.log("RESPONSE: " + response);
        	if (response.status == 403) {
        		window.location = "/";
        	}
    		return $q.reject(response);
        }
    }
});

app.config(['$routeProvider', '$locationProvider', '$httpProvider',
	function($routeProvider, $locationProvider, $httpProvider) {
		$routeProvider
		.when("/",{
			templateUrl : "/dashboard/info",
			controller : "dashInfoController"
		})
		.when("/info/:coin",{
			templateUrl : "/dashboard/info/:id",
			controller : "dashInfoAdvancedController"
		})
		.when("/settings/user",{
			templateUrl : "/dashboard/settings/user",
			controller : "dashSettingsUserController"
		})
		.when("/settings/lending",{
			templateUrl : "/dashboard/settings/lending",
			controller : "dashSettingsLendingController"
		})
		.when("/logs",{
			templateUrl : "/dashboard/logs",
			controller : "dashLogsController"
		})
		.when("/sysadmin",{
			templateUrl : "/dashboard/sysadmin",
			controller : "sysAdminController"
		})
		.when("/admin",{
			templateUrl : "/dashboard/admin",
			controller : "adminController"
		})
		.when("/admin/queuerstatus",{
			templateUrl : "/dashboard/admin/queuerstatus",
			controller : "adminControllerQuererStatus"
		})
		.when("/admin/logs",{
			templateUrl : "/dashboard/admin/logs",
			controller : "adminControllerLogs"
		})
		.when("/coin/:coin",{
			templateUrl : "/dashboard/coin",
			controller : "coinController"
		})
		.otherwise({redirectTo:'/'});


		$locationProvider.html5Mode({enabled: false, requireBase: false});

		$httpProvider.interceptors.push('redirectInterceptor');
	}]);

app.controller('dashBaseController', ['$scope', '$http', '$log', "$location", "$window", "$rootScope", "$cookies", "$interval",
	function($scope, $http, $log, $location, $window, $rootScope, $cookies, $interval) {
		var dashBaseScope = $scope;

		$rootScope.$on('$locationChangeStart', function (event) {
			if($cookies.get('HODL_TIMEOUT') == null) {
				$window.location = '/'
			}
			console.log("Time: " + $cookies.get("HODL_TIMEOUT"));
		});

		$interval(() => {
			if($cookies.get('HODL_TIMEOUT') == null) {
				$window.location = '/'
			}
		}, 15000)
	}]);

app.controller('dashInfoController', ['$scope', '$http', '$log',
	function($scope, $http, $log) {
		var dashInfoScope = $scope;

		dashInfoScope.getCurrentUserStats = function() {
			$http(
			{
				method: 'GET',
				url: '/dashboard/data/currentuserstats',
				data : {
				},
				withCredentials: true
			})
			.then((res) => {
				//success
				console.log(res.data)
				dashInfoScope.currentUserStats = res.data.CurrentUserStats;
				dashInfoScope.balances = res.data.Balances;
				init_chart_doughnut(dashInfoScope.balances)
			}, (err) => {
				//error
				$log.error("CurrentUserStats: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
			});
		}

		dashInfoScope.getLendingHistory = function() {
			$http(
			{
				method: 'GET',
				url: '/dashboard/data/lendinghistory',
				data : {
				},
				withCredentials: true
			})
			.then((res) => {
				//success
				console.log(res.data);
				$.each(res.data.CompleteLoans, (index, val) => {
					var tr = $("<tr>")
					.append($("<td>").html(val.currency))
					.append($("<td>").html(parseFloat(val.rate).toFixed(4)))
					.append($("<td>").html(val.amount))
					.append($("<td>").html(val.earned))
					.append($("<td>").html(val.fee))
					.append($("<td>").html(val.close));
					// .append($("<td>").html(val.duration));
					$("#lendingHistory").append(tr);
				});
				if (!$.fn.DataTable.isDataTable('#datatable-responsive')) {
					$('#datatable-responsive').DataTable({
						filter: false
					});
				}
			}, (err) => {
				//error
				$log.error("LendingHistory: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
			});
		}

		//init
		dashInfoScope.getCurrentUserStats();
		dashInfoScope.getLendingHistory();
		dashInfoScope.backgroundColor = backgroundColor;
		//----
	}]);

app.controller('dashInfoAdvancedController', ['$scope', '$http', '$log',
	function($scope, $http, $log) {
		var dashInfoAdvScope = $scope;
	}]);

app.controller('dashSettingsUserController', ['$scope', '$http', '$log',
	function($scope, $http, $log) {
		var dashSettingsUserScope = $scope;

		dashSettingsUserScope.create2FA = function() {
			dashSettingsUserScope.loadingCreate2FA = true;
			dashSettingsUserScope.create2FAError = '';
			$http(
			{
				method: 'POST',
				url: '/dashboard/settings/create2fa',
				data : {
					pass: dashSettingsUserScope.pass2FA,
				},
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("2FA: Success.");
				dashSettingsUserScope.qrcode = 'data:image/png;base64,' + res.data.data 
				dashSettingsUserScope.has2FA = true;
			}, (err) => {
				//error
				$log.error("2FA: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsUserScope.create2FAError = err.data.error;
			})
			.then(() => {
				dashSettingsUserScope.pass2FA = '';
			});
			dashSettingsUserScope.loadingCreate2FA = false;
		}

		dashSettingsUserScope.enable2FA = function(bool) {
			dashSettingsUserScope.loadingEnable2FA = true;
			dashSettingsUserScope.enable2FAError = '';
			dashSettingsUserScope.enable2FASuccess = '';
			$http(
			{
				method: 'POST',
				url: '/dashboard/settings/enable2fa',
				data : {
					pass: dashSettingsUserScope.pass2FA,
					token: dashSettingsUserScope.token,
					enable: bool,
				},
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("2FA: Success.");
				dashSettingsUserScope.enabled2FA = (res.data.data === 'true')
				if (dashSettingsUserScope.enabled2FA) {
					dashSettingsUserScope.enable2FASuccess = 'Success! 2FA is enabled.'
				} else {
					dashSettingsUserScope.enable2FASuccess = 'Success! 2FA is disabled.'
				}
			}, (err) => {
				//error
				$log.error("2FA: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsUserScope.enable2FAError = err.data.error;
			})
			.then(() => {
				dashSettingsUserScope.pass2FA = '';
				dashSettingsUserScope.token = '';
				dashSettingsUserScope.loadingEnable2FA = false;
			});
		}

		dashSettingsUserScope.verifyEmail = function() {
			dashSettingsUserScope.loadingVerified = false;
			dashSettingsUserScope.verifiedSuccess = '';
			dashSettingsUserScope.verifiedError = '';
			$http(
			{
				method: 'GET',
				url: '/verify/request',
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("VerifyEmail: Success.");
				dashSettingsUserScope.verifiedSuccess = 'Verification email sent!';
			}, (err) => {
				//error
				$log.error("VerifyEmail: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsUserScope.verifiedError = err.data.error;
			})
			.then(() => {
				dashSettingsUserScope.loadingVerified = false;
			})
		}

		dashSettingsUserScope.changePass = function() {
			dashSettingsUserScope.changePassSuccess = '';
			dashSettingsUserScope.changePassError = '';
			$http(
			{
				method: 'POST',
				url: '/dashboard/settings/changepass',
				data : $.param({
					pass: dashSettingsUserScope.pass,
					passnew: dashSettingsUserScope.passNew,
				}),
				headers: {'Content-Type': 'application/x-www-form-urlencoded'},
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("ChangePass: Success.");
				dashSettingsUserScope.changePassSuccess = 'Password changed successfully!';
			}, (err) => {
				//error
				$log.error("ChangePass: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsUserScope.changePassError = err.data.error;
			})
			.then(() => {
				dashSettingsUserScope.pass = '';
				dashSettingsUserScope.passNew = '';
				dashSettingsUserScope.passNew2 = '';
			})
		}

		//init
		dashSettingsUserScope.loadingEnable2FA = false;
		dashSettingsUserScope.loadingCreate2FA = false;
		dashSettingsUserScope.loadingVerified = false;

		dashSettingsUserScope.create2FAError = '';
		dashSettingsUserScope.enable2FAError = '';
		dashSettingsUserScope.verifiedError = '';
		dashSettingsUserScope.changePassError = '';

		dashSettingsUserScope.enable2FASuccess = '';
		dashSettingsUserScope.verifiedSuccess = '';
		dashSettingsUserScope.changePassSuccess = '';

		dashSettingsUserScope.pass2FA = '';
		dashSettingsUserScope.token = '';
		dashSettingsUserScope.pass = '';
		dashSettingsUserScope.passNew = '';
		dashSettingsUserScope.passNew2 = '';
		//----
	}]);

app.controller('dashSettingsLendingController', ['$scope', '$http', '$log', '$timeout',
	function($scope, $http, $log, $timeout) {
		var dashSettingsLendingScope = $scope;

		dashSettingsLendingScope.resetExchangeKeys = function() {
			dashSettingsLendingScope.exchangeKey = dashSettingsLendingScope.exchangeKeyOrig;
			dashSettingsLendingScope.exchangeSecret = dashSettingsLendingScope.exchangeSecretOrig;
		}

		dashSettingsLendingScope.getEnableExchangeLending = function() {
			dashSettingsLendingScope.loadingEnableExchangeLending = true;
			$http(
			{
				method: 'GET',
				url: '/dashboard/settings/enableuserlending',
				data : $.param({
				}),
				headers: {'Content-Type': 'application/x-www-form-urlencoded'},
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("getEnablePoloniexLending: Success.");
				dashSettingsLendingScope.coinsEnabled = res.data.data.enable;
				dashSettingsLendingScope.coinsMinLend = res.data.data.min;
				
				$timeout(()=>{
					dashSettingsLendingScope.init_switch();
				})
			}, (err) => {
				//error
				$log.error("getEnablePoloniexLending: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsLendingScope.exchangeKeysEnabledError = 'Unable to load poloniex lending information. Error: ' + err.data.error;
			})
			.then(() => {
				dashSettingsLendingScope.loadingEnablePoloniexLending = false;
			});
		}


		dashSettingsLendingScope.setEnableExchangeLending = function() {
			dashSettingsLendingScope.loadingEnableExchangeLending = true;
			dashSettingsLendingScope.exchangeKeysEnabledError = '';
			dashSettingsLendingScope.exchangeKeysEnableSuccess = '';
			$http(
			{
				method: 'POST',
				url: '/dashboard/settings/enableuserlending',
				data : $.param({
					enable: JSON.stringify(dashSettingsLendingScope.coinsEnabled),
					min: JSON.stringify(dashSettingsLendingScope.coinsMinLend),
				}),
				headers: {'Content-Type': 'application/x-www-form-urlencoded'},
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("setEnablePoloniexLending: Success.");
				dashSettingsLendingScope.exchangeKeysEnableSuccess = 'Poloniex Lending successfully updated values.'
			}, (err) => {
				//error
				$log.error("setEnablePoloniexLending: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsLendingScope.exchangeKeysEnabledError = 'Unable to update poloniex lending information. Error: ' + err.data.error;
			})
			.then(() => {
				dashSettingsLendingScope.loadingEnableExchangeLending = false;
			});
		}

		dashSettingsLendingScope.setExchangeKeys = function() {
			dashSettingsLendingScope.loadingExchangeKeys = true;
			dashSettingsLendingScope.exchangeKeysSetError = '';
			dashSettingsLendingScope.exchangeKeysSetSuccess = '';
			$http(
			{
				method: 'POST',
				url: '/dashboard/settings/setexchangekeys',
				data : $.param({
					exchangekey: dashSettingsLendingScope.exchangeKey,
					exchangesecret: dashSettingsLendingScope.exchangeSecret,
				}),
				headers: {'Content-Type': 'application/x-www-form-urlencoded'},
				withCredentials: true
			})
			.then((res) => {
				//success
				$log.info("SetPoloniexKeys: Success.");
				var tempData = JSON.parse(res.data.data);
				dashSettingsLendingScope.exchangeKeyOrig = tempData.exchangekey;
				dashSettingsLendingScope.exchangeSecretOrig = tempData.exchangesecret;
				//resets to new originals
				dashSettingsLendingScope.resetExchangeKeys();
				dashSettingsLendingScope.exchangeKeysSetSuccess = 'Successfully set poloniex keys.';
			}, (err) => {
				//error
				$log.error("SetPoloniexKeys: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
				dashSettingsLendingScope.exchangeKeysSetError =  err.data.error;
			})
			.then(() => {
				dashSettingsLendingScope.loadingExchangeKeys = false;
			});
		}

		// Switchery
		dashSettingsLendingScope.init_switch = function() {
			if ($(".js-switch")[0]) {
				var elems = Array.prototype.slice.call(document.querySelectorAll('.js-switch'));
				elems.forEach(function (html) {
                // if ($(el).data('switchery') != true) {
                	var switchery = new Switchery(html, {
                		color: '#26B99A'
                	});
                	html.onchange = function(e) {
                		dashSettingsLendingScope.$apply(() => {
                			var me = $(this);
                			dashSettingsLendingScope.coinsEnabled[me.attr('id')] = me.is(':checked');
                		});
                	}
                });
			}
		}
		// /Switchery

		//init
		// init_InputMask();

		dashSettingsLendingScope.loadingExchangeKeys = false;
		dashSettingsLendingScope.loadingEnableExchangeLending = false;

		dashSettingsLendingScope.exchangeKeysEnabledError = '';
		dashSettingsLendingScope.exchangeKeysSetError = '';

		dashSettingsLendingScope.exchangeKeysEnableSuccess = ''
		dashSettingsLendingScope.exchangeKeysSetSuccess = '';

		dashSettingsLendingScope.parseInt = parseInt;
		dashSettingsLendingScope.getEnableExchangeLending();
		//------

	}]);

app.controller('dashLogsController', ['$scope', '$http', '$log',
	function($scope, $http, $log) {
		var dashLogsScope = $scope;
	}]);

function init_chart_doughnut(balanceData){			
	if( typeof (Chart) === 'undefined'){ return; }

	console.log('init_chart_doughnut');
	Object.keys(balanceData.currencymap).forEach(function(key) {
	    value = balanceData.currencymap[key];
	    balanceData.currencymap[key] = value.toFixed(3)
	});


	if ($('.canvasDoughnut').length){
		var chart_doughnut_settings = {
			type: 'doughnut',
			tooltipFillColor: "rgba(51, 51, 51, 0.55)",
			data: {
				labels: Object.keys(balanceData.percentmap),
				datasets: [{
					data: Object.values(balanceData.currencymap),
					backgroundColor: backgroundColor,
				}]
			},
			options: { 
				legend: false, 
				responsive: false 
			}
		}
		
		$('.canvasDoughnut').each(function(){

			var chart_element = $(this);
			var chart_doughnut = new Chart( chart_element, chart_doughnut_settings);

		});			
		
	}  
}

function init_InputMask() {

	if( typeof ($.fn.inputmask) === 'undefined'){ return; }
	console.log('init_InputMask');

	$(":input").inputmask();

};

var backgroundColor =[
"#00BFFF",
"#FF69B4",
"#7CFC00",
"#800000",
"#FFA500",
"#FF4500",
"#800080",
"#00FF7F",
"#FFFF00",
"#9ACD32",
"#FF6347"
]
