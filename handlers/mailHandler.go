package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pravandkatyare/mailing-service/dto"
	"github.com/pravandkatyare/mailing-service/errs"
	"github.com/pravandkatyare/mailing-service/logging"
	"github.com/pravandkatyare/mailing-service/mail"
	sgclient "github.com/pravandkatyare/mailing-service/mailingClients/sendGrid"
	sesclient "github.com/pravandkatyare/mailing-service/mailingClients/ses"
	"github.com/pravandkatyare/mailing-service/utils"
)

const (
	MAIL_JET   = "MJ"
	SEND_GRID  = "SG"
	AMAZON_SES = "SES"
)

// MailHandler sends new mail with each new reqest
func MailHandler(c *gin.Context) {
	var mailDTO dto.MailDTO
	var mail mail.Mail

	if err := c.ShouldBindJSON(&mailDTO); err != nil {
		logging.Errorf("Error binding data to struct: %s", err)
		c.JSON(http.StatusBadRequest, errs.ErrBadRequest)
		return
	}

	switch mailDTO.Client {
	case MAIL_JET:
		logging.Infof("Generating mail template for MailJet")
		mj, err := utils.GetMailjetTemplate(mailDTO.Mail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Something went wrong")
			return
		}
		mailFailures, err := mail.SendMail(mj)
		if err != nil {
			logging.Errorf("Error converting status code: %s", err)
			c.JSON(http.StatusInternalServerError, "Something went wrong")
		}
		if len(mailFailures) == 0 {
			c.JSON(http.StatusOK, "Mail sent to all recipient/s")
		} else {
			c.JSON(http.StatusOK, mailFailures)
		}
		break
	case AMAZON_SES:
		ses := &sesclient.SESClient{}
		mail.SendMail(ses)
		c.JSON(http.StatusOK, "Amazom SES is currently is yet to be implemented as an enhancement")
		break
	case SEND_GRID:
		sg := &sgclient.SGClient{}
		mail.SendMail(sg)
		c.JSON(http.StatusOK, "SendGrid is currently is yet to be implemented as an enhancement")
		break
	}
}
