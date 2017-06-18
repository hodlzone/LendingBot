
app.controller('coinController', ['$scope', '$http', '$log', '$timeout','$routeParams',
	function($scope, $http, $log, $timeout, $routeParams) {
		var coinScope = $scope;
		var earnedFeeChart;

		coinScope.resetPoloniexKeys = function() {
			coinScope.poloniexKey = coinScope.poloniexKeyOrig;
			coinScope.poloniexSecret = coinScope.poloniexSecretOrig;
		}

		coinScope.getLendingHistory = function() {
			$http(
			{
				method: 'GET',
				url: '/dashboard/data/lendinghistorysummary', //' + coinScope.coin
				data : {
				},
				withCredentials: true
			})
			.then((res) => {
				//success
				coinScope.hasCompleteLoans = res.data.LoanHistory ? true : false;
				if (true || coinScope.hasCompleteLoans) {
					earnedFeeChart = echarts.init(document.getElementById('lendingHistoryChart')),
					earned = [],
					fee = [],
					dates = [];
					d = -1
					// res.data.LoanHistory
					for(i = 0; i < res.data.LoanHistory.length; i++) {
						fee.push(parseFloat(res.data.LoanHistory[i].data[coinScope.coin].fees));
						earned.push(parseFloat(res.data.LoanHistory[i].data[coinScope.coin].earned));
					 	dates.push(new Date(res.data.LoanHistory[i].time));
					}
					// for(i = 0; i < res.data.CompleteLoans.length; i++) {
					// 	if (i > 0 && d >= 0 && new Date(res.data.CompleteLoans[i].close).getDate() == new Date(dates[d]).getDate()) {
					// 		fee[d] += parseFloat(res.data.CompleteLoans[i].fee);
					// 		earned[d] += parseFloat(res.data.CompleteLoans[i].earned);
					// 	} else {
					// 		fee.push(parseFloat(res.data.CompleteLoans[i].fee));
					// 		earned.push(parseFloat(res.data.CompleteLoans[i].earned));
					// 		dates.push(new Date(res.data.CompleteLoans[i].close));
					// 		d++
					// 	}
					// }
					var option = {
						dataZoom : {
							show : true,
							realtime: true,
							start : 50,
							end : 100
						},
						tooltip : {
							trigger: 'axis',
							axisPointer : {
								type : 'shadow'
							},
							formatter: function (params){
								return "Date: " + params[0].name + '<br/>'
								+ params[0].seriesName + ' : ' + params[0].value + '<br/>'
								+ params[1].seriesName + ' : ' + (params[1].value + params[0].value);
							}
						},
						legend: {
							selectedMode:false,
							data:['Earned', 'Fee']
						},
						toolbox: {
							show : true,
							feature : {
								mark : {show: true},
								dataView : {show: true, readOnly: false},
								restore : {show: true},
								saveAsImage : {show: true}
							}
						},
						calculable : true,
						xAxis : [
						{
							type : 'category',
							data : dates,
						}
						],
						yAxis : [
						{
							type : 'value',
							boundaryGap: [0, 0.1]
						}
						],
						series : [
						{
							name:'Earned',
							type:'bar',
							stack: 'sum',
							barCategoryGap: '50%',
							itemStyle: {
								normal: {
									color: '#46D246',
									barBorderColor: '#46D246',
									barBorderWidth: 6,
									barBorderRadius:0,
									label : {
										show: true, position: 'top'
									}
								}
							},
							data: earned,
						},
						{
							name:'Fee',
							type:'bar',
							stack: 'sum',
							itemStyle: {
								normal: {
									color: '#ff4c4c',
									barBorderColor: '#ff4c4c',
									barBorderWidth: 6,
									barBorderRadius:0,
									label : {
										show: false, 
										position: 'top',
										formatter: function (params) {
											for (var i = 0, l = option.xAxis[0].data.length; i < l; i++) {
												if (option.xAxis[0].data[i] == params.name) {
													return option.series[0].data[i] + params.value;
												}
											}
										},
										textStyle: {
											color: '#ff4c4c'
										}
									}
								}
							},
							data: fee,
						}
						]
					}
					earnedFeeChart.setOption(option);
					$timeout(() => {
						earnedFeeChart.resize();
					});
				}
			}, (err) => {
				//error
				$log.error("LendingHistory: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
			});
		}

		coinScope.getDetailedLendingHistory = function() {
			$http(
			{
				method: 'GET',
				url: '/dashboard/data/detstats',
				data : {
				},
				withCredentials: true
			})
			.then((res) => {
				//success
				console.log(res.data.data);
			}, (err) => {
				//error
				$log.error("LendingHistory: Error: [" + JSON.stringify(err) + "] Status [" + err.status + "]");
			});
		}


		// /Coin
		coinScope.coin = $routeParams.coin;
		coinScope.getLendingHistory();
		coinScope.getDetailedLendingHistory();
		//resize charts
		window.onresize = function() {
			if (coinScope.hasCompleteLoans) {
				$timeout(() => {
					earnedFeeChart.resize();
				});
			}
		};
		//------

	}]);

