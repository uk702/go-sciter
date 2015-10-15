// generated by stringer -type=BehaviorEvent,MouseEvent,CursorType,KeyEvent,FocusEvent,ScrollEvent,GestureCmd,GestureState,GestureTypeFlag,DrawEvent,EventReason,EditChangedReason,BehaviorMethodIdentifier,SCDOM_RESULT,VALUE_RESULT -output types_string.go; DO NOT EDIT

package sciter

import "fmt"

const (
	_BehaviorEvent_name_0 = "BUTTON_CLICKBUTTON_PRESSBUTTON_STATE_CHANGEDEDIT_VALUE_CHANGINGEDIT_VALUE_CHANGEDSELECT_SELECTION_CHANGEDSELECT_STATE_CHANGEDPOPUP_REQUESTPOPUP_READYPOPUP_DISMISSEDMENU_ITEM_ACTIVEMENU_ITEM_CLICK"
	_BehaviorEvent_name_1 = "CONTEXT_MENU_REQUESTVISIUAL_STATUS_CHANGEDDISABLED_STATUS_CHANGEDPOPUP_DISMISSING"
	_BehaviorEvent_name_2 = "CONTENT_CHANGED"
	_BehaviorEvent_name_3 = "HYPERLINK_CLICK"
	_BehaviorEvent_name_4 = "ELEMENT_COLLAPSEDELEMENT_EXPANDEDACTIVATE_CHILDINIT_DATA_VIEWROWS_DATA_REQUESTUI_STATE_CHANGEDFORM_SUBMITFORM_RESETDOCUMENT_COMPLETEHISTORY_PUSHHISTORY_DROPHISTORY_PRIORHISTORY_NEXTHISTORY_STATE_CHANGEDCLOSE_POPUPREQUEST_TOOLTIPANIMATION"
	_BehaviorEvent_name_5 = "DOCUMENT_CREATEDDOCUMENT_CLOSE_REQUESTDOCUMENT_CLOSEDOCUMENT_READY"
	_BehaviorEvent_name_6 = "VIDEO_INITIALIZEDVIDEO_STARTEDVIDEO_STOPPEDVIDEO_BIND_RQ"
)

var (
	_BehaviorEvent_index_0 = [...]uint8{0, 12, 24, 44, 63, 81, 105, 125, 138, 149, 164, 180, 195}
	_BehaviorEvent_index_1 = [...]uint8{0, 20, 42, 65, 81}
	_BehaviorEvent_index_2 = [...]uint8{0, 15}
	_BehaviorEvent_index_3 = [...]uint8{0, 15}
	_BehaviorEvent_index_4 = [...]uint8{0, 17, 33, 47, 61, 78, 94, 105, 115, 132, 144, 156, 169, 181, 202, 213, 228, 237}
	_BehaviorEvent_index_5 = [...]uint8{0, 16, 38, 52, 66}
	_BehaviorEvent_index_6 = [...]uint8{0, 17, 30, 43, 56}
)

func (i BehaviorEvent) String() string {
	switch {
	case 0 <= i && i <= 11:
		return _BehaviorEvent_name_0[_BehaviorEvent_index_0[i]:_BehaviorEvent_index_0[i+1]]
	case 16 <= i && i <= 19:
		i -= 16
		return _BehaviorEvent_name_1[_BehaviorEvent_index_1[i]:_BehaviorEvent_index_1[i+1]]
	case i == 21:
		return _BehaviorEvent_name_2
	case i == 128:
		return _BehaviorEvent_name_3
	case 144 <= i && i <= 160:
		i -= 144
		return _BehaviorEvent_name_4[_BehaviorEvent_index_4[i]:_BehaviorEvent_index_4[i+1]]
	case 192 <= i && i <= 195:
		i -= 192
		return _BehaviorEvent_name_5[_BehaviorEvent_index_5[i]:_BehaviorEvent_index_5[i+1]]
	case 209 <= i && i <= 212:
		i -= 209
		return _BehaviorEvent_name_6[_BehaviorEvent_index_6[i]:_BehaviorEvent_index_6[i+1]]
	default:
		return fmt.Sprintf("BehaviorEvent(%d)", i)
	}
}

const (
	_MouseEvent_name_0 = "MOUSE_ENTERMOUSE_LEAVEMOUSE_MOVEMOUSE_UPMOUSE_DOWNMOUSE_DCLICKMOUSE_WHEELMOUSE_TICKMOUSE_IDLEDROPDRAG_ENTERDRAG_LEAVEDRAG_REQUEST"
	_MouseEvent_name_1 = "MOUSE_CLICKDRAGGING"
)

var (
	_MouseEvent_index_0 = [...]uint8{0, 11, 22, 32, 40, 50, 62, 73, 83, 93, 97, 107, 117, 129}
	_MouseEvent_index_1 = [...]uint8{0, 11, 19}
)

func (i MouseEvent) String() string {
	switch {
	case 0 <= i && i <= 12:
		return _MouseEvent_name_0[_MouseEvent_index_0[i]:_MouseEvent_index_0[i+1]]
	case 255 <= i && i <= 256:
		i -= 255
		return _MouseEvent_name_1[_MouseEvent_index_1[i]:_MouseEvent_index_1[i+1]]
	default:
		return fmt.Sprintf("MouseEvent(%d)", i)
	}
}

const _CursorType_name = "CURSOR_ARROWCURSOR_IBEAMCURSOR_WAITCURSOR_CROSSCURSOR_UPARROWCURSOR_SIZENWSECURSOR_SIZENESWCURSOR_SIZEWECURSOR_SIZENSCURSOR_SIZEALLCURSOR_NOCURSOR_APPSTARTINGCURSOR_HELPCURSOR_HANDCURSOR_DRAG_MOVECURSOR_DRAG_COPY"

var _CursorType_index = [...]uint8{0, 12, 24, 35, 47, 61, 76, 91, 104, 117, 131, 140, 158, 169, 180, 196, 212}

func (i CursorType) String() string {
	if i+1 >= CursorType(len(_CursorType_index)) {
		return fmt.Sprintf("CursorType(%d)", i)
	}
	return _CursorType_name[_CursorType_index[i]:_CursorType_index[i+1]]
}

const _KeyEvent_name = "KEY_DOWNKEY_UPKEY_CHAR"

var _KeyEvent_index = [...]uint8{0, 8, 14, 22}

func (i KeyEvent) String() string {
	if i+1 >= KeyEvent(len(_KeyEvent_index)) {
		return fmt.Sprintf("KeyEvent(%d)", i)
	}
	return _KeyEvent_name[_KeyEvent_index[i]:_KeyEvent_index[i+1]]
}

const _FocusEvent_name = "FOCUS_LOSTFOCUS_GOT"

var _FocusEvent_index = [...]uint8{0, 10, 19}

func (i FocusEvent) String() string {
	if i+1 >= FocusEvent(len(_FocusEvent_index)) {
		return fmt.Sprintf("FocusEvent(%d)", i)
	}
	return _FocusEvent_name[_FocusEvent_index[i]:_FocusEvent_index[i+1]]
}

const _ScrollEvent_name = "SCROLL_HOMESCROLL_ENDSCROLL_STEP_PLUSSCROLL_STEP_MINUSSCROLL_PAGE_PLUSSCROLL_PAGE_MINUSSCROLL_POSSCROLL_SLIDER_RELEASEDSCROLL_CORNER_PRESSEDSCROLL_CORNER_RELEASED"

var _ScrollEvent_index = [...]uint8{0, 11, 21, 37, 54, 70, 87, 97, 119, 140, 162}

func (i ScrollEvent) String() string {
	if i+1 >= ScrollEvent(len(_ScrollEvent_index)) {
		return fmt.Sprintf("ScrollEvent(%d)", i)
	}
	return _ScrollEvent_name[_ScrollEvent_index[i]:_ScrollEvent_index[i+1]]
}

const _GestureCmd_name = "GESTURE_REQUESTGESTURE_ZOOMGESTURE_PANGESTURE_ROTATEGESTURE_TAP1GESTURE_TAP2"

var _GestureCmd_index = [...]uint8{0, 15, 27, 38, 52, 64, 76}

func (i GestureCmd) String() string {
	if i+1 >= GestureCmd(len(_GestureCmd_index)) {
		return fmt.Sprintf("GestureCmd(%d)", i)
	}
	return _GestureCmd_name[_GestureCmd_index[i]:_GestureCmd_index[i+1]]
}

const (
	_GestureState_name_0 = "GESTURE_STATE_BEGINGESTURE_STATE_INERTIA"
	_GestureState_name_1 = "GESTURE_STATE_END"
)

var (
	_GestureState_index_0 = [...]uint8{0, 19, 40}
	_GestureState_index_1 = [...]uint8{0, 17}
)

func (i GestureState) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _GestureState_name_0[_GestureState_index_0[i]:_GestureState_index_0[i+1]]
	case i == 4:
		return _GestureState_name_1
	default:
		return fmt.Sprintf("GestureState(%d)", i)
	}
}

const (
	_GestureTypeFlag_name_0 = "GESTURE_FLAG_ZOOMGESTURE_FLAG_ROTATE"
	_GestureTypeFlag_name_1 = "GESTURE_FLAG_PAN_VERTICAL"
	_GestureTypeFlag_name_2 = "GESTURE_FLAG_TAP1"
	_GestureTypeFlag_name_3 = "GESTURE_FLAG_TAP2"
	_GestureTypeFlag_name_4 = "GESTURE_FLAG_PAN_WITH_GUTTER"
	_GestureTypeFlag_name_5 = "GESTURE_FLAG_PAN_WITH_INERTIA"
	_GestureTypeFlag_name_6 = "GESTURE_FLAGS_ALL"
)

var (
	_GestureTypeFlag_index_0 = [...]uint8{0, 17, 36}
	_GestureTypeFlag_index_1 = [...]uint8{0, 25}
	_GestureTypeFlag_index_2 = [...]uint8{0, 17}
	_GestureTypeFlag_index_3 = [...]uint8{0, 17}
	_GestureTypeFlag_index_4 = [...]uint8{0, 28}
	_GestureTypeFlag_index_5 = [...]uint8{0, 29}
	_GestureTypeFlag_index_6 = [...]uint8{0, 17}
)

func (i GestureTypeFlag) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _GestureTypeFlag_name_0[_GestureTypeFlag_index_0[i]:_GestureTypeFlag_index_0[i+1]]
	case i == 4:
		return _GestureTypeFlag_name_1
	case i == 16:
		return _GestureTypeFlag_name_2
	case i == 32:
		return _GestureTypeFlag_name_3
	case i == 16384:
		return _GestureTypeFlag_name_4
	case i == 32768:
		return _GestureTypeFlag_name_5
	case i == 65535:
		return _GestureTypeFlag_name_6
	default:
		return fmt.Sprintf("GestureTypeFlag(%d)", i)
	}
}

const _DrawEvent_name = "DRAW_CONTENTDRAW_FOREGROUND"

var _DrawEvent_index = [...]uint8{0, 12, 27}

func (i DrawEvent) String() string {
	i -= 1
	if i+1 >= DrawEvent(len(_DrawEvent_index)) {
		return fmt.Sprintf("DrawEvent(%d)", i+1)
	}
	return _DrawEvent_name[_DrawEvent_index[i]:_DrawEvent_index[i+1]]
}

const _EventReason_name = "BY_MOUSE_CLICKBY_KEY_CLICKSYNTHESIZED"

var _EventReason_index = [...]uint8{0, 14, 26, 37}

func (i EventReason) String() string {
	if i+1 >= EventReason(len(_EventReason_index)) {
		return fmt.Sprintf("EventReason(%d)", i)
	}
	return _EventReason_name[_EventReason_index[i]:_EventReason_index[i+1]]
}

const _EditChangedReason_name = "BY_INS_CHARBY_INS_CHARSBY_DEL_CHARBY_DEL_CHARS"

var _EditChangedReason_index = [...]uint8{0, 11, 23, 34, 46}

func (i EditChangedReason) String() string {
	if i+1 >= EditChangedReason(len(_EditChangedReason_index)) {
		return fmt.Sprintf("EditChangedReason(%d)", i)
	}
	return _EditChangedReason_name[_EditChangedReason_index[i]:_EditChangedReason_index[i+1]]
}

const (
	_BehaviorMethodIdentifier_name_0 = "DO_CLICKGET_TEXT_VALUESET_TEXT_VALUETEXT_EDIT_GET_SELECTIONTEXT_EDIT_SET_SELECTIONTEXT_EDIT_REPLACE_SELECTIONSCROLL_BAR_GET_VALUESCROLL_BAR_SET_VALUETEXT_EDIT_GET_CARET_POSITIONTEXT_EDIT_GET_SELECTION_TEXTTEXT_EDIT_GET_SELECTION_HTMLTEXT_EDIT_CHAR_POS_AT_XY"
	_BehaviorMethodIdentifier_name_1 = "IS_EMPTYGET_VALUESET_VALUE"
	_BehaviorMethodIdentifier_name_2 = "FIRST_APPLICATION_METHOD_ID"
)

var (
	_BehaviorMethodIdentifier_index_0 = [...]uint16{0, 8, 22, 36, 59, 82, 109, 129, 149, 177, 205, 233, 257}
	_BehaviorMethodIdentifier_index_1 = [...]uint8{0, 8, 17, 26}
	_BehaviorMethodIdentifier_index_2 = [...]uint8{0, 27}
)

func (i BehaviorMethodIdentifier) String() string {
	switch {
	case 0 <= i && i <= 11:
		return _BehaviorMethodIdentifier_name_0[_BehaviorMethodIdentifier_index_0[i]:_BehaviorMethodIdentifier_index_0[i+1]]
	case 252 <= i && i <= 254:
		i -= 252
		return _BehaviorMethodIdentifier_name_1[_BehaviorMethodIdentifier_index_1[i]:_BehaviorMethodIdentifier_index_1[i+1]]
	case i == 256:
		return _BehaviorMethodIdentifier_name_2
	default:
		return fmt.Sprintf("BehaviorMethodIdentifier(%d)", i)
	}
}

const _SCDOM_RESULT_name = "SCDOM_OK_NOT_HANDLEDSCDOM_OKSCDOM_INVALID_HWNDSCDOM_INVALID_HANDLESCDOM_PASSIVE_HANDLESCDOM_INVALID_PARAMETERSCDOM_OPERATION_FAILED"

var _SCDOM_RESULT_index = [...]uint8{0, 20, 28, 46, 66, 86, 109, 131}

func (i SCDOM_RESULT) String() string {
	i -= -1
	if i < 0 || i+1 >= SCDOM_RESULT(len(_SCDOM_RESULT_index)) {
		return fmt.Sprintf("SCDOM_RESULT(%d)", i+-1)
	}
	return _SCDOM_RESULT_name[_SCDOM_RESULT_index[i]:_SCDOM_RESULT_index[i+1]]
}

const _VALUE_RESULT_name = "HV_OK_TRUEHV_OKHV_BAD_PARAMETERHV_INCOMPATIBLE_TYPE"

var _VALUE_RESULT_index = [...]uint8{0, 10, 15, 31, 51}

func (i VALUE_RESULT) String() string {
	i -= -1
	if i < 0 || i+1 >= VALUE_RESULT(len(_VALUE_RESULT_index)) {
		return fmt.Sprintf("VALUE_RESULT(%d)", i+-1)
	}
	return _VALUE_RESULT_name[_VALUE_RESULT_index[i]:_VALUE_RESULT_index[i+1]]
}
