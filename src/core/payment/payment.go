package payment

import (
	"fmt"
	"sync"
	"time"

	"github.com/Emyrk/LendingBot/src/core/database/mongo"
	"github.com/Emyrk/LendingBot/src/core/userdb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PaymentDatabase struct {
	db *mongo.MongoDB

	//mux for generating code
	referralMux sync.Mutex
}

func NewPaymentDatabase(uri, dbu, dbp string) (*PaymentDatabase, error) {
	db, err := mongo.CreatePaymentDB(uri, dbu, dbp)
	if err != nil {
		return nil, fmt.Errorf("Error creating payment db: %s\n", err.Error())
	}
	return &PaymentDatabase{db: db}, err
}

func NewPaymentDatabaseGiven(db *mongo.MongoDB) *PaymentDatabase {
	return &PaymentDatabase{db: db}
}

func (p *PaymentDatabase) Close() error {
	// if p.db == nil {
	// 	return p.db.Close()
	// }
	return nil
}

type Status struct {
	Username              string  `json:"email" bson:"_id"`
	TotalDebt             float64 `json:"tdebt" bson:"tdebt"`
	UnspentCredits        float64 `json:"unspentcred" bson:"unspentcred"`
	SpentCredits          float64 `json:"spentcred" bson:"spentcred"`
	CustomChargeReduction float64 `json:"customchargereduc" bson:"customchargereduc"`
	RefereeCode           string  `json:"refereecode" bson:"refereecode"` //(Person code who referred you)
	RefereeTime           string  `json:"refereetime" bson:"refereetime"`
	ReferralCode          string  `json:"referralcode" bson:"referralcode"`
}

func (p *PaymentDatabase) SetUserReferee(username, refereeCode string) error {
	s, c, err := p.db.GetCollection(mongo.C_Status)
	if err != nil {
		return fmt.Errorf("SetUserReferee: createSession: %s", err.Error())
	}
	defer s.Close()

	st, err := p.getStatusRefereeGiven(refereeCode, c)
	if err != nil {
		//referee code does not exist
		return fmt.Errorf("SetUserReferee: getref: %s", err.Error())
	}

	st, err = p.getStatusGiven(username, c)
	if err != nil {
		return fmt.Errorf("SetUserReferee: getRef: %s", err.Error())
	}

	if st.RefereeCode != "" {
		return fmt.Errorf("Referee already set for user[%s]", username)
	} else if st.ReferralCode == refereeCode {
		return fmt.Errorf("Referee code[%s] is same as users[%s]", st.ReferralCode, refereeCode)
	}
	st.RefereeCode = refereeCode

	//CAN OPTIMIZE LATER
	upsertKey := bson.M{
		"_id": username,
	}
	upsertAction := bson.M{"$set": st}
	_, err = c.Upsert(upsertKey, upsertAction)
	if err != nil {
		return fmt.Errorf("SetUserReferee: upsert: %s", err)
	}
	return nil
}

func (p *PaymentDatabase) GetUserReferralsIfFound(username string) ([]Status, error) {
	s, c, err := p.db.GetCollection(mongo.C_Status)
	if err != nil {
		var sr []Status
		return sr, fmt.Errorf("AddUserReferral: createSession: %s", err.Error())
	}
	defer s.Close()
	ref, err := p.getUserReferralsGiven(username, c)
	if err != nil && err.Error() == mgo.ErrNotFound.Error() {
		return ref, nil
	} else if err != nil {
		return ref, err
	}
	return ref, nil
}

func (p *PaymentDatabase) GetUserReferrals(username string) ([]Status, error) {
	s, c, err := p.db.GetCollection(mongo.C_Status)
	if err != nil {
		var sr []Status
		return sr, fmt.Errorf("AddUserReferral: createSession: %s", err.Error())
	}
	defer s.Close()
	return p.getUserReferralsGiven(username, c)
}

func (p *PaymentDatabase) getUserReferralsGiven(username string, c *mgo.Collection) ([]Status, error) {
	var result []Status
	find := bson.M{"_id": username}
	//CAN OPTIMIZE to use less data
	err := c.Find(find).All(&result)
	return result, err
}

func (p *PaymentDatabase) GetStatus(username string) (*Status, error) {
	s, c, err := p.db.GetCollection(mongo.C_Status)
	if err != nil {
		return nil, fmt.Errorf("GetStatus: getcol: %s", err)
	}
	defer s.Close()
	var result Status
	err = c.Find(bson.M{"_id": username}).One(&result)
	if err != nil {
		return nil, fmt.Errorf("getStatusGiven: one: %s", err.Error())
	}
	return &result, nil
}

func (p *PaymentDatabase) SetStatus(status Status) error {
	s, c, err := p.db.GetCollection(mongo.C_Status)
	if err != nil {
		return fmt.Errorf("SetStatus: getcol: %s", err)
	}
	defer s.Close()

	_, err = c.UpsertId(status.Username, bson.M{"$set": status})
	if err != nil {
		return fmt.Errorf("SetStatus: upsert: %s", err)
	}
	return nil
}

func (p *PaymentDatabase) ReferralCodeExists(refereeCode string) bool {
	s, c, err := p.db.GetCollection(mongo.C_Status)
	if err != nil {
		return true
	}
	defer s.Close()
	_, err = p.getStatusRefereeGiven(refereeCode, c)
	if err != nil {
		return true
	}
	return false
}

func (p *PaymentDatabase) getStatusGiven(username string, c *mgo.Collection) (*Status, error) {
	var result Status
	err := c.Find(bson.M{"_id": username}).One(&result)
	if err != nil {
		return nil, fmt.Errorf("getStatusGiven: one: %s", err.Error())
	}
	return &result, nil
}

func (p *PaymentDatabase) getStatusRefereeGiven(refereeCode string, c *mgo.Collection) (*Status, error) {
	var result Status
	err := c.Find(bson.M{"referee": refereeCode}).One(&result)
	if err != nil {
		return nil, fmt.Errorf("getStatusRefereeGiven: one: %s", err.Error())
	}
	return &result, nil
}

type Debt struct {
	ID                    *bson.ObjectId      `json:"_id,omitempty" bson:"_id,omitempty"`
	LoanDate              time.Time           `json:"loandate" bson:"loandate"`
	Charge                float64             `json:"charge" bson:"charge"`
	AmountLoaned          float64             `json:"amountloaned" bson:"amountloaned"`
	LoanRate              float64             `json:"loanrate" bson:"loanrate"`
	GrossAmountEarned     float64             `json:"gae" bson:"gae"`
	Currency              string              `json:"cur" bson:"cur"`
	CurrencyToBTC         float64             `json:"curBTC" bson:"curBTC"`
	CurrencyToETH         float64             `json:"curETH" bson:"curETH"`
	Exchange              userdb.UserExchange `json:"exch" bson:"exch"`
	Username              string              `json:"email" bson:"email"`
	FullPaid              bool                `json:"fullpaid" bson:"fullpaid"`
	PaymentPercentageRate float64             `json:"ppr" bson:"ppr"`
}

func (p *PaymentDatabase) SetMultiDebt(debt []Debt) error {
	s, c, err := p.db.GetCollection(mongo.C_Debt)
	if err != nil {
		return fmt.Errorf("SetMultiDebt: getcol: %s", err)
	}
	defer s.Close()

	bulk := c.Bulk()
	for _, o := range debt {
		if o.ID == nil {
			//if nil id assume that this is new record so insert
			bulk.Insert(o)
		} else {
			//upsert to prevent update vs insert error for dups
			bulk.Upsert(
				bson.M{"_id": o.ID},
				bson.M{"$set": o},
			)
		}
	}

	_, err = bulk.Run()
	if err != nil {
		return fmt.Errorf("SetMultiDebt: run: %s", err)
	}
	return nil
}

// PAID
// 0 - Both paid and unpaid
// 1 - paid
// 2 - not paid
func (p *PaymentDatabase) GetAllDebts(username string, paid int) ([]Debt, error) {
	var results []Debt

	s, c, err := p.db.GetCollection(mongo.C_Debt)
	if err != nil {
		return results, fmt.Errorf("GetAllDebts: getcol: %s", err)
	}
	defer s.Close()

	find := bson.M{"_id": username}
	if paid == 1 {
		find["fullpaid"] = true
	} else if paid == 2 {
		find["fullpaid"] = false
	}
	err = c.Find(find).All(&results)
	if err != nil {
		return nil, fmt.Errorf("GetAllDebts: all: %s", err.Error())
	}
	return results, nil
}

type Paid struct {
	ID                 *bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	PaymentDate        time.Time      `json:"paymentdate" bson:"paymentdate"`
	BTCPaid            float64        `json:"btcpaid" bson:"btcpaid"`
	BTCTransactionDate time.Time      `json:"btctrandate" bson:"btctrandate"`
	BTCTransactionID   int64          `json:"btctranid" bson:"btctranid"`
	ETHPaid            float64        `json:"ethpaid" bson:"ethpaid"`
	ETHTransactionDate time.Time      `json:"ethtrandate" bson:"ethtrandate"`
	ETHTransactionID   int64          `json:"ethtranid" bson:"ethtranid"`
	AddressPaidFrom    string         `json:"addr" bson:"addr"`
	Username           string         `json:"email" bson:"email"`
}

func (p *PaymentDatabase) SetMultiPaid(paid []Paid) error {
	s, c, err := p.db.GetCollection(mongo.C_Paid)
	if err != nil {
		return fmt.Errorf("SetMultiPaid: getcol: %s", err)
	}
	defer s.Close()

	bulk := c.Bulk()
	for _, o := range paid {
		if o.ID == nil {
			//if nil id assume that this is new record so insert
			bulk.Insert(o)
		} else {
			//upsert to prevent update vs insert error for dups
			bulk.Upsert(
				bson.M{"_id": o.ID},
				bson.M{"$set": o},
			)
		}
	}

	_, err = bulk.Run()
	if err != nil {
		return fmt.Errorf("SetMultiPaid: run: %s", err)
	}
	return nil
}

// DATE AFTER
// dateAfter
//  - if nil then will return all dates
//  - else will get all paid after date given
func (p *PaymentDatabase) GetAllPaid(username string, dateAfter *time.Time) ([]Paid, error) {
	var results []Paid

	s, c, err := p.db.GetCollection(mongo.C_Paid)
	if err != nil {
		return results, fmt.Errorf("GetAllPaid: getcol: %s", err)
	}
	defer s.Close()

	find := bson.M{"_id": username}
	if dateAfter != nil {
		find["paymentdate"] = bson.M{"$gt": dateAfter}
	}
	err = c.Find(find).All(&results)
	if err != nil {
		return nil, fmt.Errorf("GetAllPaid: all: %s", err.Error())
	}
	return results, nil
}

func (p *PaymentDatabase) GenerateReferralCode(username string) (string, error) {
	//must lock to avoid conflicts
	p.referralMux.Lock()
	defer p.referralMux.Unlock()
	st, err := p.GetStatus(username)
	if err != nil {
		return "", err
	}
	if st.ReferralCode != "" {
		return "", fmt.Errorf("Referral code already set")
	}

	if len(username) < 5 {
		return "", fmt.Errorf("Length is less than 5")
	}
	base := username[0:5]
	i := 0
	for {
		if p.ReferralCodeExists(st.ReferralCode) == false {
			break
		}
		st.ReferralCode = fmt.Sprintf("%s%d", base, i)
		i++
	}
	return st.ReferralCode, p.SetStatus(*st)
}
