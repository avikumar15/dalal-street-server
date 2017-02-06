package actions

import (
	"github.com/Sirupsen/logrus"

	"github.com/thakkarparth007/dalal-street-server/session"
	"github.com/thakkarparth007/dalal-street-server/utils"
	"github.com/thakkarparth007/dalal-street-server/models"
	actions_proto "github.com/thakkarparth007/dalal-street-server/socketapi/proto_build/actions"
	errors_proto "github.com/thakkarparth007/dalal-street-server/socketapi/proto_build/errors"
)

var logger *logrus.Entry

func InitActions() {
	logger = utils.Logger.WithFields(logrus.Fields{
		"module": "socketapi.actions",
	})
}

func BuyStocksFromExchange(sess *session.Session, req *actions_proto.BuyStocksFromExchangeRequest) *actions_proto.BuyStocksFromExchangeResponse {

	return nil
}

func CancelAskOrder(sess *session.Session, req *actions_proto.CancelAskOrderRequest) *actions_proto.CancelAskOrderResponse {
	return nil
}

func CancelBidOrder(sess *session.Session, req *actions_proto.CancelBidOrderRequest) *actions_proto.CancelBidOrderResponse {
	return nil
}

func Login(sess *session.Session, req *actions_proto.LoginRequest) *actions_proto.LoginResponse {
	resp := &actions_proto.LoginResponse{}

	email := req.GetEmail()
	password := req.GetPassword()

	user, err := models.Login(email, password)

	switch err {
	case models.UnauthorizedError:
		resp.Response = &actions_proto.LoginResponse_InvalidCredentialsError_{
			InvalidCredentialsError: &actions_proto.LoginResponse_InvalidCredentialsError{
				Reason: "Invalid Credentials",
			},
		}
	case models.NotRegisteredError:
		resp.Response = &actions_proto.LoginResponse_InvalidCredentialsError_{
			InvalidCredentialsError: &actions_proto.LoginResponse_InvalidCredentialsError{
				Reason: "You have not registered for Dalal Street on the Pragyan site",
			},
		}
	case models.InternalError:
		resp.Response = &actions_proto.LoginResponse_InternalServerError{
			InternalServerError: &errors_proto.InternalServerError{
				Reason: "We are facing some issues on the server. Please try again in some time",
			},
		}
	case nil:
		sess.Set("userId", string(user.Id))

		resp.Response = &actions_proto.LoginResponse_Result{
			Result: &actions_proto.LoginResponse_LoginSuccessResponse{
				SessionId: sess.Id,
				User: user.ToProto(),
				// populate other fields :)
			},
		}
	}

	return resp
}

func Logout(ses *session.Session, req *actions_proto.LogoutRequest) *actions_proto.LogoutResponse {
	return nil
}

func MortgageStocks(sess *session.Session, req *actions_proto.MortgageStocksRequest) *actions_proto.MortgageStocksResponse {
	return nil
}

func PlaceAskOrder(sess *session.Session, req *actions_proto.PlaceAskOrderRequest) *actions_proto.PlaceAskOrderResponse {
	return nil
}

func PlaceBidOrder(sess *session.Session, req *actions_proto.PlaceBidOrderRequest) *actions_proto.PlaceBidOrderResponse {
	return nil
}

func RetrieveMortgageStocks(sess *session.Session, req *actions_proto.RetrieveMortgageStocksRequest) *actions_proto.RetrieveMortgageStocksResponse {
	return nil
}

func Unsubscribe(sess *session.Session, req *actions_proto.UnsubscribeRequest) *actions_proto.UnsubscribeResponse {
	return nil
}

func Subscribe(done <-chan struct{}, updates chan interface{}, sess *session.Session, req *actions_proto.SubscribeRequest) *actions_proto.SubscribeResponse {
	return nil
}

func GetCompanyProfile(sess *session.Session, req *actions_proto.GetCompanyProfileRequest) *actions_proto.GetCompanyProfileResponse {
	return nil
}

func GetMarketEvents(sess *session.Session, req *actions_proto.GetMarketEventsRequest) *actions_proto.GetMarketEventsResponse {
	return nil
}

func GetMyAsks(sess *session.Session, req *actions_proto.GetMyAsksRequest) *actions_proto.GetMyAsksResponse {
	return nil
}

func GetMyBids(sess *session.Session, req *actions_proto.GetMyBidsRequest) *actions_proto.GetMyBidsResponse {
	return nil
}

func GetNotifications(sess *session.Session, req *actions_proto.GetNotificationsRequest) *actions_proto.GetNotificationsResponse {
	return nil
}

func GetTransactions(sess *session.Session, req *actions_proto.GetTransactionsRequest) *actions_proto.GetTransactionsResponse {
	return nil
}

func GetMortgageDetails(sess *session.Session, req *actions_proto.GetMortgageDetailsRequest) *actions_proto.GetMortgageDetailsResponse {
	return nil
}

func GetLeaderboard(sess *session.Session, req *actions_proto.GetLeaderboardRequest) *actions_proto.GetLeaderboardResponse {
	return nil
}
