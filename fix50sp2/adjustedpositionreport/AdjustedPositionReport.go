//Package adjustedpositionreport msg type = BL.
package adjustedpositionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/positionqty"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a AdjustedPositionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BL"`
	fixt11.Header
	//PosMaintRptID is a required field for AdjustedPositionReport.
	PosMaintRptID string `fix:"721"`
	//PosReqType is a non-required field for AdjustedPositionReport.
	PosReqType *int `fix:"724"`
	//ClearingBusinessDate is a required field for AdjustedPositionReport.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for AdjustedPositionReport.
	SettlSessID *string `fix:"716"`
	//Parties is a required component for AdjustedPositionReport.
	parties.Parties
	//PositionQty is a required component for AdjustedPositionReport.
	positionqty.PositionQty
	//InstrmtGrp is a non-required component for AdjustedPositionReport.
	InstrmtGrp *instrmtgrp.InstrmtGrp
	//SettlPrice is a non-required field for AdjustedPositionReport.
	SettlPrice *float64 `fix:"730"`
	//PriorSettlPrice is a non-required field for AdjustedPositionReport.
	PriorSettlPrice *float64 `fix:"734"`
	//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
	PosMaintRptRefID *string `fix:"714"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetPosMaintRptID(v string)                { m.PosMaintRptID = v }
func (m *Message) SetPosReqType(v int)                      { m.PosReqType = &v }
func (m *Message) SetClearingBusinessDate(v string)         { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)                  { m.SettlSessID = &v }
func (m *Message) SetParties(v parties.Parties)             { m.Parties = v }
func (m *Message) SetPositionQty(v positionqty.PositionQty) { m.PositionQty = v }
func (m *Message) SetInstrmtGrp(v instrmtgrp.InstrmtGrp)    { m.InstrmtGrp = &v }
func (m *Message) SetSettlPrice(v float64)                  { m.SettlPrice = &v }
func (m *Message) SetPriorSettlPrice(v float64)             { m.PriorSettlPrice = &v }
func (m *Message) SetPosMaintRptRefID(v string)             { m.PosMaintRptRefID = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BL", r
}
