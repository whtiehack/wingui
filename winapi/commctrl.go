package winapi

import "github.com/lxn/win"

/*
 * NOTE: All Message Numbers below 0x0400 are RESERVED.
 *
 * Private Window Messages Start Here:
 */
// #define WM_USER                         0x0400
const (
	WM_USER = win.WM_USER

	// /*****************************************************************************\
	// *                                                                             *
	// * commctrl.h - - Interface for the Windows Common Controls                    *
	// *                                                                             *
	// * Version 1.2                                                                 *
	// *                                                                             *
	// * Copyright (c) Microsoft Corporation. All rights reserved.                   *
	// *                                                                             *
	// \*****************************************************************************/

	// // #ifndef _INC_COMMCTRL
	// // #define _INC_COMMCTRL

	// #if defined(_MSC_VER) && (_MSC_VER >= 1200)
	// // #pragma warning(push)
	// // #pragma warning(disable:4001) /* nonstandard extension : single line comment */
	// // #pragma warning(disable:4201) /* nonstandard extension used : nameless struct/union */
	// // #pragma warning(disable:4820) /* padding added after data member */
	// // #pragma once
	// // #endif
	// // #include <winapifamily.h>

	// // #pragma region Desktop Family
	// // #if WINAPI_FAMILY_PARTITION(WINAPI_PARTITION_DESKTOP)

	// #ifndef _HRESULT_DEFINED
	// _HRESULT_DEFINED = typedef _Return_type_success_(return >= 0) long HRESULT;
	// #endif // !_HRESULT_DEFINED

	// #ifndef NOUSER

	// //
	// // Define API decoration for direct importing of DLL references.
	// //
	// #ifndef WINCOMMCTRLAPI
	// #if !defined(_COMCTL32_) && defined(_WIN32)
	// WINCOMMCTRLAPI = DECLSPEC_IMPORT
	// #else
	// #define WINCOMMCTRLAPI
	// #endif
	// #endif // WINCOMMCTRLAPI

	// //
	// // For compilers that don't support nameless unions
	// //
	// #ifndef DUMMYUNIONNAME
	// #ifdef NONAMELESSUNION
	//DUMMYUNIONNAME = u
	//DUMMYUNIONNAME2 = u2
	//DUMMYUNIONNAME3 = u3
	//DUMMYUNIONNAME4 = u4
	//DUMMYUNIONNAME5 = u5
	// #else
	// #define DUMMYUNIONNAME
	// #define DUMMYUNIONNAME2
	// #define DUMMYUNIONNAME3
	// #define DUMMYUNIONNAME4
	// #define DUMMYUNIONNAME5
	// #endif
	// #endif // DUMMYUNIONNAME

	// #ifdef __cplusplus
	// extern "C" {
	// #endif

	// //
	// // Users of this header may define any number of these constants to avoid
	// // the definitions of each functional group.
	// //
	// //    NOTOOLBAR    Customizable bitmap-button toolbar control.
	// //    NOUPDOWN     Up and Down arrow increment/decrement control.
	// //    NOSTATUSBAR  Status bar control.
	// //    NOMENUHELP   APIs to help manage menus, especially with a status bar.
	// //    NOTRACKBAR   Customizable column-width tracking control.
	// //    NODRAGLIST   APIs to make a listbox source and sink drag&drop actions.
	// //    NOPROGRESS   Progress gas gauge.
	// //    NOHOTKEY     HotKey control
	// //    NOHEADER     Header bar control.
	// //    NOIMAGEAPIS  ImageList apis.
	// //    NOLISTVIEW   ListView control.
	// //    NOTREEVIEW   TreeView control.
	// //    NOTABCONTROL Tab control.
	// //    NOANIMATE    Animate control.
	// //    NOBUTTON     Button control.
	// //    NOSTATIC     Static control.
	// //    NOEDIT       Edit control.
	// //    NOLISTBOX    Listbox control.
	// //    NOCOMBOBOX   Combobox control.
	// //    NOSCROLLBAR  Scrollbar control.
	// //    NOTASKDIALOG Task Dialog.
	// //
	//=============================================================================

	// #include <prsht.h>

	// #ifndef SNDMSG
	// #ifdef __cplusplus
	// #ifndef _MAC
	// #define SNDMSG ::SendMessage
	// #else
	// #define SNDMSG ::AfxSendMessage
	// #endif
	// #else
	// #ifndef _MAC
	// SNDMSG = SendMessage
	// #else
	// SNDMSG = AfxSendMessage
	// #endif //_MAC
	// #endif
	// #endif // ifndef SNDMSG

	// #ifdef _MAC
	// #ifndef RC_INVOKED
	// #ifndef _WLM_NOFORCE_LIBS

	// #ifndef _WLMDLL
	//     #ifdef _DEBUG
	//         #pragma comment(lib, "comctld.lib")
	//     #else
	//         #pragma comment(lib, "comctl.lib")
	//     #endif
	//     #pragma comment(linker, "/macres:comctl.rsc")
	//     #else
	//     #ifdef _DEBUG
	//         #pragma comment(lib, "msvcctld.lib")
	//     #else
	//         #pragma comment(lib, "msvcctl.lib")
	//     #endif
	// #endif // _WLMDLL

	// #endif // _WLM_NOFORCE_LIBS
	// #endif // RC_INVOKED
	// #endif //_MAC

	// WINCOMMCTRLAPI void WINAPI InitCommonControls(void);

	// typedef struct tagINITCOMMONCONTROLSEX {
	//     DWORD dwSize;             // size of this structure
	//     DWORD dwICC;              // flags indicating which classes to be initialized
	// } INITCOMMONCONTROLSEX, *LPINITCOMMONCONTROLSEX;
	ICC_LISTVIEW_CLASSES   = 0 // listview, header
	ICC_TREEVIEW_CLASSES   = 0 // treeview, tooltips
	ICC_BAR_CLASSES        = 0 // toolbar, statusbar, trackbar, tooltips
	ICC_TAB_CLASSES        = 0 // tab, tooltips
	ICC_UPDOWN_CLASS       = 0 // updown
	ICC_PROGRESS_CLASS     = 0 // progress
	ICC_HOTKEY_CLASS       = 0 // hotkey
	ICC_ANIMATE_CLASS      = 0 // animate
	ICC_WIN95_CLASSES      = 0
	ICC_DATE_CLASSES       = 0 // month picker, date picker, time picker, updown
	ICC_USEREX_CLASSES     = 0 // comboex
	ICC_COOL_CLASSES       = 0 // rebar (coolbar) control
	ICC_INTERNET_CLASSES   = 0
	ICC_PAGESCROLLER_CLASS = 0 // page scroller
	ICC_NATIVEFNTCTL_CLASS = 0 // native font control
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	ICC_STANDARD_CLASSES = 0
	ICC_LINK_CLASS       = 0
	// #endif // (NTDDI_VERSION >= NTDDI_WINXP)

	// WINCOMMCTRLAPI BOOL WINAPI InitCommonControlsEx(_In_ const INITCOMMONCONTROLSEX *picce);

	ODT_HEADER   = 100
	ODT_TAB      = 101
	ODT_LISTVIEW = 102

	//====== Ranges for control message IDs =======================================
	// TODO: need fix value
	LVM_FIRST = 0 // ListView messages
	TV_FIRST  = 0 // TreeView messages
	HDM_FIRST = 0 // Header messages
	TCM_FIRST = 0x1300 // Tab control messages

	PGM_FIRST = 0 // Pager control messages

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	ECM_FIRST = 0 // Edit control messages
	BCM_FIRST = 0 // Button control messages
	CBM_FIRST = 0 // Combobox control messages
	// #endif // (NTDDI_VERSION >= NTDDI_WINXP)

	CCM_FIRST = 0 // Common control shared messages
	CCM_LAST  = (CCM_FIRST + 0x200)

	CCM_SETBKCOLOR = (CCM_FIRST + 1) // lParam is bkColor

	// typedef struct tagCOLORSCHEME {
	//    DWORD            dwSize;
	//    COLORREF         clrBtnHighlight;       // highlight color
	//    COLORREF         clrBtnShadow;          // shadow color
	// } COLORSCHEME, *LPCOLORSCHEME;

	CCM_SETCOLORSCHEME   = (CCM_FIRST + 2) // lParam is color scheme
	CCM_GETCOLORSCHEME   = (CCM_FIRST + 3) // fills in COLORSCHEME pointed to by lParam
	CCM_GETDROPTARGET    = (CCM_FIRST + 4)
	CCM_SETUNICODEFORMAT = (CCM_FIRST + 5)
	CCM_GETUNICODEFORMAT = (CCM_FIRST + 6)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	COMCTL32_VERSION = 6
	// #else
	//COMCTL32_VERSION = 5
	// #endif

	CCM_SETVERSION      = (CCM_FIRST + 0x7)
	CCM_GETVERSION      = (CCM_FIRST + 0x8)
	CCM_SETNOTIFYWINDOW = (CCM_FIRST + 0x9) // wParam == hwndParent.
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	CCM_SETWINDOWTHEME = (CCM_FIRST + 0xb)
	CCM_DPISCALE       = (CCM_FIRST + 0xc) // wParam == Awareness
	// #endif

	// // for tooltips
	INFOTIPSIZE = 1024

	//====== WM_NOTIFY Macros =====================================================

	// #define HANDLE_WM_NOTIFY(hwnd, wParam, lParam, fn) \
	//     (fn)((hwnd), (int)(wParam), (NMHDR *)(lParam))
	// #define FORWARD_WM_NOTIFY(hwnd, idFrom, pnmhdr, fn) \
	//     (LRESULT)(fn)((hwnd), WM_NOTIFY, (WPARAM)(int)(idFrom), (LPARAM)(NMHDR *)(pnmhdr))

	//====== Generic WM_NOTIFY notification codes =================================

	NM_OUTOFMEMORY     = (NM_FIRST - 1)
	NM_CLICK           = (NM_FIRST - 2) // uses NMCLICK struct
	NM_DBLCLK          = (NM_FIRST - 3)
	NM_RETURN          = (NM_FIRST - 4)
	NM_RCLICK          = (NM_FIRST - 5) // uses NMCLICK struct
	NM_RDBLCLK         = (NM_FIRST - 6)
	NM_SETFOCUS        = (NM_FIRST - 7)
	NM_KILLFOCUS       = (NM_FIRST - 8)
	NM_CUSTOMDRAW      = (NM_FIRST - 12)
	NM_HOVER           = (NM_FIRST - 13)
	NM_NCHITTEST       = (NM_FIRST - 14) // uses NMMOUSE struct
	NM_KEYDOWN         = (NM_FIRST - 15) // uses NMKEY struct
	NM_RELEASEDCAPTURE = (NM_FIRST - 16)
	NM_SETCURSOR       = (NM_FIRST - 17) // uses NMMOUSE struct
	NM_CHAR            = (NM_FIRST - 18) // uses NMCHAR struct
	NM_TOOLTIPSCREATED = (NM_FIRST - 19) // notify of when the tooltips window is create
	NM_LDOWN           = (NM_FIRST - 20)
	NM_RDOWN           = (NM_FIRST - 21)
	NM_THEMECHANGED    = (NM_FIRST - 22)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	NM_FONTCHANGED          = (NM_FIRST - 23)
	NM_CUSTOMTEXT           = (NM_FIRST - 24) // uses NMCUSTOMTEXT struct
	NM_TVSTATEIMAGECHANGING = (NM_FIRST - 24) // uses NMTVSTATEIMAGECHANGING struct, defined after HTREEITEM
	// #endif

	// #ifndef CCSIZEOF_STRUCT
	// #define CCSIZEOF_STRUCT(structname, member)  (((int)((LPBYTE)(&((structname*)0)->member) - ((LPBYTE)((structname*)0)))) + sizeof(((structname*)0)->member))
	// #endif

	//====== Generic WM_NOTIFY notification structures ============================
	// typedef struct tagNMTOOLTIPSCREATED
	// {
	//     NMHDR hdr;
	//     HWND hwndToolTips;
	// } NMTOOLTIPSCREATED, * LPNMTOOLTIPSCREATED;

	// typedef struct tagNMMOUSE {
	//     NMHDR   hdr;
	//     DWORD_PTR dwItemSpec;
	//     DWORD_PTR dwItemData;
	//     POINT   pt;
	//     LPARAM  dwHitInfo; // any specifics about where on the item or control the mouse is
	// } NMMOUSE, *LPNMMOUSE;

	// typedef NMMOUSE NMCLICK;
	// typedef LPNMMOUSE LPNMCLICK;

	// // Generic structure to request an object of a specific type.

	// typedef struct tagNMOBJECTNOTIFY {
	//     NMHDR   hdr;
	//     int     iItem;
	// #ifdef __IID_DEFINED__
	//     const IID *piid;
	// #else
	//     const void *piid;
	// #endif
	//     void *pObject;
	//     HRESULT hResult;
	//     DWORD dwFlags;    // control specific flags (hints as to where in iItem it hit)
	// } NMOBJECTNOTIFY, *LPNMOBJECTNOTIFY;

	// // Generic structure for a key

	// typedef struct tagNMKEY
	// {
	//     NMHDR hdr;
	//     UINT  nVKey;
	//     UINT  uFlags;
	// } NMKEY, *LPNMKEY;

	// // Generic structure for a character

	// typedef struct tagNMCHAR {
	//     NMHDR   hdr;
	//     UINT    ch;
	//     DWORD   dwItemPrev;     // Item previously selected
	//     DWORD   dwItemNext;     // Item to be selected
	// } NMCHAR, *LPNMCHAR;

	// #if (_WIN32_IE >= 0x0600)

	// typedef struct tagNMCUSTOMTEXT
	// {
	//     NMHDR hdr;
	//     HDC hDC;
	//     LPCWSTR lpString;
	//     int nCount;
	//     LPRECT lpRect;
	//     UINT uFormat;
	//     BOOL fLink;
	// } NMCUSTOMTEXT, *LPNMCUSTOMTEXT;

	// #endif           // _WIN32_IE >= 0x0600
	//====== WM_NOTIFY codes (NMHDR.code values) ==================================

	NM_FIRST = (0 - 0) // generic to all controls
	NM_LAST  = (0 - 99)

	LVN_FIRST = (0 - 100) // listview
	LVN_LAST  = (0 - 199)

	// // Property sheet reserved      (0-200) -  (0-299) - see prsht.h

	HDN_FIRST = (0 - 300) // header
	HDN_LAST  = (0 - 399)

	TVN_FIRST = (0 - 400) // treeview
	TVN_LAST  = (0 - 499)

	TTN_FIRST = (0 - 520) // tooltips
	TTN_LAST  = (0 - 549)

	TCN_FIRST = (0 - 550) // tab control
	TCN_LAST  = (0 - 580)

	// // Shell reserved               (0-580) -  (0-589)

	CDN_FIRST = (0 - 601) // common dialog (new)
	CDN_LAST  = (0 - 699)

	TBN_FIRST = (0 - 700) // toolbar
	TBN_LAST  = (0 - 720)

	UDN_FIRST = (0 - 721) // updown
	UDN_LAST  = (0 - 729)
	DTN_FIRST = (0 - 740) // datetimepick
	DTN_LAST  = (0 - 745) // DTN_FIRST - 5

	MCN_FIRST = (0 - 746) // monthcal
	MCN_LAST  = (0 - 752) // MCN_FIRST - 6

	DTN_FIRST2 = (0 - 753) // datetimepick2
	DTN_LAST2  = (0 - 799)

	CBEN_FIRST = (0 - 800) // combo box ex
	CBEN_LAST  = (0 - 830)

	RBN_FIRST = (0 - 831) // rebar
	RBN_LAST  = (0 - 859)

	IPN_FIRST = (0 - 860) // internet address
	IPN_LAST  = (0 - 879) // internet address

	SBN_FIRST = (0 - 880) // status bar
	SBN_LAST  = (0 - 899)

	PGN_FIRST = (0 - 900) // Pager Control
	PGN_LAST  = (0 - 950)

	// #ifndef WMN_FIRST
	WMN_FIRST = (0 - 1000)
	WMN_LAST  = (0 - 1200)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	BCN_FIRST = (0 - 1250)
	BCN_LAST  = (0 - 1350)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TRBN_FIRST = (0 - 1501) // trackbar
	TRBN_LAST  = (0 - 1519)
	// #endif

	MSGF_COMMCTRL_BEGINDRAG   = 0
	MSGF_COMMCTRL_SIZEHEADER  = 0
	MSGF_COMMCTRL_DRAGSELECT  = 0
	MSGF_COMMCTRL_TOOLBARCUST = 0

	//==================== CUSTOM DRAW ==========================================

	// // custom draw return flags
	// // values under 0x00010000 are reserved for global custom draw values.
	// // above that are for specific controls
	CDRF_DODEFAULT     = 0
	CDRF_NEWFONT       = 0
	CDRF_SKIPDEFAULT   = 0
	CDRF_DOERASE       = 0 // draw the background
	CDRF_SKIPPOSTPAINT = 0 // don't draw the focus rect

	CDRF_NOTIFYPOSTPAINT   = 0
	CDRF_NOTIFYITEMDRAW    = 0
	CDRF_NOTIFYSUBITEMDRAW = 0 // flags are the same, we can distinguish by context
	CDRF_NOTIFYPOSTERASE   = 0

	// // drawstage flags
	// // values under 0x00010000 are reserved for global custom draw values.
	// // above that are for specific controls
	CDDS_PREPAINT  = 0
	CDDS_POSTPAINT = 0
	CDDS_PREERASE  = 0
	CDDS_POSTERASE = 0
	// // the 0x000010000 bit means it's individual item specific
	CDDS_ITEM          = 0
	CDDS_ITEMPREPAINT  = (CDDS_ITEM | CDDS_PREPAINT)
	CDDS_ITEMPOSTPAINT = (CDDS_ITEM | CDDS_POSTPAINT)
	CDDS_ITEMPREERASE  = (CDDS_ITEM | CDDS_PREERASE)
	CDDS_ITEMPOSTERASE = (CDDS_ITEM | CDDS_POSTERASE)
	CDDS_SUBITEM       = 0

	// // itemState flags
	CDIS_SELECTED      = 0
	CDIS_GRAYED        = 0
	CDIS_DISABLED      = 0
	CDIS_CHECKED       = 0
	CDIS_FOCUS         = 0
	CDIS_DEFAULT       = 0
	CDIS_HOT           = 0
	CDIS_MARKED        = 0
	CDIS_INDETERMINATE = 0
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	CDIS_SHOWKEYBOARDCUES = 0
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	CDIS_NEARHOT      = 0
	CDIS_OTHERSIDEHOT = 0
	CDIS_DROPHILITED  = 0
	// #endif

	// typedef struct tagNMCUSTOMDRAWINFO
	// {
	//     NMHDR hdr;
	//     DWORD dwDrawStage;
	//     HDC hdc;
	//     RECT rc;
	//     DWORD_PTR dwItemSpec;  // this is control specific, but it's how to specify an item.  valid only with CDDS_ITEM bit set
	//     UINT  uItemState;
	//     LPARAM lItemlParam;
	// } NMCUSTOMDRAW, *LPNMCUSTOMDRAW;

	// typedef struct tagNMTTCUSTOMDRAW
	// {
	//     NMCUSTOMDRAW nmcd;
	//     UINT uDrawFlags;
	// } NMTTCUSTOMDRAW, *LPNMTTCUSTOMDRAW;

	// typedef struct tagNMCUSTOMSPLITRECTINFO
	// {
	//     NMHDR hdr;
	//     RECT rcClient;
	//     RECT rcButton;
	//     RECT rcSplit;
	// } NMCUSTOMSPLITRECTINFO, *LPNMCUSTOMSPLITRECTINFO;

	NM_GETCUSTOMSPLITRECT = (BCN_FIRST + 0x0003)

	//====== IMAGE APIS ===========================================================

	// #ifndef NOIMAGEAPIS

	CLR_NONE    = 0
	CLR_DEFAULT = 0

	// #ifndef HIMAGELIST
	// struct _IMAGELIST;
	// typedef struct _IMAGELIST* HIMAGELIST;
	// #endif

	// #ifndef IMAGELISTDRAWPARAMS
	// typedef struct _IMAGELISTDRAWPARAMS
	// {
	//     DWORD       cbSize;
	//     HIMAGELIST  himl;
	//     int         i;
	//     HDC         hdcDst;
	//     int         x;
	//     int         y;
	//     int         cx;
	//     int         cy;
	//     int         xBitmap;        // x offest from the upperleft of bitmap
	//     int         yBitmap;        // y offset from the upperleft of bitmap
	//     COLORREF    rgbBk;
	//     COLORREF    rgbFg;
	//     UINT        fStyle;
	//     DWORD       dwRop;
	// #if (_WIN32_IE >= 0x0501)
	//     DWORD       fState;
	//     DWORD       Frame;
	//     COLORREF    crEffect;
	// #endif
	// } IMAGELISTDRAWPARAMS, *LPIMAGELISTDRAWPARAMS;

	//IMAGELISTDRAWPARAMS_V3_SIZE = CCSIZEOF_STRUCT(IMAGELISTDRAWPARAMS, dwRop)

	// #endif

	ILC_MASK     = 0
	ILC_COLOR    = 0
	ILC_COLORDDB = 0
	ILC_COLOR4   = 0
	ILC_COLOR8   = 0
	ILC_COLOR16  = 0
	ILC_COLOR24  = 0
	ILC_COLOR32  = 0
	ILC_PALETTE  = 0 // (not implemented)
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	ILC_MIRROR        = 0 // Mirror the icons contained, if the process is mirrored
	ILC_PERITEMMIRROR = 0 // Causes the mirroring code to mirror each item when inserting a set of images, verses the whole strip
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	ILC_ORIGINALSIZE     = 0 // Imagelist should accept smaller than set images and apply OriginalSize based on image added
	ILC_HIGHQUALITYSCALE = 0 // Imagelist should enable use of the high quality scaler.
	// #endif
	// WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_Create(int cx, int cy, UINT flags, int cInitial, int cGrow);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_Destroy(_In_opt_ HIMAGELIST himl);

	// WINCOMMCTRLAPI int         WINAPI ImageList_GetImageCount(_In_ HIMAGELIST himl);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_SetImageCount(_In_ HIMAGELIST himl, _In_ UINT uNewCount);

	// WINCOMMCTRLAPI int         WINAPI ImageList_Add(_In_ HIMAGELIST himl, _In_ HBITMAP hbmImage, _In_opt_ HBITMAP hbmMask);

	// WINCOMMCTRLAPI int         WINAPI ImageList_ReplaceIcon(_In_ HIMAGELIST himl, _In_ int i, _In_ HICON hicon);
	// WINCOMMCTRLAPI COLORREF    WINAPI ImageList_SetBkColor(_In_ HIMAGELIST himl, _In_ COLORREF clrBk);
	// WINCOMMCTRLAPI COLORREF    WINAPI ImageList_GetBkColor(_In_ HIMAGELIST himl);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_SetOverlayImage(_In_ HIMAGELIST himl, _In_ int iImage, _In_ int iOverlay);

	// #define     ImageList_AddIcon(himl, hicon) ImageList_ReplaceIcon(himl, -1, hicon)

	ILD_NORMAL      = 0
	ILD_TRANSPARENT = 0
	ILD_MASK        = 0
	ILD_IMAGE       = 0
	ILD_ROP         = 0
	ILD_BLEND25     = 0
	ILD_BLEND50     = 0
	ILD_OVERLAYMASK = 0
	// #define INDEXTOOVERLAYMASK(i)   ((i) << 8)
	ILD_PRESERVEALPHA = 0 // This preserves the alpha channel in dest
	ILD_SCALE         = 0 // Causes the image to be scaled to cx, cy instead of clipped
	ILD_DPISCALE      = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	ILD_ASYNC = 0
	// #endif

	ILD_SELECTED = ILD_BLEND50
	ILD_FOCUS    = ILD_BLEND25
	ILD_BLEND    = ILD_BLEND50
	CLR_HILIGHT  = CLR_DEFAULT

	ILS_NORMAL   = 0
	ILS_GLOW     = 0
	ILS_SHADOW   = 0
	ILS_SATURATE = 0
	ILS_ALPHA    = 0

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	ILGT_NORMAL = 0
	ILGT_ASYNC  = 0
	// #endif

	// WINCOMMCTRLAPI BOOL WINAPI ImageList_Draw(_In_ HIMAGELIST himl, _In_ int i, _In_ HDC hdcDst, _In_ int x, _In_ int y, _In_ UINT fStyle);

	// #ifdef _WIN32

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//HBITMAP_CALLBACK = ((HBITMAP)-1)       // only for SparseImageList
	// #endif

	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_Replace(_In_ HIMAGELIST himl, _In_ int i, _In_ HBITMAP hbmImage, _In_opt_ HBITMAP hbmMask);

	// WINCOMMCTRLAPI int         WINAPI ImageList_AddMasked(_In_ HIMAGELIST himl, _In_ HBITMAP hbmImage, _In_ COLORREF crMask);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_DrawEx(_In_ HIMAGELIST himl, _In_ int i, _In_ HDC hdcDst, _In_ int x, _In_ int y, _In_ int dx, _In_ int dy, _In_ COLORREF rgbBk, _In_ COLORREF rgbFg, _In_ UINT fStyle);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_DrawIndirect(_In_ IMAGELISTDRAWPARAMS* pimldp);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_Remove(_In_ HIMAGELIST himl, _In_ int i);
	// WINCOMMCTRLAPI HICON       WINAPI ImageList_GetIcon(_In_ HIMAGELIST himl, _In_ int i, _In_ UINT flags);
	// WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_LoadImageA(HINSTANCE hi, LPCSTR lpbmp, int cx, int cGrow, COLORREF crMask, UINT uType, UINT uFlags);
	// WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_LoadImageW(HINSTANCE hi, LPCWSTR lpbmp, int cx, int cGrow, COLORREF crMask, UINT uType, UINT uFlags);

	// #ifdef UNICODE
	// ImageList_LoadImage = ImageList_LoadImageW
	// #else
	// ImageList_LoadImage = ImageList_LoadImageA
	// #endif

	ILCF_MOVE = (0x00000000)
	ILCF_SWAP = (0x00000001)
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_Copy(_In_ HIMAGELIST himlDst, _In_ int iDst, _In_ HIMAGELIST himlSrc, _In_ int iSrc, _In_ UINT uFlags);

	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_BeginDrag(_In_ HIMAGELIST himlTrack, _In_ int iTrack, _In_ int dxHotspot, _In_ int dyHotspot);
	// WINCOMMCTRLAPI void        WINAPI ImageList_EndDrag(void);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_DragEnter(HWND hwndLock, int x, int y);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_DragLeave(HWND hwndLock);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_DragMove(int x, int y);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_SetDragCursorImage(_In_ HIMAGELIST himlDrag, _In_ int iDrag, _In_ int dxHotspot, _In_ int dyHotspot);

	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_DragShowNolock(BOOL fShow);
	// _Success_(return != NULL) WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_GetDragImage(_Out_opt_ POINT *ppt, _Out_opt_ POINT *pptHotspot);

	// #define     ImageList_RemoveAll(himl) ImageList_Remove(himl, -1)
	// #define     ImageList_ExtractIcon(hi, himl, i) ImageList_GetIcon(himl, i, 0)
	// #define     ImageList_LoadBitmap(hi, lpbmp, cx, cGrow, crMask) ImageList_LoadImage(hi, lpbmp, cx, cGrow, crMask, IMAGE_BITMAP, 0)

	// struct IStream;

	// WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_Read(_In_ struct IStream *pstm);
	// WINCOMMCTRLAPI BOOL        WINAPI ImageList_Write(_In_ HIMAGELIST himl, _In_ struct IStream *pstm);

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	ILP_NORMAL    = 0 // Writes or reads the stream using new sematics for this version of comctl32
	ILP_DOWNLEVEL = 1 // Write or reads the stream using downlevel sematics.

	// WINCOMMCTRLAPI HRESULT WINAPI ImageList_ReadEx(_In_ DWORD dwFlags, _In_ struct IStream *pstm, _In_ REFIID riid, _Outptr_ PVOID* ppv);
	// WINCOMMCTRLAPI HRESULT WINAPI ImageList_WriteEx(_In_ HIMAGELIST himl, _In_ DWORD dwFlags, _In_ struct IStream *pstm);
	// #endif

	// #ifndef IMAGEINFO
	// typedef struct _IMAGEINFO
	// {
	//     HBITMAP hbmImage;
	//     HBITMAP hbmMask;
	//     int     Unused1;
	//     int     Unused2;
	//     RECT    rcImage;
	// } IMAGEINFO, *LPIMAGEINFO;
	// #endif

	// _Success_(return) WINCOMMCTRLAPI BOOL        WINAPI ImageList_GetIconSize(_In_ HIMAGELIST himl, _Out_opt_ int *cx, _Out_opt_ int *cy);
	// _Success_(return) WINCOMMCTRLAPI BOOL        WINAPI ImageList_SetIconSize(_In_ HIMAGELIST himl, _In_ int cx, _In_ int cy);
	// _Success_(return) WINCOMMCTRLAPI BOOL        WINAPI ImageList_GetImageInfo(_In_ HIMAGELIST himl, _In_ int i, _Out_ IMAGEINFO *pImageInfo);
	// WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_Merge(_In_ HIMAGELIST himl1, _In_ int i1, _In_ HIMAGELIST himl2, _In_ int i2, _In_ int dx, _In_ int dy);
	// WINCOMMCTRLAPI HIMAGELIST  WINAPI ImageList_Duplicate(_In_ HIMAGELIST himl);

	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	// WINCOMMCTRLAPI HRESULT WINAPI HIMAGELIST_QueryInterface(_In_ HIMAGELIST himl, _In_ REFIID riid, _Outptr_ void **ppv);

	// #ifdef __cplusplus
	// FORCEINLINE HIMAGELIST IImageListToHIMAGELIST(struct IImageList *himl)
	// {
	//     return reinterpret_cast<HIMAGELIST>(himl);
	// }
	// #else
	// #define IImageListToHIMAGELIST(himl) ((HIMAGELIST)(himl))
	// #endif

	// #endif

	// #endif

	//====== HEADER CONTROL =======================================================

	// #ifndef NOHEADER

	// #ifdef _WIN32
	// #define WC_HEADERA              "SysHeader32"
	WC_HEADERW = "SysHeader32"

	// #ifdef UNICODE
	WC_HEADER = WC_HEADERW
	// #else
	//WC_HEADER = WC_HEADERA
	// #endif

	// #else
	// #define WC_HEADER               "SysHeader"
	// #endif

	// // begin_r_commctrl

	HDS_HORZ     = 0
	HDS_BUTTONS  = 0
	HDS_HOTTRACK = 0
	HDS_HIDDEN   = 0

	HDS_DRAGDROP  = 0
	HDS_FULLDRAG  = 0
	HDS_FILTERBAR = 0

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	HDS_FLAT = 0
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	HDS_CHECKBOXES = 0
	HDS_NOSIZING   = 0
	HDS_OVERFLOW   = 0
	// #endif
	// // end_r_commctrl

	HDFT_ISSTRING = 0 // HD_ITEM.pvFilter points to a HD_TEXTFILTER
	HDFT_ISNUMBER = 0 // HD_ITEM.pvFilter points to a INT
	HDFT_ISDATE   = 0 // HD_ITEM.pvFilter points to a DWORD (dos date)

	HDFT_HASNOVALUE = 0 // clear the filter, by setting this bit

	// #ifdef UNICODE
	//HD_TEXTFILTER = HD_TEXTFILTERW
	//HDTEXTFILTER = HD_TEXTFILTERW
	//LPHD_TEXTFILTER = LPHD_TEXTFILTERW
	//LPHDTEXTFILTER = LPHD_TEXTFILTERW
	// #else
	//HD_TEXTFILTER = HD_TEXTFILTERA
	//HDTEXTFILTER = HD_TEXTFILTERA
	//LPHD_TEXTFILTER = LPHD_TEXTFILTERA
	//LPHDTEXTFILTER = LPHD_TEXTFILTERA
	// #endif

	// typedef struct _HD_TEXTFILTERA
	// {
	//     LPSTR pszText;                      // [in] pointer to the buffer containing the filter (ANSI)
	//     INT cchTextMax;                     // [in] max size of buffer/edit control buffer
	// } HD_TEXTFILTERA, *LPHD_TEXTFILTERA;

	// typedef struct _HD_TEXTFILTERW
	// {
	//     LPWSTR pszText;                     // [in] pointer to the buffer contiaining the filter (UNICODE)
	//     INT cchTextMax;                     // [in] max size of buffer/edit control buffer
	// } HD_TEXTFILTERW, *LPHD_TEXTFILTERW;

	//HD_ITEMA = HDITEMA
	//HD_ITEMW = HDITEMW
	//HD_ITEM = HDITEM

	// typedef struct _HD_ITEMA
	// {
	//     UINT    mask;
	//     int     cxy;
	//     LPSTR   pszText;
	//     HBITMAP hbm;
	//     int     cchTextMax;
	//     int     fmt;
	//     LPARAM  lParam;
	//     int     iImage;        // index of bitmap in ImageList
	//     int     iOrder;        // where to draw this item
	//     UINT    type;           // [in] filter type (defined what pvFilter is a pointer to)
	//     void *  pvFilter;       // [in] filter data see above
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     UINT   state;
	// #endif
	// } HDITEMA, *LPHDITEMA;

	//HDITEMA_V1_SIZE = CCSIZEOF_STRUCT(HDITEMA, lParam)
	//HDITEMW_V1_SIZE = CCSIZEOF_STRUCT(HDITEMW, lParam)

	// typedef struct _HD_ITEMW
	// {
	//     UINT    mask;
	//     int     cxy;
	//     LPWSTR  pszText;
	//     HBITMAP hbm;
	//     int     cchTextMax;
	//     int     fmt;
	//     LPARAM  lParam;
	//     int     iImage;        // index of bitmap in ImageList
	//     int     iOrder;
	//     UINT    type;           // [in] filter type (defined what pvFilter is a pointer to)
	//     void *  pvFilter;       // [in] fillter data see above
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     UINT   state;
	// #endif
	// } HDITEMW, *LPHDITEMW;

	// #ifdef UNICODE
	//HDITEM = HDITEMW
	//LPHDITEM = LPHDITEMW
	//HDITEM_V1_SIZE = HDITEMW_V1_SIZE
	// #else
	//HDITEM = HDITEMA
	//LPHDITEM = LPHDITEMA
	//HDITEM_V1_SIZE = HDITEMA_V1_SIZE
	// #endif

	HDI_WIDTH      = 0
	HDI_HEIGHT     = HDI_WIDTH
	HDI_TEXT       = 0
	HDI_FORMAT     = 0
	HDI_LPARAM     = 0
	HDI_BITMAP     = 0
	HDI_IMAGE      = 0
	HDI_DI_SETITEM = 0
	HDI_ORDER      = 0
	HDI_FILTER     = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	HDI_STATE = 0
	// #endif

	// // HDF_ flags are shared with the listview control (LVCFMT_ flags)

	HDF_LEFT        = 0 // Same as LVCFMT_LEFT
	HDF_RIGHT       = 0 // Same as LVCFMT_RIGHT
	HDF_CENTER      = 0 // Same as LVCFMT_CENTER
	HDF_JUSTIFYMASK = 0 // Same as LVCFMT_JUSTIFYMASK
	HDF_RTLREADING  = 0 // Same as LVCFMT_LEFT

	HDF_BITMAP          = 0
	HDF_STRING          = 0
	HDF_OWNERDRAW       = 0 // Same as LVCFMT_COL_HAS_IMAGES
	HDF_IMAGE           = 0 // Same as LVCFMT_IMAGE
	HDF_BITMAP_ON_RIGHT = 0 // Same as LVCFMT_BITMAP_ON_RIGHT

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	HDF_SORTUP   = 0
	HDF_SORTDOWN = 0
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	HDF_CHECKBOX    = 0
	HDF_CHECKED     = 0
	HDF_FIXEDWIDTH  = 0 // Can't resize the column; same as LVCFMT_FIXED_WIDTH
	HDF_SPLITBUTTON = 0 // Column is a split button; same as LVCFMT_SPLITBUTTON
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	HDIS_FOCUSED = 0
	// #endif

	HDM_GETITEMCOUNT = (HDM_FIRST + 0)
	// #define Header_GetItemCount(hwndHD) \
	//     (int)SNDMSG((hwndHD), HDM_GETITEMCOUNT, 0, 0L)

	HDM_INSERTITEMA = (HDM_FIRST + 1)
	HDM_INSERTITEMW = (HDM_FIRST + 10)

	// #ifdef UNICODE
	//HDM_INSERTITEM = HDM_INSERTITEMW
	// #else
	//HDM_INSERTITEM = HDM_INSERTITEMA
	// #endif

	// #define Header_InsertItem(hwndHD, i, phdi) \
	//     (int)SNDMSG((hwndHD), HDM_INSERTITEM, (WPARAM)(int)(i), (LPARAM)(const HD_ITEM *)(phdi))

	HDM_DELETEITEM = (HDM_FIRST + 2)
	// #define Header_DeleteItem(hwndHD, i) \
	//     (BOOL)SNDMSG((hwndHD), HDM_DELETEITEM, (WPARAM)(int)(i), 0L)

	HDM_GETITEMA = (HDM_FIRST + 3)
	HDM_GETITEMW = (HDM_FIRST + 11)

	// #ifdef UNICODE
	//HDM_GETITEM = HDM_GETITEMW
	// #else
	//HDM_GETITEM = HDM_GETITEMA
	// #endif

	// #define Header_GetItem(hwndHD, i, phdi) \
	//     (BOOL)SNDMSG((hwndHD), HDM_GETITEM, (WPARAM)(int)(i), (LPARAM)(HD_ITEM *)(phdi))

	HDM_SETITEMA = (HDM_FIRST + 4)
	HDM_SETITEMW = (HDM_FIRST + 12)

	// #ifdef UNICODE
	//HDM_SETITEM = HDM_SETITEMW
	// #else
	//HDM_SETITEM = HDM_SETITEMA
	// #endif

	// #define Header_SetItem(hwndHD, i, phdi) \
	//     (BOOL)SNDMSG((hwndHD), HDM_SETITEM, (WPARAM)(int)(i), (LPARAM)(const HD_ITEM *)(phdi))

	//HD_LAYOUT = HDLAYOUT

	// typedef struct _HD_LAYOUT
	// {
	//     RECT *prc;
	//     WINDOWPOS *pwpos;
	// } HDLAYOUT, *LPHDLAYOUT;

	HDM_LAYOUT = (HDM_FIRST + 5)
	// #define Header_Layout(hwndHD, playout) \
	//     (BOOL)SNDMSG((hwndHD), HDM_LAYOUT, 0, (LPARAM)(HD_LAYOUT *)(playout))

	HHT_NOWHERE        = 0
	HHT_ONHEADER       = 0
	HHT_ONDIVIDER      = 0
	HHT_ONDIVOPEN      = 0
	HHT_ONFILTER       = 0
	HHT_ONFILTERBUTTON = 0
	HHT_ABOVE          = 0
	HHT_BELOW          = 0
	HHT_TORIGHT        = 0
	HHT_TOLEFT         = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	HHT_ONITEMSTATEICON = 0
	HHT_ONDROPDOWN      = 0
	HHT_ONOVERFLOW      = 0
	// #endif

	//HD_HITTESTINFO = HDHITTESTINFO

	// typedef struct _HD_HITTESTINFO
	// {
	//     POINT pt;
	//     UINT flags;
	//     int iItem;
	// } HDHITTESTINFO, *LPHDHITTESTINFO;

	HDSIL_NORMAL = 0
	HDSIL_STATE  = 1

	HDM_HITTEST = (HDM_FIRST + 6)

	HDM_GETITEMRECT = (HDM_FIRST + 7)
	// #define Header_GetItemRect(hwnd, iItem, lprc) \
	//         (BOOL)SNDMSG((hwnd), HDM_GETITEMRECT, (WPARAM)(iItem), (LPARAM)(lprc))

	HDM_SETIMAGELIST = (HDM_FIRST + 8)
	// #define Header_SetImageList(hwnd, himl) \
	//         (HIMAGELIST)SNDMSG((hwnd), HDM_SETIMAGELIST, HDSIL_NORMAL, (LPARAM)(himl))
	// #define Header_SetStateImageList(hwnd, himl) \
	//         (HIMAGELIST)SNDMSG((hwnd), HDM_SETIMAGELIST, HDSIL_STATE, (LPARAM)(himl))

	HDM_GETIMAGELIST = (HDM_FIRST + 9)
	// #define Header_GetImageList(hwnd) \
	//         (HIMAGELIST)SNDMSG((hwnd), HDM_GETIMAGELIST, HDSIL_NORMAL, 0)
	// #define Header_GetStateImageList(hwnd) \
	//         (HIMAGELIST)SNDMSG((hwnd), HDM_GETIMAGELIST, HDSIL_STATE, 0)

	HDM_ORDERTOINDEX = (HDM_FIRST + 15)
	// #define Header_OrderToIndex(hwnd, i) \
	//         (int)SNDMSG((hwnd), HDM_ORDERTOINDEX, (WPARAM)(i), 0)

	HDM_CREATEDRAGIMAGE = (HDM_FIRST + 16) // wparam = which item (by index)
	// #define Header_CreateDragImage(hwnd, i) \
	//         (HIMAGELIST)SNDMSG((hwnd), HDM_CREATEDRAGIMAGE, (WPARAM)(i), 0)

	HDM_GETORDERARRAY = (HDM_FIRST + 17)
	// #define Header_GetOrderArray(hwnd, iCount, lpi) \
	//         (BOOL)SNDMSG((hwnd), HDM_GETORDERARRAY, (WPARAM)(iCount), (LPARAM)(lpi))

	HDM_SETORDERARRAY = (HDM_FIRST + 18)
	// #define Header_SetOrderArray(hwnd, iCount, lpi) \
	//         (BOOL)SNDMSG((hwnd), HDM_SETORDERARRAY, (WPARAM)(iCount), (LPARAM)(lpi))
	// lparam = int array of size HDM_GETITEMCOUNT
	// // the array specifies the order that all items should be displayed.
	// // e.g.  { 2, 0, 1}
	// // says the index 2 item should be shown in the 0ths position
	// //      index 0 should be shown in the 1st position
	// //      index 1 should be shown in the 2nd position

	HDM_SETHOTDIVIDER = (HDM_FIRST + 19)
	// #define Header_SetHotDivider(hwnd, fPos, dw) \
	//         (int)SNDMSG((hwnd), HDM_SETHOTDIVIDER, (WPARAM)(fPos), (LPARAM)(dw))
	// // convenience message for external dragdrop
	// wParam = BOOL  specifying whether the lParam is a dwPos of the cursor
	// //              position or the index of which divider to hotlight
	// lParam = depends on wParam  (-1 and wParm = FALSE turns off hotlight)

	HDM_SETBITMAPMARGIN = (HDM_FIRST + 20)
	// #define Header_SetBitmapMargin(hwnd, iWidth) \
	//         (int)SNDMSG((hwnd), HDM_SETBITMAPMARGIN, (WPARAM)(iWidth), 0)

	HDM_GETBITMAPMARGIN = (HDM_FIRST + 21)
	// #define Header_GetBitmapMargin(hwnd) \
	//         (int)SNDMSG((hwnd), HDM_GETBITMAPMARGIN, 0, 0)

	HDM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	// #define Header_SetUnicodeFormat(hwnd, fUnicode)  \
	//     (BOOL)SNDMSG((hwnd), HDM_SETUNICODEFORMAT, (WPARAM)(fUnicode), 0)

	HDM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	// #define Header_GetUnicodeFormat(hwnd)  \
	//     (BOOL)SNDMSG((hwnd), HDM_GETUNICODEFORMAT, 0, 0)

	HDM_SETFILTERCHANGETIMEOUT = (HDM_FIRST + 22)
	// #define Header_SetFilterChangeTimeout(hwnd, i) \
	//         (int)SNDMSG((hwnd), HDM_SETFILTERCHANGETIMEOUT, 0, (LPARAM)(i))

	HDM_EDITFILTER = (HDM_FIRST + 23)
	// #define Header_EditFilter(hwnd, i, fDiscardChanges) \
	//         (int)SNDMSG((hwnd), HDM_EDITFILTER, (WPARAM)(i), MAKELPARAM(fDiscardChanges, 0))

	// // Clear filter takes -1 as a column value to indicate that all
	// // the filter should be cleared.  When this happens you will
	// // only receive a single filter changed notification.

	HDM_CLEARFILTER = (HDM_FIRST + 24)
	// #define Header_ClearFilter(hwnd, i) \
	//         (int)SNDMSG((hwnd), HDM_CLEARFILTER, (WPARAM)(i), 0)
	// #define Header_ClearAllFilters(hwnd) \
	//         (int)SNDMSG((hwnd), HDM_CLEARFILTER, (WPARAM)-1, 0)

	// #if (_WIN32_IE >= 0x0600)
	//HDM_TRANSLATEACCELERATOR = CCM_TRANSLATEACCELERATOR
	// #endif

	// #if(NTDDI_VERSION >= NTDDI_VISTA)

	HDM_GETITEMDROPDOWNRECT = (HDM_FIRST + 25) // rect of item's drop down button
	// #define Header_GetItemDropDownRect(hwnd, iItem, lprc) \
	//         (BOOL)SNDMSG((hwnd), HDM_GETITEMDROPDOWNRECT, (WPARAM)(iItem), (LPARAM)(lprc))

	HDM_GETOVERFLOWRECT = (HDM_FIRST + 26) // rect of overflow button
	// #define Header_GetOverflowRect(hwnd, lprc) \
	//         (BOOL)SNDMSG((hwnd), HDM_GETOVERFLOWRECT, 0, (LPARAM)(lprc))

	HDM_GETFOCUSEDITEM = (HDM_FIRST + 27)
	// #define Header_GetFocusedItem(hwnd) \
	//         (int)SNDMSG((hwnd), HDM_GETFOCUSEDITEM, (WPARAM)(0), (LPARAM)(0))

	HDM_SETFOCUSEDITEM = (HDM_FIRST + 28)
	// #define Header_SetFocusedItem(hwnd, iItem) \
	//         (BOOL)SNDMSG((hwnd), HDM_SETFOCUSEDITEM, (WPARAM)(0), (LPARAM)(iItem))

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	HDN_ITEMCHANGINGA    = (HDN_FIRST - 0)
	HDN_ITEMCHANGINGW    = (HDN_FIRST - 20)
	HDN_ITEMCHANGEDA     = (HDN_FIRST - 1)
	HDN_ITEMCHANGEDW     = (HDN_FIRST - 21)
	HDN_ITEMCLICKA       = (HDN_FIRST - 2)
	HDN_ITEMCLICKW       = (HDN_FIRST - 22)
	HDN_ITEMDBLCLICKA    = (HDN_FIRST - 3)
	HDN_ITEMDBLCLICKW    = (HDN_FIRST - 23)
	HDN_DIVIDERDBLCLICKA = (HDN_FIRST - 5)
	HDN_DIVIDERDBLCLICKW = (HDN_FIRST - 25)
	HDN_BEGINTRACKA      = (HDN_FIRST - 6)
	HDN_BEGINTRACKW      = (HDN_FIRST - 26)
	HDN_ENDTRACKA        = (HDN_FIRST - 7)
	HDN_ENDTRACKW        = (HDN_FIRST - 27)
	HDN_TRACKA           = (HDN_FIRST - 8)
	HDN_TRACKW           = (HDN_FIRST - 28)
	HDN_GETDISPINFOA     = (HDN_FIRST - 9)
	HDN_GETDISPINFOW     = (HDN_FIRST - 29)
	HDN_BEGINDRAG        = (HDN_FIRST - 10)
	HDN_ENDDRAG          = (HDN_FIRST - 11)
	HDN_FILTERCHANGE     = (HDN_FIRST - 12)
	HDN_FILTERBTNCLICK   = (HDN_FIRST - 13)

	// #if (_WIN32_IE >= 0x0600)
	HDN_BEGINFILTEREDIT = (HDN_FIRST - 14)
	HDN_ENDFILTEREDIT   = (HDN_FIRST - 15)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	HDN_ITEMSTATEICONCLICK = (HDN_FIRST - 16)
	HDN_ITEMKEYDOWN        = (HDN_FIRST - 17)
	HDN_DROPDOWN           = (HDN_FIRST - 18)
	HDN_OVERFLOWCLICK      = (HDN_FIRST - 19)
	// #endif

	// #ifdef UNICODE
	//HDN_ITEMCHANGING = HDN_ITEMCHANGINGW
	//HDN_ITEMCHANGED = HDN_ITEMCHANGEDW
	//HDN_ITEMCLICK = HDN_ITEMCLICKW
	//HDN_ITEMDBLCLICK = HDN_ITEMDBLCLICKW
	//HDN_DIVIDERDBLCLICK = HDN_DIVIDERDBLCLICKW
	//HDN_BEGINTRACK = HDN_BEGINTRACKW
	//HDN_ENDTRACK = HDN_ENDTRACKW
	//HDN_TRACK = HDN_TRACKW
	//HDN_GETDISPINFO = HDN_GETDISPINFOW
	// #else
	//HDN_ITEMCHANGING = HDN_ITEMCHANGINGA
	//HDN_ITEMCHANGED = HDN_ITEMCHANGEDA
	//HDN_ITEMCLICK = HDN_ITEMCLICKA
	//HDN_ITEMDBLCLICK = HDN_ITEMDBLCLICKA
	//HDN_DIVIDERDBLCLICK = HDN_DIVIDERDBLCLICKA
	//HDN_BEGINTRACK = HDN_BEGINTRACKA
	//HDN_ENDTRACK = HDN_ENDTRACKA
	//HDN_TRACK = HDN_TRACKA
	//HDN_GETDISPINFO = HDN_GETDISPINFOA
	// #endif
	//
	//HD_NOTIFYA = NMHEADERA
	//HD_NOTIFYW = NMHEADERW
	//HD_NOTIFY = NMHEADER

	// typedef struct tagNMHEADERA
	// {
	//     NMHDR   hdr;
	//     int     iItem;
	//     int     iButton;
	//     HDITEMA *pitem;
	// }  NMHEADERA, *LPNMHEADERA;

	// typedef struct tagNMHEADERW
	// {
	//     NMHDR   hdr;
	//     int     iItem;
	//     int     iButton;
	//     HDITEMW *pitem;
	// } NMHEADERW, *LPNMHEADERW;

	// #ifdef UNICODE
	//NMHEADER = NMHEADERW
	//LPNMHEADER = LPNMHEADERW
	// #else
	//NMHEADER = NMHEADERA
	//LPNMHEADER = LPNMHEADERA
	// #endif

	// typedef struct tagNMHDDISPINFOW
	// {
	//     NMHDR   hdr;
	//     int     iItem;
	//     UINT    mask;
	//     LPWSTR  pszText;
	//     int     cchTextMax;
	//     int     iImage;
	//     LPARAM  lParam;
	// } NMHDDISPINFOW, *LPNMHDDISPINFOW;

	// typedef struct tagNMHDDISPINFOA
	// {
	//     NMHDR   hdr;
	//     int     iItem;
	//     UINT    mask;
	//     LPSTR   pszText;
	//     int     cchTextMax;
	//     int     iImage;
	//     LPARAM  lParam;
	// } NMHDDISPINFOA, *LPNMHDDISPINFOA;

	// #ifdef UNICODE
	//NMHDDISPINFO = NMHDDISPINFOW
	//LPNMHDDISPINFO = LPNMHDDISPINFOW
	// #else
	//NMHDDISPINFO = NMHDDISPINFOA
	//LPNMHDDISPINFO = LPNMHDDISPINFOA
	// #endif

	// typedef struct tagNMHDFILTERBTNCLICK
	// {
	//     NMHDR hdr;
	//     INT iItem;
	//     RECT rc;
	// } NMHDFILTERBTNCLICK, *LPNMHDFILTERBTNCLICK;

	// #endif      // NOHEADER

	//====== TOOLBAR CONTROL ======================================================

	// #ifndef NOTOOLBAR

	// #ifdef _WIN32
	TOOLBARCLASSNAMEW = "ToolbarWindow32"
	// #define TOOLBARCLASSNAMEA       "ToolbarWindow32"

	// #ifdef  UNICODE
	TOOLBARCLASSNAME = TOOLBARCLASSNAMEW
	// #else
	//TOOLBARCLASSNAME = TOOLBARCLASSNAMEA
	// #endif

	// #else
	// #define TOOLBARCLASSNAME        "ToolbarWindow"
	// #endif

	// typedef struct _TBBUTTON {
	//     int iBitmap;
	//     int idCommand;
	//     BYTE fsState;
	//     BYTE fsStyle;
	// #ifdef _WIN64
	//     BYTE bReserved[6];          // padding for alignment
	// #elif defined(_WIN32)
	//     BYTE bReserved[2];          // padding for alignment
	// #endif
	//     DWORD_PTR dwData;
	//     INT_PTR iString;
	// } TBBUTTON, NEAR* PTBBUTTON, *LPTBBUTTON;
	// typedef const TBBUTTON *LPCTBBUTTON;

	// typedef struct _COLORMAP {
	//     COLORREF from;
	//     COLORREF to;
	// } COLORMAP, *LPCOLORMAP;

	// WINCOMMCTRLAPI HWND WINAPI CreateToolbarEx(HWND hwnd, DWORD ws, UINT wID, int nBitmaps,
	//                         HINSTANCE hBMInst, UINT_PTR wBMID, LPCTBBUTTON lpButtons,
	//                         int iNumButtons, int dxButton, int dyButton,
	//                         int dxBitmap, int dyBitmap, UINT uStructSize);

	// WINCOMMCTRLAPI HBITMAP WINAPI CreateMappedBitmap(HINSTANCE hInstance, INT_PTR idBitmap,
	//                                   UINT wFlags, _In_opt_ LPCOLORMAP lpColorMap,
	//                                   int iNumMaps);

	CMB_MASKED            = 0
	TBSTATE_CHECKED       = 0
	TBSTATE_PRESSED       = 0
	TBSTATE_ENABLED       = 0
	TBSTATE_HIDDEN        = 0
	TBSTATE_INDETERMINATE = 0
	TBSTATE_WRAP          = 0
	TBSTATE_ELLIPSES      = 0
	TBSTATE_MARKED        = 0

	// // begin_r_commctrl

	TBSTYLE_BUTTON     = 0                               // obsolete; use BTNS_BUTTON instead
	TBSTYLE_SEP        = 0                               // obsolete; use BTNS_SEP instead
	TBSTYLE_CHECK      = 0                               // obsolete; use BTNS_CHECK instead
	TBSTYLE_GROUP      = 0                               // obsolete; use BTNS_GROUP instead
	TBSTYLE_CHECKGROUP = (TBSTYLE_GROUP | TBSTYLE_CHECK) // obsolete; use BTNS_CHECKGROUP instead
	TBSTYLE_DROPDOWN   = 0                               // obsolete; use BTNS_DROPDOWN instead
	TBSTYLE_AUTOSIZE   = 0                               // obsolete; use BTNS_AUTOSIZE instead
	TBSTYLE_NOPREFIX   = 0                               // obsolete; use BTNS_NOPREFIX instead

	TBSTYLE_TOOLTIPS     = 0
	TBSTYLE_WRAPABLE     = 0
	TBSTYLE_ALTDRAG      = 0
	TBSTYLE_FLAT         = 0
	TBSTYLE_LIST         = 0
	TBSTYLE_CUSTOMERASE  = 0
	TBSTYLE_REGISTERDROP = 0
	TBSTYLE_TRANSPARENT  = 0

	// // end_r_commctrl

	TBSTYLE_EX_DRAWDDARROWS = 0

	// // begin_r_commctrl

	BTNS_BUTTON        = TBSTYLE_BUTTON     // 0x0000
	BTNS_SEP           = TBSTYLE_SEP        // 0x0001
	BTNS_CHECK         = TBSTYLE_CHECK      // 0x0002
	BTNS_GROUP         = TBSTYLE_GROUP      // 0x0004
	BTNS_CHECKGROUP    = TBSTYLE_CHECKGROUP // (TBSTYLE_GROUP | TBSTYLE_CHECK)
	BTNS_DROPDOWN      = TBSTYLE_DROPDOWN   // 0x0008
	BTNS_AUTOSIZE      = TBSTYLE_AUTOSIZE   // 0x0010; automatically calculate the cx of the button
	BTNS_NOPREFIX      = TBSTYLE_NOPREFIX   // 0x0020; this button should not have accel prefix
	BTNS_SHOWTEXT      = 0                  // ignored unless TBSTYLE_EX_MIXEDBUTTONS is set
	BTNS_WHOLEDROPDOWN = 0                  // draw drop-down arrow, but without split arrow section

	// // end_r_commctrl

	TBSTYLE_EX_MIXEDBUTTONS       = 0
	TBSTYLE_EX_HIDECLIPPEDBUTTONS = 0 // don't show partially obscured buttons

	TBSTYLE_EX_MULTICOLUMN = 0 // conflicts w/ TBSTYLE_WRAPABLE
	TBSTYLE_EX_VERTICAL    = 0

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TBSTYLE_EX_DOUBLEBUFFER = 0 // Double Buffer the toolbar
	// #endif

	// // Custom Draw Structure
	// typedef struct _NMTBCUSTOMDRAW {
	//     NMCUSTOMDRAW nmcd;
	//     HBRUSH hbrMonoDither;
	//     HBRUSH hbrLines;                // For drawing lines on buttons
	//     HPEN hpenLines;                 // For drawing lines on buttons

	//     COLORREF clrText;               // Color of text
	//     COLORREF clrMark;               // Color of text bk when marked. (only if TBSTATE_MARKED)
	//     COLORREF clrTextHighlight;      // Color of text when highlighted
	//     COLORREF clrBtnFace;            // Background of the button
	//     COLORREF clrBtnHighlight;       // 3D highlight
	//     COLORREF clrHighlightHotTrack;  // In conjunction with fHighlightHotTrack
	//                                     // will cause button to highlight like a menu
	//     RECT rcText;                    // Rect for text

	//     int nStringBkMode;
	//     int nHLStringBkMode;
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	//     int iListGap;
	// #endif
	// } NMTBCUSTOMDRAW, * LPNMTBCUSTOMDRAW;

	// // Toolbar custom draw return flags
	TBCDRF_NOEDGES        = 0 // Don't draw button edges
	TBCDRF_HILITEHOTTRACK = 0 // Use color of the button bk when hottracked
	TBCDRF_NOOFFSET       = 0 // Don't offset button if pressed
	TBCDRF_NOMARK         = 0 // Don't draw default highlight of image/text for TBSTATE_MARKED
	TBCDRF_NOETCHEDEFFECT = 0 // Don't draw etched effect for disabled items

	TBCDRF_BLENDICON    = 0 // Use ILD_BLEND50 on the icon image
	TBCDRF_NOBACKGROUND = 0 // Use ILD_BLEND50 on the icon image
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TBCDRF_USECDCOLORS = 0 // Use CustomDrawColors to RenderText regardless of VisualStyle
	// #endif

	TB_ENABLEBUTTON          = (WM_USER + 1)
	TB_CHECKBUTTON           = (WM_USER + 2)
	TB_PRESSBUTTON           = (WM_USER + 3)
	TB_HIDEBUTTON            = (WM_USER + 4)
	TB_INDETERMINATE         = (WM_USER + 5)
	TB_MARKBUTTON            = (WM_USER + 6)
	TB_ISBUTTONENABLED       = (WM_USER + 9)
	TB_ISBUTTONCHECKED       = (WM_USER + 10)
	TB_ISBUTTONPRESSED       = (WM_USER + 11)
	TB_ISBUTTONHIDDEN        = (WM_USER + 12)
	TB_ISBUTTONINDETERMINATE = (WM_USER + 13)
	TB_ISBUTTONHIGHLIGHTED   = (WM_USER + 14)
	TB_SETSTATE              = (WM_USER + 17)
	TB_GETSTATE              = (WM_USER + 18)
	TB_ADDBITMAP             = (WM_USER + 19)

	// #ifdef _WIN32
	// typedef struct tagTBADDBITMAP {
	//         HINSTANCE       hInst;
	//         UINT_PTR        nID;
	// } TBADDBITMAP, *LPTBADDBITMAP;

	//HINST_COMMCTRL = ((HINSTANCE)-1)
	IDB_STD_SMALL_COLOR  = 0
	IDB_STD_LARGE_COLOR  = 1
	IDB_VIEW_SMALL_COLOR = 4
	IDB_VIEW_LARGE_COLOR = 5
	IDB_HIST_SMALL_COLOR = 8
	IDB_HIST_LARGE_COLOR = 9
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	IDB_HIST_NORMAL   = 12
	IDB_HIST_HOT      = 13
	IDB_HIST_DISABLED = 14
	IDB_HIST_PRESSED  = 15
	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// // icon indexes for standard bitmap

	STD_CUT        = 0
	STD_COPY       = 1
	STD_PASTE      = 2
	STD_UNDO       = 3
	STD_REDOW      = 4
	STD_DELETE     = 5
	STD_FILENEW    = 6
	STD_FILEOPEN   = 7
	STD_FILESAVE   = 8
	STD_PRINTPRE   = 9
	STD_PROPERTIES = 10
	STD_HELP       = 11
	STD_FIND       = 12
	STD_REPLACE    = 13
	STD_PRINT      = 14

	// // icon indexes for standard view bitmap

	VIEW_LARGEICONS    = 0
	VIEW_SMALLICONS    = 1
	VIEW_LIST          = 2
	VIEW_DETAILS       = 3
	VIEW_SORTNAME      = 4
	VIEW_SORTSIZE      = 5
	VIEW_SORTDATE      = 6
	VIEW_SORTTYPE      = 7
	VIEW_PARENTFOLDER  = 8
	VIEW_NETCONNECT    = 9
	VIEW_NETDISCONNECT = 10
	VIEW_NEWFOLDER     = 11
	VIEW_VIEWMENU      = 12

	HIST_BACK           = 0
	HIST_FORWARD        = 1
	HIST_FAVORITES      = 2
	HIST_ADDTOFAVORITES = 3
	HIST_VIEWTREE       = 4

	// #endif

	TB_ADDBUTTONSA   = (WM_USER + 20)
	TB_INSERTBUTTONA = (WM_USER + 21)

	TB_DELETEBUTTON   = (WM_USER + 22)
	TB_GETBUTTON      = (WM_USER + 23)
	TB_BUTTONCOUNT    = (WM_USER + 24)
	TB_COMMANDTOINDEX = (WM_USER + 25)

	// #ifdef _WIN32

	// typedef struct tagTBSAVEPARAMSA {
	//     HKEY hkr;
	//     LPCSTR pszSubKey;
	//     LPCSTR pszValueName;
	// } TBSAVEPARAMSA, *LPTBSAVEPARAMSA;

	// typedef struct tagTBSAVEPARAMSW {
	//     HKEY hkr;
	//     LPCWSTR pszSubKey;
	//     LPCWSTR pszValueName;
	// } TBSAVEPARAMSW, *LPTBSAVEPARAMW;

	// #ifdef UNICODE
	//TBSAVEPARAMS = TBSAVEPARAMSW
	//LPTBSAVEPARAMS = LPTBSAVEPARAMSW
	// #else
	//TBSAVEPARAMS = TBSAVEPARAMSA
	//LPTBSAVEPARAMS = LPTBSAVEPARAMSA
	// #endif

	// #endif  // _WIN32

	TB_SAVERESTOREA         = (WM_USER + 26)
	TB_SAVERESTOREW         = (WM_USER + 76)
	TB_CUSTOMIZE            = (WM_USER + 27)
	TB_ADDSTRINGA           = (WM_USER + 28)
	TB_ADDSTRINGW           = (WM_USER + 77)
	TB_GETITEMRECT          = (WM_USER + 29)
	TB_BUTTONSTRUCTSIZE     = (WM_USER + 30)
	TB_SETBUTTONSIZE        = (WM_USER + 31)
	TB_SETBITMAPSIZE        = (WM_USER + 32)
	TB_AUTOSIZE             = (WM_USER + 33)
	TB_GETTOOLTIPS          = (WM_USER + 35)
	TB_SETTOOLTIPS          = (WM_USER + 36)
	TB_SETPARENT            = (WM_USER + 37)
	TB_SETROWS              = (WM_USER + 39)
	TB_GETROWS              = (WM_USER + 40)
	TB_SETCMDID             = (WM_USER + 42)
	TB_CHANGEBITMAP         = (WM_USER + 43)
	TB_GETBITMAP            = (WM_USER + 44)
	TB_GETBUTTONTEXTA       = (WM_USER + 45)
	TB_GETBUTTONTEXTW       = (WM_USER + 75)
	TB_REPLACEBITMAP        = (WM_USER + 46)
	TB_SETINDENT            = (WM_USER + 47)
	TB_SETIMAGELIST         = (WM_USER + 48)
	TB_GETIMAGELIST         = (WM_USER + 49)
	TB_LOADIMAGES           = (WM_USER + 50)
	TB_GETRECT              = (WM_USER + 51) // wParam is the Cmd instead of index
	TB_SETHOTIMAGELIST      = (WM_USER + 52)
	TB_GETHOTIMAGELIST      = (WM_USER + 53)
	TB_SETDISABLEDIMAGELIST = (WM_USER + 54)
	TB_GETDISABLEDIMAGELIST = (WM_USER + 55)
	TB_SETSTYLE             = (WM_USER + 56)
	TB_GETSTYLE             = (WM_USER + 57)
	TB_GETBUTTONSIZE        = (WM_USER + 58)
	TB_SETBUTTONWIDTH       = (WM_USER + 59)
	TB_SETMAXTEXTROWS       = (WM_USER + 60)
	TB_GETTEXTROWS          = (WM_USER + 61)

	// #ifdef UNICODE
	//TB_GETBUTTONTEXT = TB_GETBUTTONTEXTW
	//TB_SAVERESTORE = TB_SAVERESTOREW
	//TB_ADDSTRING = TB_ADDSTRINGW
	// #else
	//TB_GETBUTTONTEXT = TB_GETBUTTONTEXTA
	//TB_SAVERESTORE = TB_SAVERESTOREA
	//TB_ADDSTRING = TB_ADDSTRINGA
	// #endif
	TB_GETOBJECT          = (WM_USER + 62) // wParam == IID, lParam void **ppv
	TB_GETHOTITEM         = (WM_USER + 71)
	TB_SETHOTITEM         = (WM_USER + 72) // wParam == iHotItem
	TB_SETANCHORHIGHLIGHT = (WM_USER + 73) // wParam == TRUE/FALSE
	TB_GETANCHORHIGHLIGHT = (WM_USER + 74)
	TB_MAPACCELERATORA    = (WM_USER + 78) // wParam == ch, lParam int * pidBtn

	// typedef struct {
	//     int   iButton;
	//     DWORD dwFlags;
	// } TBINSERTMARK, * LPTBINSERTMARK;
	TBIMHT_AFTER      = 0 // TRUE = insert After iButton, otherwise before
	TBIMHT_BACKGROUND = 0 // TRUE iff missed buttons completely

	TB_GETINSERTMARK      = (WM_USER + 79) // lParam == LPTBINSERTMARK
	TB_SETINSERTMARK      = (WM_USER + 80) // lParam == LPTBINSERTMARK
	TB_INSERTMARKHITTEST  = (WM_USER + 81) // wParam == LPPOINT lParam == LPTBINSERTMARK
	TB_MOVEBUTTON         = (WM_USER + 82)
	TB_GETMAXSIZE         = (WM_USER + 83) // lParam == LPSIZE
	TB_SETEXTENDEDSTYLE   = (WM_USER + 84) // For TBSTYLE_EX_*
	TB_GETEXTENDEDSTYLE   = (WM_USER + 85) // For TBSTYLE_EX_*
	TB_GETPADDING         = (WM_USER + 86)
	TB_SETPADDING         = (WM_USER + 87)
	TB_SETINSERTMARKCOLOR = (WM_USER + 88)
	TB_GETINSERTMARKCOLOR = (WM_USER + 89)

	TB_SETCOLORSCHEME = CCM_SETCOLORSCHEME // lParam is color scheme
	TB_GETCOLORSCHEME = CCM_GETCOLORSCHEME // fills in COLORSCHEME pointed to by lParam

	TB_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	TB_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT

	TB_MAPACCELERATORW = (WM_USER + 90) // wParam == ch, lParam int * pidBtn
	// #ifdef UNICODE
	//TB_MAPACCELERATOR = TB_MAPACCELERATORW
	// #else
	//TB_MAPACCELERATOR = TB_MAPACCELERATORA
	// #endif

	// typedef struct {
	//     HINSTANCE       hInstOld;
	//     UINT_PTR        nIDOld;
	//     HINSTANCE       hInstNew;
	//     UINT_PTR        nIDNew;
	//     int             nButtons;
	// } TBREPLACEBITMAP, *LPTBREPLACEBITMAP;

	// #ifdef _WIN32

	TBBF_LARGE = 0

	TB_GETBITMAPFLAGS = (WM_USER + 41)

	TBIF_IMAGE   = 0
	TBIF_TEXT    = 0
	TBIF_STATE   = 0
	TBIF_STYLE   = 0
	TBIF_LPARAM  = 0
	TBIF_COMMAND = 0
	TBIF_SIZE    = 0

	TBIF_BYINDEX = 0 // this specifies that the wparam in Get/SetButtonInfo is an index, not id

	// typedef struct {
	//     UINT cbSize;
	//     DWORD dwMask;
	//     int idCommand;
	//     int iImage;
	//     BYTE fsState;
	//     BYTE fsStyle;
	//     WORD cx;
	//     DWORD_PTR lParam;
	//     LPSTR pszText;
	//     int cchText;
	// } TBBUTTONINFOA, *LPTBBUTTONINFOA;

	// typedef struct {
	//     UINT cbSize;
	//     DWORD dwMask;
	//     int idCommand;
	//     int iImage;
	//     BYTE fsState;
	//     BYTE fsStyle;
	//     WORD cx;
	//     DWORD_PTR lParam;
	//     LPWSTR pszText;
	//     int cchText;
	// } TBBUTTONINFOW, *LPTBBUTTONINFOW;

	// #ifdef UNICODE
	//TBBUTTONINFO = TBBUTTONINFOW
	//LPTBBUTTONINFO = LPTBBUTTONINFOW
	// #else
	//TBBUTTONINFO = TBBUTTONINFOA
	//LPTBBUTTONINFO = LPTBBUTTONINFOA
	// #endif

	// // BUTTONINFO APIs do NOT support the string pool.
	TB_GETBUTTONINFOW = (WM_USER + 63)
	TB_SETBUTTONINFOW = (WM_USER + 64)
	TB_GETBUTTONINFOA = (WM_USER + 65)
	TB_SETBUTTONINFOA = (WM_USER + 66)
	// #ifdef UNICODE
	//TB_GETBUTTONINFO = TB_GETBUTTONINFOW
	//TB_SETBUTTONINFO = TB_SETBUTTONINFOW
	// #else
	//TB_GETBUTTONINFO = TB_GETBUTTONINFOA
	//TB_SETBUTTONINFO = TB_SETBUTTONINFOA
	// #endif

	TB_INSERTBUTTONW = (WM_USER + 67)
	TB_ADDBUTTONSW   = (WM_USER + 68)

	TB_HITTEST = (WM_USER + 69)

	// // New post Win95/NT4 for InsertButton and AddButton.  if iString member
	// // is a pointer to a string, it will be handled as a string like listview
	// // (although LPSTR_TEXTCALLBACK is not supported).
	// #ifdef UNICODE
	//TB_INSERTBUTTON = TB_INSERTBUTTONW
	//TB_ADDBUTTONS = TB_ADDBUTTONSW
	// #else
	//TB_INSERTBUTTON = TB_INSERTBUTTONA
	//TB_ADDBUTTONS = TB_ADDBUTTONSA
	// #endif

	TB_SETDRAWTEXTFLAGS = (WM_USER + 70) // wParam == mask lParam == bit values

	TB_GETSTRINGW = (WM_USER + 91)
	TB_GETSTRINGA = (WM_USER + 92)
	// #ifdef UNICODE
	//TB_GETSTRING = TB_GETSTRINGW
	// #else
	//TB_GETSTRING = TB_GETSTRINGA
	// #endif

	TB_SETBOUNDINGSIZE   = (WM_USER + 93)
	TB_SETHOTITEM2       = (WM_USER + 94) // wParam == iHotItem,  lParam = dwFlags
	TB_HASACCELERATOR    = (WM_USER + 95) // wParam == char, lParam = &iCount
	TB_SETLISTGAP        = (WM_USER + 96)
	TB_GETIMAGELISTCOUNT = (WM_USER + 98)
	TB_GETIDEALSIZE      = (WM_USER + 99) // wParam == fHeight, lParam = psize
	// // before using WM_USER + 103, recycle old space above (WM_USER + 97)
	//TB_TRANSLATEACCELERATOR = CCM_TRANSLATEACCELERATOR

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TBMF_PAD           = 0
	TBMF_BARPAD        = 0
	TBMF_BUTTONSPACING = 0

	// typedef struct {
	//     UINT cbSize;
	//     DWORD dwMask;

	//     int cxPad;        // PAD
	//     int cyPad;
	//     int cxBarPad;     // BARPAD
	//     int cyBarPad;
	//     int cxButtonSpacing;   // BUTTONSPACING
	//     int cyButtonSpacing;
	// } TBMETRICS, * LPTBMETRICS;

	TB_GETMETRICS = (WM_USER + 101)
	TB_SETMETRICS = (WM_USER + 102)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TB_GETITEMDROPDOWNRECT = (WM_USER + 103) // Rect of item's drop down button
	TB_SETPRESSEDIMAGELIST = (WM_USER + 104)
	TB_GETPRESSEDIMAGELIST = (WM_USER + 105)
	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TB_SETWINDOWTHEME = CCM_SETWINDOWTHEME
	// #endif

	TBN_GETBUTTONINFOA = (TBN_FIRST - 0)
	TBN_BEGINDRAG      = (TBN_FIRST - 1)
	TBN_ENDDRAG        = (TBN_FIRST - 2)
	TBN_BEGINADJUST    = (TBN_FIRST - 3)
	TBN_ENDADJUST      = (TBN_FIRST - 4)
	TBN_RESET          = (TBN_FIRST - 5)
	TBN_QUERYINSERT    = (TBN_FIRST - 6)
	TBN_QUERYDELETE    = (TBN_FIRST - 7)
	TBN_TOOLBARCHANGE  = (TBN_FIRST - 8)
	TBN_CUSTHELP       = (TBN_FIRST - 9)
	TBN_DROPDOWN       = (TBN_FIRST - 10)
	TBN_GETOBJECT      = (TBN_FIRST - 12)

	// // Structure for TBN_HOTITEMCHANGE notification
	// //
	// typedef struct tagNMTBHOTITEM
	// {
	//     NMHDR   hdr;
	//     int     idOld;
	//     int     idNew;
	//     DWORD   dwFlags;           // HICF_*
	// } NMTBHOTITEM, * LPNMTBHOTITEM;

	// // Hot item change flags
	HICF_OTHER          = 0
	HICF_MOUSE          = 0 // Triggered by mouse
	HICF_ARROWKEYS      = 0 // Triggered by arrow keys
	HICF_ACCELERATOR    = 0 // Triggered by accelerator
	HICF_DUPACCEL       = 0 // This accelerator is not unique
	HICF_ENTERING       = 0 // idOld is invalid
	HICF_LEAVING        = 0 // idNew is invalid
	HICF_RESELECT       = 0 // hot item reselected
	HICF_LMOUSE         = 0 // left mouse button selected
	HICF_TOGGLEDROPDOWN = 0 // Toggle button's dropdown state

	TBN_HOTITEMCHANGE   = (TBN_FIRST - 13)
	TBN_DRAGOUT         = (TBN_FIRST - 14) // this is sent when the user clicks down on a button then drags off the button
	TBN_DELETINGBUTTON  = (TBN_FIRST - 15) // uses TBNOTIFY
	TBN_GETDISPINFOA    = (TBN_FIRST - 16) // This is sent when the  toolbar needs  some display information
	TBN_GETDISPINFOW    = (TBN_FIRST - 17) // This is sent when the  toolbar needs  some display information
	TBN_GETINFOTIPA     = (TBN_FIRST - 18)
	TBN_GETINFOTIPW     = (TBN_FIRST - 19)
	TBN_GETBUTTONINFOW  = (TBN_FIRST - 20)
	TBN_RESTORE         = (TBN_FIRST - 21)
	TBN_SAVE            = (TBN_FIRST - 22)
	TBN_INITCUSTOMIZE   = (TBN_FIRST - 23)
	TBNRF_HIDEHELP      = 0
	TBNRF_ENDCUSTOMIZE  = 0
	TBN_WRAPHOTITEM     = (TBN_FIRST - 24)
	TBN_DUPACCELERATOR  = (TBN_FIRST - 25)
	TBN_WRAPACCELERATOR = (TBN_FIRST - 26)
	TBN_DRAGOVER        = (TBN_FIRST - 27)
	TBN_MAPACCELERATOR  = (TBN_FIRST - 28)

	// typedef struct tagNMTBSAVE
	// {
	//     NMHDR hdr;
	//     DWORD* pData;
	//     DWORD* pCurrent;
	//     UINT cbData;
	//     int iItem;
	//     int cButtons;
	//     TBBUTTON tbButton;
	// } NMTBSAVE, *LPNMTBSAVE;

	// typedef struct tagNMTBRESTORE
	// {
	//     NMHDR hdr;
	//     DWORD* pData;
	//     DWORD* pCurrent;
	//     UINT cbData;
	//     int iItem;
	//     int cButtons;
	//     int cbBytesPerRecord;
	//     TBBUTTON tbButton;
	// } NMTBRESTORE, *LPNMTBRESTORE;

	// typedef struct tagNMTBGETINFOTIPA
	// {
	//     NMHDR hdr;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iItem;
	//     LPARAM lParam;
	// } NMTBGETINFOTIPA, *LPNMTBGETINFOTIPA;

	// typedef struct tagNMTBGETINFOTIPW
	// {
	//     NMHDR hdr;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iItem;
	//     LPARAM lParam;
	// } NMTBGETINFOTIPW, *LPNMTBGETINFOTIPW;

	// #ifdef UNICODE
	//TBN_GETINFOTIP = TBN_GETINFOTIPW
	//NMTBGETINFOTIP = NMTBGETINFOTIPW
	//LPNMTBGETINFOTIP = LPNMTBGETINFOTIPW
	// #else
	//TBN_GETINFOTIP = TBN_GETINFOTIPA
	//NMTBGETINFOTIP = NMTBGETINFOTIPA
	//LPNMTBGETINFOTIP = LPNMTBGETINFOTIPA
	// #endif

	TBNF_IMAGE      = 0
	TBNF_TEXT       = 0
	TBNF_DI_SETITEM = 0

	// typedef struct {
	//     NMHDR  hdr;
	//     DWORD dwMask;     // [in] Specifies the values requested .[out] Client ask the data to be set for future use
	//     int idCommand;    // [in] id of button we're requesting info for
	//     DWORD_PTR lParam;  // [in] lParam of button
	//     int iImage;       // [out] image index
	//     LPSTR pszText;    // [out] new text for item
	//     int cchText;      // [in] size of buffer pointed to by pszText
	// } NMTBDISPINFOA, *LPNMTBDISPINFOA;

	// typedef struct {
	//     NMHDR hdr;
	//     DWORD dwMask;      //[in] Specifies the values requested .[out] Client ask the data to be set for future use
	//     int idCommand;    // [in] id of button we're requesting info for
	//     DWORD_PTR lParam;  // [in] lParam of button
	//     int iImage;       // [out] image index
	//     LPWSTR pszText;   // [out] new text for item
	//     int cchText;      // [in] size of buffer pointed to by pszText
	// } NMTBDISPINFOW, *LPNMTBDISPINFOW;

	// #ifdef UNICODE
	//TBN_GETDISPINFO = TBN_GETDISPINFOW
	//NMTBDISPINFO = NMTBDISPINFOW
	//LPNMTBDISPINFO = LPNMTBDISPINFOW
	// #else
	//TBN_GETDISPINFO = TBN_GETDISPINFOA
	//NMTBDISPINFO = NMTBDISPINFOA
	//LPNMTBDISPINFO = LPNMTBDISPINFOA
	// #endif

	// // Return codes for TBN_DROPDOWN
	TBDDRET_DEFAULT      = 0
	TBDDRET_NODEFAULT    = 1
	TBDDRET_TREATPRESSED = 2 // Treat as a standard press button

	//
	//// #ifdef UNICODE
	//TBN_GETBUTTONINFO = TBN_GETBUTTONINFOW
	//// #else
	//TBN_GETBUTTONINFO = TBN_GETBUTTONINFOA
	//// #endif
	//
	//TBNOTIFYA = NMTOOLBARA
	//TBNOTIFYW = NMTOOLBARW
	//LPTBNOTIFYA = LPNMTOOLBARA
	//LPTBNOTIFYW = LPNMTOOLBARW
	//
	//TBNOTIFY = NMTOOLBAR
	//LPTBNOTIFY = LPNMTOOLBAR

	// typedef struct tagNMTOOLBARA {
	//     NMHDR   hdr;
	//     int     iItem;
	//     TBBUTTON tbButton;
	//     int     cchText;
	//     LPSTR   pszText;
	//     RECT    rcButton;
	// } NMTOOLBARA, *LPNMTOOLBARA;

	// typedef struct tagNMTOOLBARW {
	//     NMHDR   hdr;
	//     int     iItem;
	//     TBBUTTON tbButton;
	//     int     cchText;
	//     LPWSTR   pszText;
	//     RECT    rcButton;
	// } NMTOOLBARW, *LPNMTOOLBARW;

	//// #ifdef UNICODE
	//NMTOOLBAR = NMTOOLBARW
	//LPNMTOOLBAR = LPNMTOOLBARW
	//// #else
	//NMTOOLBAR = NMTOOLBARA
	//LPNMTOOLBAR = LPNMTOOLBARA
	//// #endif

	// #endif

	// #endif      // NOTOOLBAR

	//====== REBAR CONTROL ========================================================

	// #ifndef NOREBAR

	// #ifdef _WIN32
	REBARCLASSNAMEW = "ReBarWindow32"
	// #define REBARCLASSNAMEA         "ReBarWindow32"

	//// #ifdef  UNICODE
	//REBARCLASSNAME = REBARCLASSNAMEW
	//// #else
	//REBARCLASSNAME = REBARCLASSNAMEA
	//// #endif

	// #else
	// #define REBARCLASSNAME          "ReBarWindow"
	// #endif

	RBIM_IMAGELIST = 0

	// // begin_r_commctrl

	RBS_TOOLTIPS        = 0
	RBS_VARHEIGHT       = 0
	RBS_BANDBORDERS     = 0
	RBS_FIXEDORDER      = 0
	RBS_REGISTERDROP    = 0
	RBS_AUTOSIZE        = 0
	RBS_VERTICALGRIPPER = 0 // this always has the vertical gripper (default for horizontal mode)
	RBS_DBLCLKTOGGLE    = 0
	// // end_r_commctrl

	// typedef struct tagREBARINFO
	// {
	//     UINT        cbSize;
	//     UINT        fMask;
	// #ifndef NOIMAGEAPIS
	//     HIMAGELIST  himl;
	// #else
	//     HANDLE      himl;
	// #endif
	// }   REBARINFO, *LPREBARINFO;

	RBBS_BREAK          = 0 // break to new line
	RBBS_FIXEDSIZE      = 0 // band can't be sized
	RBBS_CHILDEDGE      = 0 // edge around top & bottom of child window
	RBBS_HIDDEN         = 0 // don't show
	RBBS_NOVERT         = 0 // don't show when vertical
	RBBS_FIXEDBMP       = 0 // bitmap doesn't move during band resize
	RBBS_VARIABLEHEIGHT = 0 // allow autosizing of this child vertically
	RBBS_GRIPPERALWAYS  = 0 // always show the gripper
	RBBS_NOGRIPPER      = 0 // never show the gripper
	RBBS_USECHEVRON     = 0 // display drop-down button for this band if it's sized smaller than ideal width
	RBBS_HIDETITLE      = 0 // keep band title hidden
	RBBS_TOPALIGN       = 0 // keep band in top row
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	// #endif

	RBBIM_STYLE      = 0
	RBBIM_COLORS     = 0
	RBBIM_TEXT       = 0
	RBBIM_IMAGE      = 0
	RBBIM_CHILD      = 0
	RBBIM_CHILDSIZE  = 0
	RBBIM_SIZE       = 0
	RBBIM_BACKGROUND = 0
	RBBIM_ID         = 0
	RBBIM_IDEALSIZE  = 0
	RBBIM_LPARAM     = 0
	RBBIM_HEADERSIZE = 0 // control the size of the header
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	RBBIM_CHEVRONLOCATION = 0
	RBBIM_CHEVRONSTATE    = 0
	// #endif

	// typedef struct tagREBARBANDINFOA
	// {
	//     UINT        cbSize;
	//     UINT        fMask;
	//     UINT        fStyle;
	//     COLORREF    clrFore;
	//     COLORREF    clrBack;
	//     LPSTR       lpText;
	//     UINT        cch;
	//     int         iImage;
	//     HWND        hwndChild;
	//     UINT        cxMinChild;
	//     UINT        cyMinChild;
	//     UINT        cx;
	//     HBITMAP     hbmBack;
	//     UINT        wID;
	//     UINT        cyChild;
	//     UINT        cyMaxChild;
	//     UINT        cyIntegral;
	//     UINT        cxIdeal;
	//     LPARAM      lParam;
	//     UINT        cxHeader;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     RECT        rcChevronLocation;  // the rect is in client co-ord wrt hwndChild
	//     UINT        uChevronState;      // STATE_SYSTEM_*
	// #endif
	// }   REBARBANDINFOA, *LPREBARBANDINFOA;
	// typedef REBARBANDINFOA CONST *LPCREBARBANDINFOA;

	//REBARBANDINFOA_V3_SIZE = CCSIZEOF_STRUCT(REBARBANDINFOA, wID)
	//REBARBANDINFOW_V3_SIZE = CCSIZEOF_STRUCT(REBARBANDINFOW, wID)
	//
	//REBARBANDINFOA_V6_SIZE = CCSIZEOF_STRUCT(REBARBANDINFOA, cxHeader)
	//REBARBANDINFOW_V6_SIZE = CCSIZEOF_STRUCT(REBARBANDINFOW, cxHeader)

	// typedef struct tagREBARBANDINFOW
	// {
	//     UINT        cbSize;
	//     UINT        fMask;
	//     UINT        fStyle;
	//     COLORREF    clrFore;
	//     COLORREF    clrBack;
	//     LPWSTR      lpText;
	//     UINT        cch;
	//     int         iImage;
	//     HWND        hwndChild;
	//     UINT        cxMinChild;
	//     UINT        cyMinChild;
	//     UINT        cx;
	//     HBITMAP     hbmBack;
	//     UINT        wID;
	//     UINT        cyChild;
	//     UINT        cyMaxChild;
	//     UINT        cyIntegral;
	//     UINT        cxIdeal;
	//     LPARAM      lParam;
	//     UINT        cxHeader;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     RECT        rcChevronLocation;    // the rect is in client co-ord wrt hwndChild
	//     UINT        uChevronState; // STATE_SYSTEM_*
	// #endif
	// }   REBARBANDINFOW, *LPREBARBANDINFOW;
	// typedef REBARBANDINFOW CONST *LPCREBARBANDINFOW;

	//// #ifdef UNICODE
	//REBARBANDINFO = REBARBANDINFOW
	//LPREBARBANDINFO = LPREBARBANDINFOW
	//LPCREBARBANDINFO = LPCREBARBANDINFOW
	//REBARBANDINFO_V3_SIZE = REBARBANDINFOW_V3_SIZE
	//REBARBANDINFO_V6_SIZE = REBARBANDINFOW_V6_SIZE
	//// #else
	//REBARBANDINFO = REBARBANDINFOA
	//LPREBARBANDINFO = LPREBARBANDINFOA
	//LPCREBARBANDINFO = LPCREBARBANDINFOA
	//REBARBANDINFO_V3_SIZE = REBARBANDINFOA_V3_SIZE
	//REBARBANDINFO_V6_SIZE = REBARBANDINFOA_V6_SIZE
	//// #endif

	RB_INSERTBANDA  = (WM_USER + 1)
	RB_DELETEBAND   = (WM_USER + 2)
	RB_GETBARINFO   = (WM_USER + 3)
	RB_SETBARINFO   = (WM_USER + 4)
	RB_SETBANDINFOA = (WM_USER + 6)
	RB_SETPARENT    = (WM_USER + 7)
	RB_HITTEST      = (WM_USER + 8)
	RB_GETRECT      = (WM_USER + 9)
	RB_INSERTBANDW  = (WM_USER + 10)
	RB_SETBANDINFOW = (WM_USER + 11)
	RB_GETBANDCOUNT = (WM_USER + 12)
	RB_GETROWCOUNT  = (WM_USER + 13)
	RB_GETROWHEIGHT = (WM_USER + 14)
	RB_IDTOINDEX    = (WM_USER + 16) // wParam == id
	RB_GETTOOLTIPS  = (WM_USER + 17)
	RB_SETTOOLTIPS  = (WM_USER + 18)
	RB_SETBKCOLOR   = (WM_USER + 19) // sets the default BK color
	RB_GETBKCOLOR   = (WM_USER + 20) // defaults to CLR_NONE
	RB_SETTEXTCOLOR = (WM_USER + 21)
	RB_GETTEXTCOLOR = (WM_USER + 22) // defaults to 0x00000000

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	RBSTR_CHANGERECT = 0 // flags for RB_SIZETORECT
	// #endif

	RB_SIZETORECT = (WM_USER + 23) // resize the rebar/break bands and such to this rect (lparam)

	RB_SETCOLORSCHEME = CCM_SETCOLORSCHEME // lParam is color scheme
	RB_GETCOLORSCHEME = CCM_GETCOLORSCHEME // fills in COLORSCHEME pointed to by lParam

	//// #ifdef UNICODE
	//RB_INSERTBAND = RB_INSERTBANDW
	//RB_SETBANDINFO = RB_SETBANDINFOW
	//// #else
	//RB_INSERTBAND = RB_INSERTBANDA
	//RB_SETBANDINFO = RB_SETBANDINFOA
	//// #endif

	// // for manual drag control
	// lparam == cursor pos
	//         // -1 means do it yourself.
	//         // -2 means use what you had saved before
	RB_BEGINDRAG    = (WM_USER + 24)
	RB_ENDDRAG      = (WM_USER + 25)
	RB_DRAGMOVE     = (WM_USER + 26)
	RB_GETBARHEIGHT = (WM_USER + 27)
	RB_GETBANDINFOW = (WM_USER + 28)
	RB_GETBANDINFOA = (WM_USER + 29)

	//// #ifdef UNICODE
	//RB_GETBANDINFO = RB_GETBANDINFOW
	//// #else
	//RB_GETBANDINFO = RB_GETBANDINFOA
	//// #endif

	RB_MINIMIZEBAND = (WM_USER + 30)
	RB_MAXIMIZEBAND = (WM_USER + 31)

	RB_GETDROPTARGET = (CCM_GETDROPTARGET)

	RB_GETBANDBORDERS = (WM_USER + 34) // returns in lparam = lprc the amount of edges added to band wparam

	RB_SHOWBAND   = (WM_USER + 35) // show/hide band
	RB_SETPALETTE = (WM_USER + 37)
	RB_GETPALETTE = (WM_USER + 38)
	RB_MOVEBAND   = (WM_USER + 39)

	RB_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	RB_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	RB_GETBANDMARGINS = (WM_USER + 40)
	RB_SETWINDOWTHEME = CCM_SETWINDOWTHEME
	// #endif

	// #if (_WIN32_IE >= 0x0600)
	RB_SETEXTENDEDSTYLE = (WM_USER + 41)
	RB_GETEXTENDEDSTYLE = (WM_USER + 42)
	// #endif      // _WIN32_IE >= 0x0600

	RB_PUSHCHEVRON = (WM_USER + 43)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	RB_SETBANDWIDTH = (WM_USER + 44) // set width for docked band
	// #endif      // (NTDDI_VERSION >= NTDDI_VISTA)

	RBN_HEIGHTCHANGE = (RBN_FIRST - 0)

	RBN_GETOBJECT     = (RBN_FIRST - 1)
	RBN_LAYOUTCHANGED = (RBN_FIRST - 2)
	RBN_AUTOSIZE      = (RBN_FIRST - 3)
	RBN_BEGINDRAG     = (RBN_FIRST - 4)
	RBN_ENDDRAG       = (RBN_FIRST - 5)
	RBN_DELETINGBAND  = (RBN_FIRST - 6) // Uses NMREBAR
	RBN_DELETEDBAND   = (RBN_FIRST - 7) // Uses NMREBAR
	RBN_CHILDSIZE     = (RBN_FIRST - 8)

	RBN_CHEVRONPUSHED = (RBN_FIRST - 10)

	// #if (_WIN32_IE >= 0x0600)
	RBN_SPLITTERDRAG = (RBN_FIRST - 11)
	// #endif      // _WIN32_IE >= 0x0600

	RBN_MINMAX = (RBN_FIRST - 21)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	RBN_AUTOBREAK = (RBN_FIRST - 22)
	// #endif

	// typedef struct tagNMREBARCHILDSIZE
	// {
	//     NMHDR hdr;
	//     UINT uBand;
	//     UINT wID;
	//     RECT rcChild;
	//     RECT rcBand;
	// } NMREBARCHILDSIZE, *LPNMREBARCHILDSIZE;

	// typedef struct tagNMREBAR
	// {
	//     NMHDR   hdr;
	//     DWORD   dwMask;           // RBNM_*
	//     UINT    uBand;
	//     UINT    fStyle;
	//     UINT    wID;
	//     LPARAM  lParam;
	// } NMREBAR, *LPNMREBAR;

	// // Mask flags for NMREBAR
	RBNM_ID     = 0
	RBNM_STYLE  = 0
	RBNM_LPARAM = 0

	// typedef struct tagNMRBAUTOSIZE
	// {
	//     NMHDR hdr;
	//     BOOL fChanged;
	//     RECT rcTarget;
	//     RECT rcActual;
	// } NMRBAUTOSIZE, *LPNMRBAUTOSIZE;

	// typedef struct tagNMREBARCHEVRON
	// {
	//     NMHDR hdr;
	//     UINT uBand;
	//     UINT wID;
	//     LPARAM lParam;
	//     RECT rc;
	//     LPARAM lParamNM;
	// } NMREBARCHEVRON, *LPNMREBARCHEVRON;

	// #if (_WIN32_IE >= 0x0600)
	// typedef struct tagNMREBARSPLITTER
	// {
	//     NMHDR hdr;
	//     RECT  rcSizing;
	// } NMREBARSPLITTER, *LPNMREBARSPLITTER;
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	RBAB_AUTOSIZE = 0 // These are not flags and are all mutually exclusive
	RBAB_ADDBAND  = 0

	// typedef struct tagNMREBARAUTOBREAK
	// {
	//     NMHDR hdr;
	//     UINT uBand;
	//     UINT wID;
	//     LPARAM lParam;
	//     UINT uMsg;
	//     UINT fStyleCurrent;
	//     BOOL fAutoBreak;
	// } NMREBARAUTOBREAK, *LPNMREBARAUTOBREAK;
	// #endif

	RBHT_NOWHERE = 0
	RBHT_CAPTION = 0
	RBHT_CLIENT  = 0
	RBHT_GRABBER = 0
	RBHT_CHEVRON = 0
	// #if (_WIN32_IE >= 0x0600)
	RBHT_SPLITTER = 0
	// #endif

	// typedef struct _RB_HITTESTINFO
	// {
	//     POINT pt;
	//     UINT flags;
	//     int iBand;
	// } RBHITTESTINFO, *LPRBHITTESTINFO;

	// #endif      // NOREBAR

	//====== TOOLTIPS CONTROL =====================================================

	// #ifndef NOTOOLTIPS

	// #ifdef _WIN32

	TOOLTIPS_CLASSW = "tooltips_class32"
	// #define TOOLTIPS_CLASSA         "tooltips_class32"

	// #ifdef UNICODE
	TOOLTIPS_CLASS = TOOLTIPS_CLASSW
	// #else
	//TOOLTIPS_CLASS = TOOLTIPS_CLASSA
	// #endif

	// #else
	// #define TOOLTIPS_CLASS          "tooltips_class"
	// #endif
	//
	//LPTOOLINFOA = LPTTTOOLINFOA
	//LPTOOLINFOW = LPTTTOOLINFOW
	//TOOLINFOA = TTTOOLINFOA
	//TOOLINFOW = TTTOOLINFOW

	//LPTOOLINFO = LPTTTOOLINFO
	//TOOLINFO = TTTOOLINFO
	//
	//TTTOOLINFOA_V1_SIZE = CCSIZEOF_STRUCT(TTTOOLINFOA, lpszText)
	//TTTOOLINFOW_V1_SIZE = CCSIZEOF_STRUCT(TTTOOLINFOW, lpszText)
	//TTTOOLINFOA_V2_SIZE = CCSIZEOF_STRUCT(TTTOOLINFOA, lParam)
	//TTTOOLINFOW_V2_SIZE = CCSIZEOF_STRUCT(TTTOOLINFOW, lParam)
	//TTTOOLINFOA_V3_SIZE = CCSIZEOF_STRUCT(TTTOOLINFOA, lpReserved)
	//TTTOOLINFOW_V3_SIZE = CCSIZEOF_STRUCT(TTTOOLINFOW, lpReserved)

	// typedef struct tagTOOLINFOA {
	//     UINT cbSize;
	//     UINT uFlags;
	//     HWND hwnd;
	//     UINT_PTR uId;
	//     RECT rect;
	//     HINSTANCE hinst;
	//     LPSTR lpszText;
	//     LPARAM lParam;
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	//     void *lpReserved;
	// #endif
	// } TTTOOLINFOA, NEAR *PTOOLINFOA, *LPTTTOOLINFOA;

	// typedef struct tagTOOLINFOW {
	//     UINT cbSize;
	//     UINT uFlags;
	//     HWND hwnd;
	//     UINT_PTR uId;
	//     RECT rect;
	//     HINSTANCE hinst;
	//     LPWSTR lpszText;
	//     LPARAM lParam;
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	//     void *lpReserved;
	// #endif
	// } TTTOOLINFOW, NEAR *PTOOLINFOW, *LPTTTOOLINFOW;
	//
	//// #ifdef UNICODE
	//TTTOOLINFO = TTTOOLINFOW
	//PTOOLINFO = PTOOLINFOW
	//LPTTTOOLINFO = LPTTTOOLINFOW
	//TTTOOLINFO_V1_SIZE = TTTOOLINFOW_V1_SIZE
	//// #else
	//PTOOLINFO = PTOOLINFOA
	//TTTOOLINFO = TTTOOLINFOA
	//LPTTTOOLINFO = LPTTTOOLINFOA
	//TTTOOLINFO_V1_SIZE = TTTOOLINFOA_V1_SIZE
	//// #endif

	// // begin_r_commctrl

	TTS_ALWAYSTIP = 0
	TTS_NOPREFIX  = 0
	TTS_NOANIMATE = 0
	TTS_NOFADE    = 0
	TTS_BALLOON   = 0
	TTS_CLOSE     = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TTS_USEVISUALSTYLE = 0 // Use themed hyperlinks

	// #endif

	// // end_r_commctrl

	TTF_IDISHWND = 0

	// // Use this to center around trackpoint in trackmode
	// // -OR- to center around tool in normal mode.
	// // Use TTF_ABSOLUTE to place the tip exactly at the track coords when
	// // in tracking mode.  TTF_ABSOLUTE can be used in conjunction with TTF_CENTERTIP
	// // to center the tip absolutely about the track point.

	TTF_CENTERTIP   = 0
	TTF_RTLREADING  = 0
	TTF_SUBCLASS    = 0
	TTF_TRACK       = 0
	TTF_ABSOLUTE    = 0
	TTF_TRANSPARENT = 0
	TTF_PARSELINKS  = 0
	TTF_DI_SETITEM  = 0 // valid only on the TTN_NEEDTEXT callback

	TTDT_AUTOMATIC = 0
	TTDT_RESHOW    = 1
	TTDT_AUTOPOP   = 2
	TTDT_INITIAL   = 3

	// // ToolTip Icons (Set with TTM_SETTITLE)
	TTI_NONE    = 0
	TTI_INFO    = 1
	TTI_WARNING = 2
	TTI_ERROR   = 3
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TTI_INFO_LARGE    = 4
	TTI_WARNING_LARGE = 5
	TTI_ERROR_LARGE   = 6
	// #endif  // (NTDDI_VERSION >= NTDDI_VISTA)

	// // Tool Tip Messages
	TTM_ACTIVATE     = (WM_USER + 1)
	TTM_SETDELAYTIME = (WM_USER + 3)
	TTM_ADDTOOLA     = (WM_USER + 4)
	TTM_ADDTOOLW     = (WM_USER + 50)
	TTM_DELTOOLA     = (WM_USER + 5)
	TTM_DELTOOLW     = (WM_USER + 51)
	TTM_NEWTOOLRECTA = (WM_USER + 6)
	TTM_NEWTOOLRECTW = (WM_USER + 52)
	TTM_RELAYEVENT   = (WM_USER + 7) // Win7: wParam = GetMessageExtraInfo() when relaying WM_MOUSEMOVE

	TTM_GETTOOLINFOA = (WM_USER + 8)
	TTM_GETTOOLINFOW = (WM_USER + 53)

	TTM_SETTOOLINFOA = (WM_USER + 9)
	TTM_SETTOOLINFOW = (WM_USER + 54)

	TTM_HITTESTA        = (WM_USER + 10)
	TTM_HITTESTW        = (WM_USER + 55)
	TTM_GETTEXTA        = (WM_USER + 11)
	TTM_GETTEXTW        = (WM_USER + 56)
	TTM_UPDATETIPTEXTA  = (WM_USER + 12)
	TTM_UPDATETIPTEXTW  = (WM_USER + 57)
	TTM_GETTOOLCOUNT    = (WM_USER + 13)
	TTM_ENUMTOOLSA      = (WM_USER + 14)
	TTM_ENUMTOOLSW      = (WM_USER + 58)
	TTM_GETCURRENTTOOLA = (WM_USER + 15)
	TTM_GETCURRENTTOOLW = (WM_USER + 59)
	TTM_WINDOWFROMPOINT = (WM_USER + 16)
	TTM_TRACKACTIVATE   = (WM_USER + 17) // wParam = TRUE/FALSE start end  lparam = LPTOOLINFO
	TTM_TRACKPOSITION   = (WM_USER + 18) // lParam = dwPos
	TTM_SETTIPBKCOLOR   = (WM_USER + 19)
	TTM_SETTIPTEXTCOLOR = (WM_USER + 20)
	TTM_GETDELAYTIME    = (WM_USER + 21)
	TTM_GETTIPBKCOLOR   = (WM_USER + 22)
	TTM_GETTIPTEXTCOLOR = (WM_USER + 23)
	TTM_SETMAXTIPWIDTH  = (WM_USER + 24)
	TTM_GETMAXTIPWIDTH  = (WM_USER + 25)
	TTM_SETMARGIN       = (WM_USER + 26) // lParam = lprc
	TTM_GETMARGIN       = (WM_USER + 27) // lParam = lprc
	TTM_POP             = (WM_USER + 28)
	TTM_UPDATE          = (WM_USER + 29)
	TTM_GETBUBBLESIZE   = (WM_USER + 30)
	TTM_ADJUSTRECT      = (WM_USER + 31)
	TTM_SETTITLEA       = (WM_USER + 32) // wParam = TTI_*, lParam = char* szTitle
	TTM_SETTITLEW       = (WM_USER + 33) // wParam = TTI_*, lParam = wchar* szTitle

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TTM_POPUP    = (WM_USER + 34)
	TTM_GETTITLE = (WM_USER + 35) // wParam = 0, lParam = TTGETTITLE*

	// typedef struct _TTGETTITLE
	// {
	//     DWORD dwSize;
	//     UINT uTitleBitmap;
	//     UINT cch;
	//     WCHAR* pszTitle;
	// } TTGETTITLE, *PTTGETTITLE;
	// #endif
	//
	//// #ifdef UNICODE
	//TTM_ADDTOOL = TTM_ADDTOOLW
	//TTM_DELTOOL = TTM_DELTOOLW
	//TTM_NEWTOOLRECT = TTM_NEWTOOLRECTW
	//TTM_GETTOOLINFO = TTM_GETTOOLINFOW
	//TTM_SETTOOLINFO = TTM_SETTOOLINFOW
	//TTM_HITTEST = TTM_HITTESTW
	//TTM_GETTEXT = TTM_GETTEXTW
	//TTM_UPDATETIPTEXT = TTM_UPDATETIPTEXTW
	//TTM_ENUMTOOLS = TTM_ENUMTOOLSW
	//TTM_GETCURRENTTOOL = TTM_GETCURRENTTOOLW
	//TTM_SETTITLE = TTM_SETTITLEW
	//// #else
	//TTM_ADDTOOL = TTM_ADDTOOLA
	//TTM_DELTOOL = TTM_DELTOOLA
	//TTM_NEWTOOLRECT = TTM_NEWTOOLRECTA
	//TTM_GETTOOLINFO = TTM_GETTOOLINFOA
	//TTM_SETTOOLINFO = TTM_SETTOOLINFOA
	//TTM_HITTEST = TTM_HITTESTA
	//TTM_GETTEXT = TTM_GETTEXTA
	//TTM_UPDATETIPTEXT = TTM_UPDATETIPTEXTA
	//TTM_ENUMTOOLS = TTM_ENUMTOOLSA
	//TTM_GETCURRENTTOOL = TTM_GETCURRENTTOOLA
	//TTM_SETTITLE = TTM_SETTITLEA
	//// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TTM_SETWINDOWTHEME = CCM_SETWINDOWTHEME
	// #endif

	//
	//LPHITTESTINFOW = LPTTHITTESTINFOW
	//LPHITTESTINFOA = LPTTHITTESTINFOA
	//LPHITTESTINFO = LPTTHITTESTINFO

	// typedef struct _TT_HITTESTINFOA {
	//     HWND hwnd;
	//     POINT pt;
	//     TTTOOLINFOA ti;
	// } TTHITTESTINFOA, *LPTTHITTESTINFOA;

	// typedef struct _TT_HITTESTINFOW {
	//     HWND hwnd;
	//     POINT pt;
	//     TTTOOLINFOW ti;
	// } TTHITTESTINFOW, *LPTTHITTESTINFOW;

	//// #ifdef UNICODE
	//TTHITTESTINFO = TTHITTESTINFOW
	//LPTTHITTESTINFO = LPTTHITTESTINFOW
	//// #else
	//TTHITTESTINFO = TTHITTESTINFOA
	//LPTTHITTESTINFO = LPTTHITTESTINFOA
	//// #endif

	TTN_GETDISPINFOA = (TTN_FIRST - 0)
	TTN_GETDISPINFOW = (TTN_FIRST - 10)
	TTN_SHOW         = (TTN_FIRST - 1)
	TTN_POP          = (TTN_FIRST - 2)
	TTN_LINKCLICK    = (TTN_FIRST - 3)
	//
	//// #ifdef UNICODE
	//TTN_GETDISPINFO = TTN_GETDISPINFOW
	//// #else
	//TTN_GETDISPINFO = TTN_GETDISPINFOA
	//// #endif
	//
	//TTN_NEEDTEXT = TTN_GETDISPINFO
	//TTN_NEEDTEXTA = TTN_GETDISPINFOA
	//TTN_NEEDTEXTW = TTN_GETDISPINFOW
	//
	//
	//TOOLTIPTEXTW = NMTTDISPINFOW
	//TOOLTIPTEXTA = NMTTDISPINFOA
	//LPTOOLTIPTEXTA = LPNMTTDISPINFOA
	//LPTOOLTIPTEXTW = LPNMTTDISPINFOW
	//
	//TOOLTIPTEXT = NMTTDISPINFO
	//LPTOOLTIPTEXT = LPNMTTDISPINFO
	//
	//NMTTDISPINFOA_V1_SIZE = CCSIZEOF_STRUCT(NMTTDISPINFOA, uFlags)
	//NMTTDISPINFOW_V1_SIZE = CCSIZEOF_STRUCT(NMTTDISPINFOW, uFlags)

	// typedef struct tagNMTTDISPINFOA {
	//     NMHDR hdr;
	//     LPSTR lpszText;
	//     char szText[80];
	//     HINSTANCE hinst;
	//     UINT uFlags;
	//     LPARAM lParam;
	// } NMTTDISPINFOA, *LPNMTTDISPINFOA;

	// typedef struct tagNMTTDISPINFOW {
	//     NMHDR hdr;
	//     LPWSTR lpszText;
	//     WCHAR szText[80];
	//     HINSTANCE hinst;
	//     UINT uFlags;
	//     LPARAM lParam;
	// } NMTTDISPINFOW, *LPNMTTDISPINFOW;
	//
	//// #ifdef UNICODE
	//NMTTDISPINFO = NMTTDISPINFOW
	//LPNMTTDISPINFO = LPNMTTDISPINFOW
	//NMTTDISPINFO_V1_SIZE = NMTTDISPINFOW_V1_SIZE
	//// #else
	//NMTTDISPINFO = NMTTDISPINFOA
	//LPNMTTDISPINFO = LPNMTTDISPINFOA
	//NMTTDISPINFO_V1_SIZE = NMTTDISPINFOA_V1_SIZE
	//// #endif

	// #endif      // NOTOOLTIPS

	//====== STATUS BAR CONTROL ===================================================

	// #ifndef NOSTATUSBAR

	// // begin_r_commctrl

	SBARS_SIZEGRIP = 0
	SBARS_TOOLTIPS = 0

	// // this is a status bar flag, preference to SBARS_TOOLTIPS
	SBT_TOOLTIPS = 0

	// // end_r_commctrl

	// WINCOMMCTRLAPI void WINAPI DrawStatusTextA(HDC hDC, LPCRECT lprc, LPCSTR pszText, UINT uFlags);
	// WINCOMMCTRLAPI void WINAPI DrawStatusTextW(HDC hDC, LPCRECT lprc, LPCWSTR pszText, UINT uFlags);

	// WINCOMMCTRLAPI HWND WINAPI CreateStatusWindowA(LONG style, LPCSTR lpszText, HWND hwndParent, UINT wID);
	// WINCOMMCTRLAPI HWND WINAPI CreateStatusWindowW(LONG style, LPCWSTR lpszText, HWND hwndParent, UINT wID);
	//
	//// #ifdef UNICODE
	//CreateStatusWindow = CreateStatusWindowW
	//DrawStatusText = DrawStatusTextW
	//// #else
	//CreateStatusWindow = CreateStatusWindowA
	//DrawStatusText = DrawStatusTextA
	//// #endif

	// #ifdef _WIN32
	STATUSCLASSNAMEW = "msctls_statusbar32"
	// #define STATUSCLASSNAMEA        "msctls_statusbar32"
	//
	//// #ifdef UNICODE
	//STATUSCLASSNAME = STATUSCLASSNAMEW
	//// #else
	//STATUSCLASSNAME = STATUSCLASSNAMEA
	//// #endif

	// #else
	// #define STATUSCLASSNAME         "msctls_statusbar"
	// #endif

	SB_SETTEXTA       = (WM_USER + 1)
	SB_SETTEXTW       = (WM_USER + 11)
	SB_GETTEXTA       = (WM_USER + 2)
	SB_GETTEXTW       = (WM_USER + 13)
	SB_GETTEXTLENGTHA = (WM_USER + 3)
	SB_GETTEXTLENGTHW = (WM_USER + 12)
	//
	//// #ifdef UNICODE
	//SB_GETTEXT = SB_GETTEXTW
	//SB_SETTEXT = SB_SETTEXTW
	//SB_GETTEXTLENGTH = SB_GETTEXTLENGTHW
	//SB_SETTIPTEXT = SB_SETTIPTEXTW
	//SB_GETTIPTEXT = SB_GETTIPTEXTW
	//// #else
	//SB_GETTEXT = SB_GETTEXTA
	//SB_SETTEXT = SB_SETTEXTA
	//SB_GETTEXTLENGTH = SB_GETTEXTLENGTHA
	//SB_SETTIPTEXT = SB_SETTIPTEXTA
	//SB_GETTIPTEXT = SB_GETTIPTEXTA
	//// #endif

	SB_SETPARTS         = (WM_USER + 4)
	SB_GETPARTS         = (WM_USER + 6)
	SB_GETBORDERS       = (WM_USER + 7)
	SB_SETMINHEIGHT     = (WM_USER + 8)
	SB_SIMPLE           = (WM_USER + 9)
	SB_GETRECT          = (WM_USER + 10)
	SB_ISSIMPLE         = (WM_USER + 14)
	SB_SETICON          = (WM_USER + 15)
	SB_SETTIPTEXTA      = (WM_USER + 16)
	SB_SETTIPTEXTW      = (WM_USER + 17)
	SB_GETTIPTEXTA      = (WM_USER + 18)
	SB_GETTIPTEXTW      = (WM_USER + 19)
	SB_GETICON          = (WM_USER + 20)
	SB_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	SB_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT

	SBT_OWNERDRAW    = 0
	SBT_NOBORDERS    = 0
	SBT_POPOUT       = 0
	SBT_RTLREADING   = 0
	SBT_NOTABPARSING = 0

	SB_SETBKCOLOR = CCM_SETBKCOLOR // lParam = bkColor

	// // status bar notifications
	SBN_SIMPLEMODECHANGE = (SBN_FIRST - 0)

	// // refers to the data saved for simple mode
	SB_SIMPLEID = 0

	// #endif      // NOSTATUSBAR

	//====== MENU HELP ============================================================

	// #ifndef NOMENUHELP

	//WINCOMMCTRLAPI void WINAPI MenuHelp(UINT uMsg, WPARAM wParam, LPARAM lParam, HMENU hMainMenu, HINSTANCE hInst, HWND hwndStatus, _In_reads_(_Inexpressible_(2 + 2n && n >= 1)) UINT *lpwIDs);
	// WINCOMMCTRLAPI BOOL WINAPI ShowHideMenuCtl(_In_ HWND hWnd, _In_ UINT_PTR uFlags, _In_z_ LPINT lpInfo);
	// WINCOMMCTRLAPI void WINAPI GetEffectiveClientRect(_In_ HWND hWnd, _Out_ LPRECT lprc,  _In_z_ const INT *lpInfo);

	//MINSYSCOMMAND = SC_SIZE

	// #endif

	//====== TRACKBAR CONTROL =====================================================

	// #ifndef NOTRACKBAR

	// #ifdef _WIN32

	// #define TRACKBAR_CLASSA         "msctls_trackbar32"
	TRACKBAR_CLASSW = "msctls_trackbar32"

	// #ifdef UNICODE
	TRACKBAR_CLASS = TRACKBAR_CLASSW
	// #else
	//TRACKBAR_CLASS = TRACKBAR_CLASSA
	// #endif

	// #else
	// #define TRACKBAR_CLASS          "msctls_trackbar"
	// #endif

	// // begin_r_commctrl

	TBS_AUTOTICKS      = 0
	TBS_VERT           = 0
	TBS_HORZ           = 0
	TBS_TOP            = 0
	TBS_BOTTOM         = 0
	TBS_LEFT           = 0
	TBS_RIGHT          = 0
	TBS_BOTH           = 0
	TBS_NOTICKS        = 0
	TBS_ENABLESELRANGE = 0
	TBS_FIXEDLENGTH    = 0
	TBS_NOTHUMB        = 0
	TBS_TOOLTIPS       = 0
	TBS_REVERSED       = 0 // Accessibility hint: the smaller number (usually the min value) means "high" and the larger number (usually the max value) means "low"

	TBS_DOWNISLEFT = 0 // Down=Left and Up=Right (default is Down=Right and Up=Left)

	// #if (_WIN32_IE >= 0x0600)
	TBS_NOTIFYBEFOREMOVE = 0 // Trackbar should notify parent before repositioning the slider due to user action (enables snapping)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TBS_TRANSPARENTBKGND = 0 // Background is painted by the parent via WM_PRINTCLIENT
	// #endif

	// // end_r_commctrl

	TBM_GETPOS         = (WM_USER)
	TBM_GETRANGEMIN    = (WM_USER + 1)
	TBM_GETRANGEMAX    = (WM_USER + 2)
	TBM_GETTIC         = (WM_USER + 3)
	TBM_SETTIC         = (WM_USER + 4)
	TBM_SETPOS         = (WM_USER + 5)
	TBM_SETRANGE       = (WM_USER + 6)
	TBM_SETRANGEMIN    = (WM_USER + 7)
	TBM_SETRANGEMAX    = (WM_USER + 8)
	TBM_CLEARTICS      = (WM_USER + 9)
	TBM_SETSEL         = (WM_USER + 10)
	TBM_SETSELSTART    = (WM_USER + 11)
	TBM_SETSELEND      = (WM_USER + 12)
	TBM_GETPTICS       = (WM_USER + 14)
	TBM_GETTICPOS      = (WM_USER + 15)
	TBM_GETNUMTICS     = (WM_USER + 16)
	TBM_GETSELSTART    = (WM_USER + 17)
	TBM_GETSELEND      = (WM_USER + 18)
	TBM_CLEARSEL       = (WM_USER + 19)
	TBM_SETTICFREQ     = (WM_USER + 20)
	TBM_SETPAGESIZE    = (WM_USER + 21)
	TBM_GETPAGESIZE    = (WM_USER + 22)
	TBM_SETLINESIZE    = (WM_USER + 23)
	TBM_GETLINESIZE    = (WM_USER + 24)
	TBM_GETTHUMBRECT   = (WM_USER + 25)
	TBM_GETCHANNELRECT = (WM_USER + 26)
	TBM_SETTHUMBLENGTH = (WM_USER + 27)
	TBM_GETTHUMBLENGTH = (WM_USER + 28)
	TBM_SETTOOLTIPS    = (WM_USER + 29)
	TBM_GETTOOLTIPS    = (WM_USER + 30)
	TBM_SETTIPSIDE     = (WM_USER + 31)
	// // TrackBar Tip Side flags
	TBTS_TOP    = 0
	TBTS_LEFT   = 1
	TBTS_BOTTOM = 2
	TBTS_RIGHT  = 3

	TBM_SETBUDDY         = (WM_USER + 32) // wparam = BOOL fLeft; (or right)
	TBM_GETBUDDY         = (WM_USER + 33) // wparam = BOOL fLeft; (or right)
	TBM_SETPOSNOTIFY     = (WM_USER + 34)
	TBM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	TBM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT

	TB_LINEUP        = 0
	TB_LINEDOWN      = 1
	TB_PAGEUP        = 2
	TB_PAGEDOWN      = 3
	TB_THUMBPOSITION = 4
	TB_THUMBTRACK    = 5
	TB_TOP           = 6
	TB_BOTTOM        = 7
	TB_ENDTRACK      = 8

	// // custom draw item specs
	TBCD_TICS    = 0
	TBCD_THUMB   = 0
	TBCD_CHANNEL = 0

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TRBN_THUMBPOSCHANGING = (TRBN_FIRST - 1)

	// // Structure for Trackbar's TRBN_THUMBPOSCHANGING notification
	// typedef struct tagTRBTHUMBPOSCHANGING
	// {
	//     NMHDR hdr;
	//     DWORD dwPos;
	//     int nReason;
	// } NMTRBTHUMBPOSCHANGING;
	// #endif

	// #endif // trackbar

	//====== DRAG LIST CONTROL ====================================================

	// #ifndef NODRAGLIST

	// typedef struct tagDRAGLISTINFO {
	//     UINT uNotification;
	//     HWND hWnd;
	//     POINT ptCursor;
	// } DRAGLISTINFO, *LPDRAGLISTINFO;

	DL_BEGINDRAG  = (WM_USER + 133)
	DL_DRAGGING   = (WM_USER + 134)
	DL_DROPPED    = (WM_USER + 135)
	DL_CANCELDRAG = (WM_USER + 136)

	DL_CURSORSET  = 0
	DL_STOPCURSOR = 1
	DL_COPYCURSOR = 2
	DL_MOVECURSOR = 3

	DRAGLISTMSGSTRING = "commctrl_DragListMsg"

	// WINCOMMCTRLAPI BOOL WINAPI MakeDragList(HWND hLB);
	// WINCOMMCTRLAPI void WINAPI DrawInsert(HWND handParent, HWND hLB, int nItem);

	// WINCOMMCTRLAPI int WINAPI LBItemFromPt(HWND hLB, POINT pt, BOOL bAutoScroll);

	// #endif

	//====== UPDOWN CONTROL =======================================================

	// #ifndef NOUPDOWN

	// #ifdef _WIN32

	// #define UPDOWN_CLASSA           "msctls_updown32"
	UPDOWN_CLASSW = "msctls_updown32"

	// #ifdef UNICODE
	UPDOWN_CLASS = UPDOWN_CLASSW
	// #else
	//UPDOWN_CLASS = UPDOWN_CLASSA
	// #endif

	// #else
	// #define UPDOWN_CLASS            "msctls_updown"
	// #endif

	// typedef struct _UDACCEL {
	//     UINT nSec;
	//     UINT nInc;
	// } UDACCEL, *LPUDACCEL;

	UD_MAXVAL = 0
	UD_MINVAL = (-UD_MAXVAL)

	// // begin_r_commctrl

	UDS_WRAP        = 0
	UDS_SETBUDDYINT = 0
	UDS_ALIGNRIGHT  = 0
	UDS_ALIGNLEFT   = 0
	UDS_AUTOBUDDY   = 0
	UDS_ARROWKEYS   = 0
	UDS_HORZ        = 0
	UDS_NOTHOUSANDS = 0
	UDS_HOTTRACK    = 0

	// // end_r_commctrl

	UDM_SETRANGE         = (WM_USER + 101)
	UDM_GETRANGE         = (WM_USER + 102)
	UDM_SETPOS           = (WM_USER + 103)
	UDM_GETPOS           = (WM_USER + 104)
	UDM_SETBUDDY         = (WM_USER + 105)
	UDM_GETBUDDY         = (WM_USER + 106)
	UDM_SETACCEL         = (WM_USER + 107)
	UDM_GETACCEL         = (WM_USER + 108)
	UDM_SETBASE          = (WM_USER + 109)
	UDM_GETBASE          = (WM_USER + 110)
	UDM_SETRANGE32       = (WM_USER + 111)
	UDM_GETRANGE32       = (WM_USER + 112) // wParam & lParam are LPINT
	UDM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	UDM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	UDM_SETPOS32         = (WM_USER + 113)
	UDM_GETPOS32         = (WM_USER + 114)

	// WINCOMMCTRLAPI HWND WINAPI CreateUpDownControl(DWORD dwStyle, int x, int y, int cx, int cy,
	//                                 HWND hParent, int nID, HINSTANCE hInst,
	//                                 HWND hBuddy,
	//                                 int nUpper, int nLower, int nPos);
	//
	//NM_UPDOWN = NMUPDOWN
	//LPNM_UPDOWN = LPNMUPDOWN

	// typedef struct _NM_UPDOWN
	// {
	//     NMHDR hdr;
	//     int iPos;
	//     int iDelta;
	// } NMUPDOWN, *LPNMUPDOWN;

	UDN_DELTAPOS = (UDN_FIRST - 1)

	// #endif  // NOUPDOWN

	//====== PROGRESS CONTROL =====================================================

	// #ifndef NOPROGRESS

	// #ifdef _WIN32

	// #define PROGRESS_CLASSA         "msctls_progress32"
	PROGRESS_CLASSW = "msctls_progress32"

	// #ifdef UNICODE
	PROGRESS_CLASS = PROGRESS_CLASSW
	// #else
	//PROGRESS_CLASS = PROGRESS_CLASSA
	// #endif

	// #else
	// #define PROGRESS_CLASS          "msctls_progress"
	// #endif

	// // begin_r_commctrl

	PBS_SMOOTH   = 0
	PBS_VERTICAL = 0

	// // end_r_commctrl

	PBM_SETRANGE   = (WM_USER + 1)
	PBM_SETPOS     = (WM_USER + 2)
	PBM_DELTAPOS   = (WM_USER + 3)
	PBM_SETSTEP    = (WM_USER + 4)
	PBM_STEPIT     = (WM_USER + 5)
	PBM_SETRANGE32 = (WM_USER + 6) // lParam = high, wParam = low
	// typedef struct
	// {
	//    int iLow;
	//    int iHigh;
	// } PBRANGE, *PPBRANGE;
	PBM_GETRANGE    = (WM_USER + 7) // wParam = return (TRUE ? low : high). lParam = PPBRANGE or NULL
	PBM_GETPOS      = (WM_USER + 8)
	PBM_SETBARCOLOR = (WM_USER + 9)  // lParam = bar color
	PBM_SETBKCOLOR  = CCM_SETBKCOLOR // lParam = bkColor

	// // begin_r_commctrl

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	PBS_MARQUEE = 0
	// #endif       // (NTDDI_VERSION >= NTDDI_WINXP)

	// // end_r_commctrl

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	PBM_SETMARQUEE = (WM_USER + 10)
	// #endif      // (NTDDI_VERSION >= NTDDI_WINXP)

	// // begin_r_commctrl
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	PBS_SMOOTHREVERSE = 0
	// #endif       // (NTDDI_VERSION >= NTDDI_VISTA)

	// // end_r_commctrl

	// #if (NTDDI_VERSION >= NTDDI_VISTA)

	PBM_GETSTEP     = (WM_USER + 13)
	PBM_GETBKCOLOR  = (WM_USER + 14)
	PBM_GETBARCOLOR = (WM_USER + 15)
	PBM_SETSTATE    = (WM_USER + 16) // wParam = PBST_[State] (NORMAL, ERROR, PAUSED)
	PBM_GETSTATE    = (WM_USER + 17)

	PBST_NORMAL = 0
	PBST_ERROR  = 0
	PBST_PAUSED = 0
	// #endif      // (NTDDI_VERSION >= NTDDI_VISTA)

	// #endif  // NOPROGRESS

	//====== HOTKEY CONTROL =======================================================

	// #ifndef NOHOTKEY

	HOTKEYF_SHIFT   = 0
	HOTKEYF_CONTROL = 0
	HOTKEYF_ALT     = 0
	// #ifdef _MAC
	//HOTKEYF_EXT = 0
	// #else
	HOTKEYF_EXT = 0
	// #endif

	HKCOMB_NONE = 0
	HKCOMB_S    = 0
	HKCOMB_C    = 0
	HKCOMB_A    = 0
	HKCOMB_SC   = 0
	HKCOMB_SA   = 0
	HKCOMB_CA   = 0
	HKCOMB_SCA  = 0

	HKM_SETHOTKEY = (WM_USER + 1)
	HKM_GETHOTKEY = (WM_USER + 2)
	HKM_SETRULES  = (WM_USER + 3)

	// #ifdef _WIN32

	// #define HOTKEY_CLASSA           "msctls_hotkey32"
	HOTKEY_CLASSW = "msctls_hotkey32"

	//// #ifdef UNICODE
	//HOTKEY_CLASS = HOTKEY_CLASSW
	//// #else
	//HOTKEY_CLASS = HOTKEY_CLASSA
	//// #endif

	// #else
	// #define HOTKEY_CLASS            "msctls_hotkey"
	// #endif

	// #endif  // NOHOTKEY

	// // begin_r_commctrl

	//====== COMMON CONTROL STYLES ================================================

	CCS_TOP           = 0
	CCS_NOMOVEY       = 0
	CCS_BOTTOM        = 0
	CCS_NORESIZE      = 0
	CCS_NOPARENTALIGN = 0
	CCS_ADJUSTABLE    = 0
	CCS_NODIVIDER     = 0
	CCS_VERT          = 0
	CCS_LEFT          = (CCS_VERT | CCS_TOP)
	CCS_RIGHT         = (CCS_VERT | CCS_BOTTOM)
	CCS_NOMOVEX       = (CCS_VERT | CCS_NOMOVEY)

	// // end_r_commctrl

	//====== SysLink control =========================================

	// #ifdef _WIN32
	// #if (NTDDI_VERSION >= NTDDI_WINXP)

	INVALID_LINK_INDEX = (-1)
	MAX_LINKID_TEXT    = 48
	L_MAX_URL_LENGTH   = (2048 + 32 + len("://")) // sizeof("://"))

	WC_LINK = "SysLink"

	// // begin_r_commctrl

	LWS_TRANSPARENT  = 0
	LWS_IGNORERETURN = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LWS_NOPREFIX       = 0
	LWS_USEVISUALSTYLE = 0
	LWS_USECUSTOMTEXT  = 0
	LWS_RIGHT          = 0
	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// // end_r_commctrl

	LIF_ITEMINDEX = 0
	LIF_STATE     = 0
	LIF_ITEMID    = 0
	LIF_URL       = 0

	LIS_FOCUSED = 0
	LIS_ENABLED = 0
	LIS_VISITED = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LIS_HOTTRACK      = 0
	LIS_DEFAULTCOLORS = 0 // Don't use any custom text colors
	// #endif

	// typedef struct tagLITEM
	// {
	//     UINT        mask ;
	//     int         iLink ;
	//     UINT        state ;
	//     UINT        stateMask ;
	//     WCHAR       szID[MAX_LINKID_TEXT] ;
	//     WCHAR       szUrl[L_MAX_URL_LENGTH] ;
	// } LITEM, * PLITEM ;

	// typedef struct tagLHITTESTINFO
	// {
	//     POINT       pt ;
	//     LITEM     item ;
	// } LHITTESTINFO, *PLHITTESTINFO ;

	// typedef struct tagNMLINK
	// {
	//     NMHDR       hdr;
	//     LITEM     item ;
	// } NMLINK,  *PNMLINK;

	// //  SysLink notifications
	// //  NM_CLICK   // wParam: control ID, lParam: PNMLINK, ret: ignored.

	// //  LinkWindow messages
	LM_HITTEST        = (WM_USER + 0x300)   // wParam: n/a, lparam: PLHITTESTINFO, ret: BOOL
	LM_GETIDEALHEIGHT = (WM_USER + 0x301)   // wParam: cxMaxWidth, lparam: n/a, ret: cy
	LM_SETITEM        = (WM_USER + 0x302)   // wParam: n/a, lparam: LITEM*, ret: BOOL
	LM_GETITEM        = (WM_USER + 0x303)   // wParam: n/a, lparam: LITEM*, ret: BOOL
	LM_GETIDEALSIZE   = (LM_GETIDEALHEIGHT) // wParam: cxMaxWidth, lparam: SIZE*, ret: cy

	// #endif

	// #endif // _WIN32
	//====== End SysLink control =========================================

	//====== LISTVIEW CONTROL =====================================================

	// #ifndef NOLISTVIEW

	// #ifdef _WIN32

	// #define WC_LISTVIEWA            "SysListView32"
	WC_LISTVIEWW = "SysListView32"

	// #ifdef UNICODE
	WC_LISTVIEW = WC_LISTVIEWW
	// #else
	//WC_LISTVIEW = WC_LISTVIEWA
	// #endif

	// #else
	// #define WC_LISTVIEW             "SysListView"
	// #endif

	// // begin_r_commctrl

	LVS_ICON            = 0
	LVS_REPORT          = 0
	LVS_SMALLICON       = 0
	LVS_LIST            = 0
	LVS_TYPEMASK        = 0
	LVS_SINGLESEL       = 0
	LVS_SHOWSELALWAYS   = 0
	LVS_SORTASCENDING   = 0
	LVS_SORTDESCENDING  = 0
	LVS_SHAREIMAGELISTS = 0
	LVS_NOLABELWRAP     = 0
	LVS_AUTOARRANGE     = 0
	LVS_EDITLABELS      = 0
	LVS_OWNERDATA       = 0
	LVS_NOSCROLL        = 0

	LVS_TYPESTYLEMASK = 0

	LVS_ALIGNTOP  = 0
	LVS_ALIGNLEFT = 0
	LVS_ALIGNMASK = 0

	LVS_OWNERDRAWFIXED = 0
	LVS_NOCOLUMNHEADER = 0
	LVS_NOSORTHEADER   = 0

	// // end_r_commctrl

	LVM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	// #define ListView_SetUnicodeFormat(hwnd, fUnicode)  \
	//     (BOOL)SNDMSG((hwnd), LVM_SETUNICODEFORMAT, (WPARAM)(fUnicode), 0)

	LVM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	// #define ListView_GetUnicodeFormat(hwnd)  \
	//     (BOOL)SNDMSG((hwnd), LVM_GETUNICODEFORMAT, 0, 0)

	LVM_GETBKCOLOR = (LVM_FIRST + 0)
	// #define ListView_GetBkColor(hwnd)  \
	//     (COLORREF)SNDMSG((hwnd), LVM_GETBKCOLOR, 0, 0L)

	LVM_SETBKCOLOR = (LVM_FIRST + 1)
	// #define ListView_SetBkColor(hwnd, clrBk) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETBKCOLOR, 0, (LPARAM)(COLORREF)(clrBk))

	LVM_GETIMAGELIST = (LVM_FIRST + 2)
	// #define ListView_GetImageList(hwnd, iImageList) \
	//     (HIMAGELIST)SNDMSG((hwnd), LVM_GETIMAGELIST, (WPARAM)(INT)(iImageList), 0L)

	LVSIL_NORMAL      = 0
	LVSIL_SMALL       = 1
	LVSIL_STATE       = 2
	LVSIL_GROUPHEADER = 3

	LVM_SETIMAGELIST = (LVM_FIRST + 3)
	// #define ListView_SetImageList(hwnd, himl, iImageList) \
	//     (HIMAGELIST)SNDMSG((hwnd), LVM_SETIMAGELIST, (WPARAM)(iImageList), (LPARAM)(HIMAGELIST)(himl))

	LVM_GETITEMCOUNT = (LVM_FIRST + 4)
	// #define ListView_GetItemCount(hwnd) \
	//     (int)SNDMSG((hwnd), LVM_GETITEMCOUNT, 0, 0L)

	LVIF_TEXT        = 0
	LVIF_IMAGE       = 0
	LVIF_PARAM       = 0
	LVIF_STATE       = 0
	LVIF_INDENT      = 0
	LVIF_NORECOMPUTE = 0
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	LVIF_GROUPID = 0
	LVIF_COLUMNS = 0
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVIF_COLFMT = 0 // The piColFmt member is valid in addition to puColumns
	// #endif

	LVIS_FOCUSED     = 0
	LVIS_SELECTED    = 0
	LVIS_CUT         = 0
	LVIS_DROPHILITED = 0
	LVIS_GLOW        = 0
	LVIS_ACTIVATING  = 0

	LVIS_OVERLAYMASK    = 0
	LVIS_STATEIMAGEMASK = 0

	// #define INDEXTOSTATEIMAGEMASK(i) ((i) << 12)

	I_INDENTCALLBACK = (-1)
	//LV_ITEMA = LVITEMA
	//LV_ITEMW = LVITEMW

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	I_GROUPIDCALLBACK = (-1)
	I_GROUPIDNONE     = (-2)
	// #endif
	//LV_ITEM = LVITEM
	//
	//LVITEMA_V1_SIZE = CCSIZEOF_STRUCT(LVITEMA, lParam)
	//LVITEMW_V1_SIZE = CCSIZEOF_STRUCT(LVITEMW, lParam)
	//
	//// #if (NTDDI_VERSION >= NTDDI_VISTA) // Will be unused downlevel, but sizeof(LVITEMA) must be equal to sizeof(LVITEMW)
	//LVITEMA_V5_SIZE = CCSIZEOF_STRUCT(LVITEMA, puColumns)
	//LVITEMW_V5_SIZE = CCSIZEOF_STRUCT(LVITEMW, puColumns)
	//
	//// #ifdef UNICODE
	//LVITEM_V5_SIZE = LVITEMW_V5_SIZE
	//// #else
	//LVITEM_V5_SIZE = LVITEMA_V5_SIZE
	//// #endif
	//// #endif

	// typedef struct tagLVITEMA
	// {
	//     UINT mask;
	//     int iItem;
	//     int iSubItem;
	//     UINT state;
	//     UINT stateMask;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iImage;
	//     LPARAM lParam;
	//     int iIndent;
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	//     int iGroupId;
	//     UINT cColumns; // tile view columns
	//     PUINT puColumns;
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_VISTA) // Will be unused downlevel, but sizeof(LVITEMA) must be equal to sizeof(LVITEMW)
	//     int* piColFmt;
	//     int iGroup; // readonly. only valid for owner data.
	// #endif
	// } LVITEMA, *LPLVITEMA;

	// typedef struct tagLVITEMW
	// {
	//     UINT mask;
	//     int iItem;
	//     int iSubItem;
	//     UINT state;
	//     UINT stateMask;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iImage;
	//     LPARAM lParam;
	//     int iIndent;
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	//     int iGroupId;
	//     UINT cColumns; // tile view columns
	//     PUINT puColumns;
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     int* piColFmt;
	//     int iGroup; // readonly. only valid for owner data.
	// #endif
	// } LVITEMW, *LPLVITEMW;

	//// #ifdef UNICODE
	//LVITEM = LVITEMW
	//LPLVITEM = LPLVITEMW
	//LVITEM_V1_SIZE = LVITEMW_V1_SIZE
	//// #else
	//LVITEM = LVITEMA
	//LPLVITEM = LPLVITEMA
	//LVITEM_V1_SIZE = LVITEMA_V1_SIZE
	//// #endif

	//LPSTR_TEXTCALLBACKW = ((LPWSTR)-1L)
	//LPSTR_TEXTCALLBACKA = ((LPSTR)-1L)
	//// #ifdef UNICODE
	//LPSTR_TEXTCALLBACK = LPSTR_TEXTCALLBACKW
	//// #else
	//LPSTR_TEXTCALLBACK = LPSTR_TEXTCALLBACKA
	//// #endif

	I_IMAGECALLBACK = (-1)
	I_IMAGENONE     = (-2)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	// // For tileview
	I_COLUMNSCALLBACK = -1
	// #endif

	//LVM_GETITEMA = (LVM_FIRST + 5)
	//LVM_GETITEMW = (LVM_FIRST + 75)
	//// #ifdef UNICODE
	//LVM_GETITEM = LVM_GETITEMW
	//// #else
	//LVM_GETITEM = LVM_GETITEMA
	//// #endif

	// #define ListView_GetItem(hwnd, pitem) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETITEM, 0, (LPARAM)(LV_ITEM *)(pitem))

	LVM_SETITEMA = (LVM_FIRST + 6)
	LVM_SETITEMW = (LVM_FIRST + 76)
	//// #ifdef UNICODE
	//LVM_SETITEM = LVM_SETITEMW
	//// #else
	//LVM_SETITEM = LVM_SETITEMA
	//// #endif

	// #define ListView_SetItem(hwnd, pitem) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETITEM, 0, (LPARAM)(const LV_ITEM *)(pitem))

	LVM_INSERTITEMA = (LVM_FIRST + 7)
	LVM_INSERTITEMW = (LVM_FIRST + 77)
	//// #ifdef UNICODE
	//LVM_INSERTITEM = LVM_INSERTITEMW
	//// #else
	//LVM_INSERTITEM = LVM_INSERTITEMA
	//// #endif
	// #define ListView_InsertItem(hwnd, pitem)   \
	//     (int)SNDMSG((hwnd), LVM_INSERTITEM, 0, (LPARAM)(const LV_ITEM *)(pitem))

	LVM_DELETEITEM = (LVM_FIRST + 8)
	// #define ListView_DeleteItem(hwnd, i) \
	//     (BOOL)SNDMSG((hwnd), LVM_DELETEITEM, (WPARAM)(int)(i), 0L)

	LVM_DELETEALLITEMS = (LVM_FIRST + 9)
	// #define ListView_DeleteAllItems(hwnd) \
	//     (BOOL)SNDMSG((hwnd), LVM_DELETEALLITEMS, 0, 0L)

	LVM_GETCALLBACKMASK = (LVM_FIRST + 10)
	// #define ListView_GetCallbackMask(hwnd) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETCALLBACKMASK, 0, 0)

	LVM_SETCALLBACKMASK = (LVM_FIRST + 11)
	// #define ListView_SetCallbackMask(hwnd, mask) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETCALLBACKMASK, (WPARAM)(UINT)(mask), 0)

	LVNI_ALL = 0

	LVNI_FOCUSED     = 0
	LVNI_SELECTED    = 0
	LVNI_CUT         = 0
	LVNI_DROPHILITED = 0
	LVNI_STATEMASK   = (LVNI_FOCUSED | LVNI_SELECTED | LVNI_CUT | LVNI_DROPHILITED)

	LVNI_VISIBLEORDER  = 0
	LVNI_PREVIOUS      = 0
	LVNI_VISIBLEONLY   = 0
	LVNI_SAMEGROUPONLY = 0

	LVNI_ABOVE         = 0
	LVNI_BELOW         = 0
	LVNI_TOLEFT        = 0
	LVNI_TORIGHT       = 0
	LVNI_DIRECTIONMASK = (LVNI_ABOVE | LVNI_BELOW | LVNI_TOLEFT | LVNI_TORIGHT)

	LVM_GETNEXTITEM = (LVM_FIRST + 12)
	// #define ListView_GetNextItem(hwnd, i, flags) \
	//     (int)SNDMSG((hwnd), LVM_GETNEXTITEM, (WPARAM)(int)(i), MAKELPARAM((flags), 0))

	LVFI_PARAM     = 0
	LVFI_STRING    = 0
	LVFI_SUBSTRING = 0 // Same as LVFI_PARTIAL
	LVFI_PARTIAL   = 0
	LVFI_WRAP      = 0
	LVFI_NEARESTXY = 0

	//LV_FINDINFOA = LVFINDINFOA
	//LV_FINDINFOW = LVFINDINFOW
	//
	//LV_FINDINFO = LVFINDINFO

	// typedef struct tagLVFINDINFOA
	// {
	//     UINT flags;
	//     LPCSTR psz;
	//     LPARAM lParam;
	//     POINT pt;
	//     UINT vkDirection;
	// } LVFINDINFOA, *LPFINDINFOA;

	// typedef struct tagLVFINDINFOW
	// {
	//     UINT flags;
	//     LPCWSTR psz;
	//     LPARAM lParam;
	//     POINT pt;
	//     UINT vkDirection;
	// } LVFINDINFOW, *LPFINDINFOW;

	//// #ifdef UNICODE
	//LVFINDINFO = LVFINDINFOW
	//// #else
	//LVFINDINFO = LVFINDINFOA
	//// #endif

	LVM_FINDITEMA = (LVM_FIRST + 13)
	LVM_FINDITEMW = (LVM_FIRST + 83)
	//// #ifdef UNICODE
	//LVM_FINDITEM = LVM_FINDITEMW
	//// #else
	//LVM_FINDITEM = LVM_FINDITEMA
	//// #endif

	// #define ListView_FindItem(hwnd, iStart, plvfi) \
	//     (int)SNDMSG((hwnd), LVM_FINDITEM, (WPARAM)(int)(iStart), (LPARAM)(const LV_FINDINFO *)(plvfi))

	LVIR_BOUNDS       = 0
	LVIR_ICON         = 1
	LVIR_LABEL        = 2
	LVIR_SELECTBOUNDS = 3

	LVM_GETITEMRECT = (LVM_FIRST + 14)
	// #define ListView_GetItemRect(hwnd, i, prc, code) \
	//      (BOOL)SNDMSG((hwnd), LVM_GETITEMRECT, (WPARAM)(int)(i), \
	//((prc) ? (((RECT *)(prc))->left = (code),(LPARAM)(RECT *)(prc)) : (LPARAM)(RECT *)NULL))

	LVM_SETITEMPOSITION = (LVM_FIRST + 15)
	// #define ListView_SetItemPosition(hwndLV, i, x, y) \
	//     (BOOL)SNDMSG((hwndLV), LVM_SETITEMPOSITION, (WPARAM)(int)(i), MAKELPARAM((x), (y)))

	LVM_GETITEMPOSITION = (LVM_FIRST + 16)
	// #define ListView_GetItemPosition(hwndLV, i, ppt) \
	//     (BOOL)SNDMSG((hwndLV), LVM_GETITEMPOSITION, (WPARAM)(int)(i), (LPARAM)(POINT *)(ppt))

	LVM_GETSTRINGWIDTHA = (LVM_FIRST + 17)
	LVM_GETSTRINGWIDTHW = (LVM_FIRST + 87)
	//// #ifdef UNICODE
	//LVM_GETSTRINGWIDTH = LVM_GETSTRINGWIDTHW
	//// #else
	//LVM_GETSTRINGWIDTH = LVM_GETSTRINGWIDTHA
	//// #endif

	// #define ListView_GetStringWidth(hwndLV, psz) \
	//     (int)SNDMSG((hwndLV), LVM_GETSTRINGWIDTH, 0, (LPARAM)(LPCTSTR)(psz))

	LVHT_NOWHERE         = 0
	LVHT_ONITEMICON      = 0
	LVHT_ONITEMLABEL     = 0
	LVHT_ONITEMSTATEICON = 0
	LVHT_ONITEM          = (LVHT_ONITEMICON | LVHT_ONITEMLABEL | LVHT_ONITEMSTATEICON)

	LVHT_ABOVE   = 0
	LVHT_BELOW   = 0
	LVHT_TORIGHT = 0
	LVHT_TOLEFT  = 0

	LVHT_EX_GROUP_HEADER     = 0
	LVHT_EX_GROUP_FOOTER     = 0
	LVHT_EX_GROUP_COLLAPSE   = 0
	LVHT_EX_GROUP_BACKGROUND = 0
	LVHT_EX_GROUP_STATEICON  = 0
	LVHT_EX_GROUP_SUBSETLINK = 0
	LVHT_EX_GROUP            = (LVHT_EX_GROUP_BACKGROUND | LVHT_EX_GROUP_COLLAPSE | LVHT_EX_GROUP_FOOTER | LVHT_EX_GROUP_HEADER | LVHT_EX_GROUP_STATEICON | LVHT_EX_GROUP_SUBSETLINK)
	LVHT_EX_ONCONTENTS       = 0 // On item AND not on the background
	LVHT_EX_FOOTER           = 0

	//LV_HITTESTINFO = LVHITTESTINFO
	//
	//LVHITTESTINFO_V1_SIZE = CCSIZEOF_STRUCT(LVHITTESTINFO, iItem)
	//
	// typedef struct tagLVHITTESTINFO
	// {
	//     POINT pt;
	//     UINT flags;
	//     int iItem;
	//     int iSubItem;    // this is was NOT in win95.  valid only for LVM_SUBITEMHITTEST
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     int iGroup; // readonly. index of group. only valid for owner data.
	//                 // supports single item in multiple groups.
	// #endif
	// } LVHITTESTINFO, *LPLVHITTESTINFO;

	LVM_HITTEST = (LVM_FIRST + 18)
	// #define ListView_HitTest(hwndLV, pinfo) \
	//     (int)SNDMSG((hwndLV), LVM_HITTEST, 0, (LPARAM)(LV_HITTESTINFO *)(pinfo))
	// #define ListView_HitTestEx(hwndLV, pinfo) \
	//     (int)SNDMSG((hwndLV), LVM_HITTEST, (WPARAM)-1, (LPARAM)(LV_HITTESTINFO *)(pinfo))

	LVM_ENSUREVISIBLE = (LVM_FIRST + 19)
	// #define ListView_EnsureVisible(hwndLV, i, fPartialOK) \
	//     (BOOL)SNDMSG((hwndLV), LVM_ENSUREVISIBLE, (WPARAM)(int)(i), MAKELPARAM((fPartialOK), 0))

	LVM_SCROLL = (LVM_FIRST + 20)
	// #define ListView_Scroll(hwndLV, dx, dy) \
	//     (BOOL)SNDMSG((hwndLV), LVM_SCROLL, (WPARAM)(int)(dx), (LPARAM)(int)(dy))

	LVM_REDRAWITEMS = (LVM_FIRST + 21)
	// #define ListView_RedrawItems(hwndLV, iFirst, iLast) \
	//     (BOOL)SNDMSG((hwndLV), LVM_REDRAWITEMS, (WPARAM)(int)(iFirst), (LPARAM)(int)(iLast))

	LVA_DEFAULT    = 0
	LVA_ALIGNLEFT  = 0
	LVA_ALIGNTOP   = 0
	LVA_SNAPTOGRID = 0

	LVM_ARRANGE = (LVM_FIRST + 22)
	// #define ListView_Arrange(hwndLV, code) \
	//     (BOOL)SNDMSG((hwndLV), LVM_ARRANGE, (WPARAM)(UINT)(code), 0L)

	LVM_EDITLABELA = (LVM_FIRST + 23)
	LVM_EDITLABELW = (LVM_FIRST + 118)
	//// #ifdef UNICODE
	//LVM_EDITLABEL = LVM_EDITLABELW
	//// #else
	//LVM_EDITLABEL = LVM_EDITLABELA
	//// #endif

	// #define ListView_EditLabel(hwndLV, i) \
	//     (HWND)SNDMSG((hwndLV), LVM_EDITLABEL, (WPARAM)(int)(i), 0L)

	LVM_GETEDITCONTROL = (LVM_FIRST + 24)
	// #define ListView_GetEditControl(hwndLV) \
	//     (HWND)SNDMSG((hwndLV), LVM_GETEDITCONTROL, 0, 0L)

	//LV_COLUMNA = LVCOLUMNA
	//LV_COLUMNW = LVCOLUMNW
	//
	//LV_COLUMN = LVCOLUMN
	//
	//LVCOLUMNA_V1_SIZE = CCSIZEOF_STRUCT(LVCOLUMNA, iSubItem)
	//LVCOLUMNW_V1_SIZE = CCSIZEOF_STRUCT(LVCOLUMNW, iSubItem)

	// typedef struct tagLVCOLUMNA
	// {
	//     UINT mask;
	//     int fmt;
	//     int cx;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iSubItem;
	//     int iImage;
	//     int iOrder;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     int cxMin;       // min snap point
	//     int cxDefault;   // default snap point
	//     int cxIdeal;     // read only. ideal may not eqaul current width if auto sized (LVS_EX_AUTOSIZECOLUMNS) to a lesser width.
	// #endif
	// } LVCOLUMNA, *LPLVCOLUMNA;

	// typedef struct tagLVCOLUMNW
	// {
	//     UINT mask;
	//     int fmt;
	//     int cx;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iSubItem;
	//     int iImage;
	//     int iOrder;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     int cxMin;       // min snap point
	//     int cxDefault;   // default snap point
	//     int cxIdeal;     // read only. ideal may not eqaul current width if auto sized (LVS_EX_AUTOSIZECOLUMNS) to a lesser width.
	// #endif
	// } LVCOLUMNW, *LPLVCOLUMNW;

	//// #ifdef UNICODE
	//LVCOLUMN = LVCOLUMNW
	//LPLVCOLUMN = LPLVCOLUMNW
	//LVCOLUMN_V1_SIZE = LVCOLUMNW_V1_SIZE
	//// #else
	//LVCOLUMN = LVCOLUMNA
	//LPLVCOLUMN = LPLVCOLUMNA
	//LVCOLUMN_V1_SIZE = LVCOLUMNA_V1_SIZE
	//// #endif

	LVCF_FMT     = 0
	LVCF_WIDTH   = 0
	LVCF_TEXT    = 0
	LVCF_SUBITEM = 0
	LVCF_IMAGE   = 0
	LVCF_ORDER   = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVCF_MINWIDTH     = 0
	LVCF_DEFAULTWIDTH = 0
	LVCF_IDEALWIDTH   = 0
	// #endif

	// // LVCFMT_ flags up to FFFF are shared with the header control (HDF_ flags).
	// // Flags above FFFF are listview-specific.

	LVCFMT_LEFT        = 0 // Same as HDF_LEFT
	LVCFMT_RIGHT       = 0 // Same as HDF_RIGHT
	LVCFMT_CENTER      = 0 // Same as HDF_CENTER
	LVCFMT_JUSTIFYMASK = 0 // Same as HDF_JUSTIFYMASK

	LVCFMT_IMAGE           = 0 // Same as HDF_IMAGE
	LVCFMT_BITMAP_ON_RIGHT = 0 // Same as HDF_BITMAP_ON_RIGHT
	LVCFMT_COL_HAS_IMAGES  = 0 // Same as HDF_OWNERDRAW

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVCFMT_FIXED_WIDTH  = 0 // Can't resize the column; same as HDF_FIXEDWIDTH
	LVCFMT_NO_DPI_SCALE = 0 // If not set, CCM_DPISCALE will govern scaling up fixed width
	LVCFMT_FIXED_RATIO  = 0 // Width will augment with the row height

	// // The following flags
	LVCFMT_LINE_BREAK         = 0 // Move to the top of the next list of columns
	LVCFMT_FILL               = 0 // Fill the remainder of the tile area. Might have a title.
	LVCFMT_WRAP               = 0 // This sub-item can be wrapped.
	LVCFMT_NO_TITLE           = 0 // This sub-item doesn't have an title.
	LVCFMT_TILE_PLACEMENTMASK = (LVCFMT_LINE_BREAK | LVCFMT_FILL)

	LVCFMT_SPLITBUTTON = 0 // Column is a split button; same as HDF_SPLITBUTTON
	// #endif

	LVM_GETCOLUMNA = (LVM_FIRST + 25)
	LVM_GETCOLUMNW = (LVM_FIRST + 95)
	//// #ifdef UNICODE
	//LVM_GETCOLUMN = LVM_GETCOLUMNW
	//// #else
	//LVM_GETCOLUMN = LVM_GETCOLUMNA
	//// #endif

	// #define ListView_GetColumn(hwnd, iCol, pcol) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETCOLUMN, (WPARAM)(int)(iCol), (LPARAM)(LV_COLUMN *)(pcol))

	LVM_SETCOLUMNA = (LVM_FIRST + 26)
	LVM_SETCOLUMNW = (LVM_FIRST + 96)
	//// #ifdef UNICODE
	//LVM_SETCOLUMN = LVM_SETCOLUMNW
	//// #else
	//LVM_SETCOLUMN = LVM_SETCOLUMNA
	//// #endif

	// #define ListView_SetColumn(hwnd, iCol, pcol) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETCOLUMN, (WPARAM)(int)(iCol), (LPARAM)(const LV_COLUMN *)(pcol))

	LVM_INSERTCOLUMNA = (LVM_FIRST + 27)
	LVM_INSERTCOLUMNW = (LVM_FIRST + 97)
	// #ifdef UNICODE
	// #   define  LVM_INSERTCOLUMN    LVM_INSERTCOLUMNW
	// #else
	// #   define  LVM_INSERTCOLUMN    LVM_INSERTCOLUMNA
	// #endif

	// #define ListView_InsertColumn(hwnd, iCol, pcol) \
	//     (int)SNDMSG((hwnd), LVM_INSERTCOLUMN, (WPARAM)(int)(iCol), (LPARAM)(const LV_COLUMN *)(pcol))

	LVM_DELETECOLUMN = (LVM_FIRST + 28)
	// #define ListView_DeleteColumn(hwnd, iCol) \
	//     (BOOL)SNDMSG((hwnd), LVM_DELETECOLUMN, (WPARAM)(int)(iCol), 0)

	LVM_GETCOLUMNWIDTH = (LVM_FIRST + 29)
	// #define ListView_GetColumnWidth(hwnd, iCol) \
	//     (int)SNDMSG((hwnd), LVM_GETCOLUMNWIDTH, (WPARAM)(int)(iCol), 0)

	// #define LVSCW_AUTOSIZE              -1
	// #define LVSCW_AUTOSIZE_USEHEADER    -2
	LVM_SETCOLUMNWIDTH = (LVM_FIRST + 30)

	// #define ListView_SetColumnWidth(hwnd, iCol, cx) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETCOLUMNWIDTH, (WPARAM)(int)(iCol), MAKELPARAM((cx), 0))

	LVM_GETHEADER = (LVM_FIRST + 31)
	// #define ListView_GetHeader(hwnd)\
	//     (HWND)SNDMSG((hwnd), LVM_GETHEADER, 0, 0L)

	LVM_CREATEDRAGIMAGE = (LVM_FIRST + 33)
	// #define ListView_CreateDragImage(hwnd, i, lpptUpLeft) \
	//     (HIMAGELIST)SNDMSG((hwnd), LVM_CREATEDRAGIMAGE, (WPARAM)(int)(i), (LPARAM)(LPPOINT)(lpptUpLeft))

	LVM_GETVIEWRECT = (LVM_FIRST + 34)
	// #define ListView_GetViewRect(hwnd, prc) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETVIEWRECT, 0, (LPARAM)(RECT *)(prc))

	LVM_GETTEXTCOLOR = (LVM_FIRST + 35)
	// #define ListView_GetTextColor(hwnd)  \
	//     (COLORREF)SNDMSG((hwnd), LVM_GETTEXTCOLOR, 0, 0L)

	LVM_SETTEXTCOLOR = (LVM_FIRST + 36)
	// #define ListView_SetTextColor(hwnd, clrText) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETTEXTCOLOR, 0, (LPARAM)(COLORREF)(clrText))

	LVM_GETTEXTBKCOLOR = (LVM_FIRST + 37)
	// #define ListView_GetTextBkColor(hwnd)  \
	//     (COLORREF)SNDMSG((hwnd), LVM_GETTEXTBKCOLOR, 0, 0L)

	LVM_SETTEXTBKCOLOR = (LVM_FIRST + 38)
	// #define ListView_SetTextBkColor(hwnd, clrTextBk) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETTEXTBKCOLOR, 0, (LPARAM)(COLORREF)(clrTextBk))

	LVM_GETTOPINDEX = (LVM_FIRST + 39)
	// #define ListView_GetTopIndex(hwndLV) \
	//     (int)SNDMSG((hwndLV), LVM_GETTOPINDEX, 0, 0)

	LVM_GETCOUNTPERPAGE = (LVM_FIRST + 40)
	// #define ListView_GetCountPerPage(hwndLV) \
	//     (int)SNDMSG((hwndLV), LVM_GETCOUNTPERPAGE, 0, 0)

	LVM_GETORIGIN = (LVM_FIRST + 41)
	// #define ListView_GetOrigin(hwndLV, ppt) \
	//     (BOOL)SNDMSG((hwndLV), LVM_GETORIGIN, (WPARAM)0, (LPARAM)(POINT *)(ppt))

	LVM_UPDATE = (LVM_FIRST + 42)
	// #define ListView_Update(hwndLV, i) \
	//     (BOOL)SNDMSG((hwndLV), LVM_UPDATE, (WPARAM)(i), 0L)

	LVM_SETITEMSTATE = (LVM_FIRST + 43)
	// #define ListView_SetItemState(hwndLV, i, data, mask) \
	// { LV_ITEM _macro_lvi;\
	//_macro_lvi.stateMask = (mask);\
	//_macro_lvi.state = (data);\
	//   SNDMSG((hwndLV), LVM_SETITEMSTATE, (WPARAM)(i), (LPARAM)(LV_ITEM *)&_macro_lvi);\
	// }

	// #define ListView_SetCheckState(hwndLV, i, fCheck) \
	//   ListView_SetItemState(hwndLV, i, INDEXTOSTATEIMAGEMASK((fCheck)?2:1), LVIS_STATEIMAGEMASK)

	LVM_GETITEMSTATE = (LVM_FIRST + 44)
	// #define ListView_GetItemState(hwndLV, i, mask) \
	//    (UINT)SNDMSG((hwndLV), LVM_GETITEMSTATE, (WPARAM)(i), (LPARAM)(mask))

	// #define ListView_GetCheckState(hwndLV, i) \
	//    ((((UINT)(SNDMSG((hwndLV), LVM_GETITEMSTATE, (WPARAM)(i), LVIS_STATEIMAGEMASK))) >> 12) -1)

	LVM_GETITEMTEXTA = (LVM_FIRST + 45)
	LVM_GETITEMTEXTW = (LVM_FIRST + 115)

	//// #ifdef UNICODE
	//LVM_GETITEMTEXT = LVM_GETITEMTEXTW
	//// #else
	//LVM_GETITEMTEXT = LVM_GETITEMTEXTA
	//// #endif

	// #define ListView_GetItemText(hwndLV, i, iSubItem_, pszText_, cchTextMax_) \
	// { LV_ITEM _macro_lvi;\
	//_macro_lvi.iSubItem = (iSubItem_);\
	//_macro_lvi.cchTextMax = (cchTextMax_);\
	//_macro_lvi.pszText = (pszText_);\
	//   SNDMSG((hwndLV), LVM_GETITEMTEXT, (WPARAM)(i), (LPARAM)(LV_ITEM *)&_macro_lvi);\
	// }

	LVM_SETITEMTEXTA = (LVM_FIRST + 46)
	LVM_SETITEMTEXTW = (LVM_FIRST + 116)

	//// #ifdef UNICODE
	//LVM_SETITEMTEXT = LVM_SETITEMTEXTW
	//// #else
	//LVM_SETITEMTEXT = LVM_SETITEMTEXTA
	//// #endif

	// #define ListView_SetItemText(hwndLV, i, iSubItem_, pszText_) \
	// { LV_ITEM _macro_lvi;\
	//_macro_lvi.iSubItem = (iSubItem_);\
	//_macro_lvi.pszText = (pszText_);\
	//   SNDMSG((hwndLV), LVM_SETITEMTEXT, (WPARAM)(i), (LPARAM)(LV_ITEM *)&_macro_lvi);\
	// }

	// // these flags only apply to LVS_OWNERDATA listviews in report or list mode
	LVSICF_NOINVALIDATEALL = 0
	LVSICF_NOSCROLL        = 0

	LVM_SETITEMCOUNT = (LVM_FIRST + 47)
	// #define ListView_SetItemCount(hwndLV, cItems) \
	//   SNDMSG((hwndLV), LVM_SETITEMCOUNT, (WPARAM)(cItems), 0)

	// #define ListView_SetItemCountEx(hwndLV, cItems, dwFlags) \
	//   SNDMSG((hwndLV), LVM_SETITEMCOUNT, (WPARAM)(cItems), (LPARAM)(dwFlags))

	// typedef int (CALLBACK *PFNLVCOMPARE)(LPARAM, LPARAM, LPARAM);

	LVM_SORTITEMS = (LVM_FIRST + 48)
	// #define ListView_SortItems(hwndLV, _pfnCompare, _lPrm) \
	//   (BOOL)SNDMSG((hwndLV), LVM_SORTITEMS, (WPARAM)(LPARAM)(_lPrm), \
	//   (LPARAM)(PFNLVCOMPARE)(_pfnCompare))

	LVM_SETITEMPOSITION32 = (LVM_FIRST + 49)
	// #define ListView_SetItemPosition32(hwndLV, i, x0, y0) \
	// {   POINT ptNewPos; \
	//ptNewPos.x = (x0); ptNewPos.y = (y0); \
	//     SNDMSG((hwndLV), LVM_SETITEMPOSITION32, (WPARAM)(int)(i), (LPARAM)&ptNewPos); \
	// }

	LVM_GETSELECTEDCOUNT = (LVM_FIRST + 50)
	// #define ListView_GetSelectedCount(hwndLV) \
	//     (UINT)SNDMSG((hwndLV), LVM_GETSELECTEDCOUNT, 0, 0L)

	LVM_GETITEMSPACING = (LVM_FIRST + 51)
	// #define ListView_GetItemSpacing(hwndLV, fSmall) \
	//         (DWORD)SNDMSG((hwndLV), LVM_GETITEMSPACING, fSmall, 0L)

	LVM_GETISEARCHSTRINGA = (LVM_FIRST + 52)
	LVM_GETISEARCHSTRINGW = (LVM_FIRST + 117)

	//// #ifdef UNICODE
	//LVM_GETISEARCHSTRING = LVM_GETISEARCHSTRINGW
	//// #else
	//LVM_GETISEARCHSTRING = LVM_GETISEARCHSTRINGA
	//// #endif

	// #define ListView_GetISearchString(hwndLV, lpsz) \
	//         (BOOL)SNDMSG((hwndLV), LVM_GETISEARCHSTRING, 0, (LPARAM)(LPTSTR)(lpsz))

	LVM_SETICONSPACING = (LVM_FIRST + 53)
	// // -1 for cx and cy means we'll use the default (system settings)
	// // 0 for cx or cy means use the current setting (allows you to change just one param)
	// #define ListView_SetIconSpacing(hwndLV, cx, cy) \
	//         (DWORD)SNDMSG((hwndLV), LVM_SETICONSPACING, 0, MAKELONG(cx,cy))

	LVM_SETEXTENDEDLISTVIEWSTYLE = (LVM_FIRST + 54) // optional wParam == mask
	// #define ListView_SetExtendedListViewStyle(hwndLV, dw)\
	//         (DWORD)SNDMSG((hwndLV), LVM_SETEXTENDEDLISTVIEWSTYLE, 0, dw)
	// #define ListView_SetExtendedListViewStyleEx(hwndLV, dwMask, dw)\
	//         (DWORD)SNDMSG((hwndLV), LVM_SETEXTENDEDLISTVIEWSTYLE, dwMask, dw)

	LVM_GETEXTENDEDLISTVIEWSTYLE = (LVM_FIRST + 55)
	// #define ListView_GetExtendedListViewStyle(hwndLV)\
	//         (DWORD)SNDMSG((hwndLV), LVM_GETEXTENDEDLISTVIEWSTYLE, 0, 0)

	LVS_EX_GRIDLINES        = 0
	LVS_EX_SUBITEMIMAGES    = 0
	LVS_EX_CHECKBOXES       = 0
	LVS_EX_TRACKSELECT      = 0
	LVS_EX_HEADERDRAGDROP   = 0
	LVS_EX_FULLROWSELECT    = 0 // applies to report mode only
	LVS_EX_ONECLICKACTIVATE = 0
	LVS_EX_TWOCLICKACTIVATE = 0
	LVS_EX_FLATSB           = 0
	LVS_EX_REGIONAL         = 0
	LVS_EX_INFOTIP          = 0 // listview does InfoTips for you
	LVS_EX_UNDERLINEHOT     = 0
	LVS_EX_UNDERLINECOLD    = 0
	LVS_EX_MULTIWORKAREAS   = 0
	LVS_EX_LABELTIP         = 0 // listview unfolds partly hidden labels if it does not have infotip text
	LVS_EX_BORDERSELECT     = 0 // border selection style instead of highlight
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	LVS_EX_DOUBLEBUFFER = 0
	LVS_EX_HIDELABELS   = 0
	LVS_EX_SINGLEROW    = 0
	LVS_EX_SNAPTOGRID   = 0 // Icons automatically snap to grid.
	LVS_EX_SIMPLESELECT = 0 // Also changes overlay rendering to top right for icon mode.
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVS_EX_JUSTIFYCOLUMNS        = 0 // Icons are lined up in columns that use up the whole view area.
	LVS_EX_TRANSPARENTBKGND      = 0 // Background is painted by the parent via WM_PRINTCLIENT
	LVS_EX_TRANSPARENTSHADOWTEXT = 0 // Enable shadow text on transparent backgrounds only (useful with bitmaps)
	LVS_EX_AUTOAUTOARRANGE       = 0 // Icons automatically arrange if no icon positions have been set
	LVS_EX_HEADERINALLVIEWS      = 0 // Display column header in all view modes
	LVS_EX_AUTOCHECKSELECT       = 0
	LVS_EX_AUTOSIZECOLUMNS       = 0
	LVS_EX_COLUMNSNAPPOINTS      = 0
	LVS_EX_COLUMNOVERFLOW        = 0

	// #endif

	LVM_GETSUBITEMRECT = (LVM_FIRST + 56)
	// #define ListView_GetSubItemRect(hwnd, iItem, iSubItem, code, prc) \
	//         (BOOL)SNDMSG((hwnd), LVM_GETSUBITEMRECT, (WPARAM)(int)(iItem), \
	//((prc) ? ((((LPRECT)(prc))->top = (iSubItem)), (((LPRECT)(prc))->left = (code)), (LPARAM)(prc)) : (LPARAM)(LPRECT)NULL))

	LVM_SUBITEMHITTEST = (LVM_FIRST + 57)
	// #define ListView_SubItemHitTest(hwnd, plvhti) \
	//         (int)SNDMSG((hwnd), LVM_SUBITEMHITTEST, 0, (LPARAM)(LPLVHITTESTINFO)(plvhti))
	// #define ListView_SubItemHitTestEx(hwnd, plvhti) \
	//         (int)SNDMSG((hwnd), LVM_SUBITEMHITTEST, (WPARAM)-1, (LPARAM)(LPLVHITTESTINFO)(plvhti))

	LVM_SETCOLUMNORDERARRAY = (LVM_FIRST + 58)
	// #define ListView_SetColumnOrderArray(hwnd, iCount, pi) \
	//         (BOOL)SNDMSG((hwnd), LVM_SETCOLUMNORDERARRAY, (WPARAM)(iCount), (LPARAM)(LPINT)(pi))

	LVM_GETCOLUMNORDERARRAY = (LVM_FIRST + 59)
	// #define ListView_GetColumnOrderArray(hwnd, iCount, pi) \
	//         (BOOL)SNDMSG((hwnd), LVM_GETCOLUMNORDERARRAY, (WPARAM)(iCount), (LPARAM)(LPINT)(pi))

	LVM_SETHOTITEM = (LVM_FIRST + 60)
	// #define ListView_SetHotItem(hwnd, i) \
	//         (int)SNDMSG((hwnd), LVM_SETHOTITEM, (WPARAM)(i), 0)

	LVM_GETHOTITEM = (LVM_FIRST + 61)
	// #define ListView_GetHotItem(hwnd) \
	//         (int)SNDMSG((hwnd), LVM_GETHOTITEM, 0, 0)

	LVM_SETHOTCURSOR = (LVM_FIRST + 62)
	// #define ListView_SetHotCursor(hwnd, hcur) \
	//         (HCURSOR)SNDMSG((hwnd), LVM_SETHOTCURSOR, 0, (LPARAM)(hcur))

	LVM_GETHOTCURSOR = (LVM_FIRST + 63)
	// #define ListView_GetHotCursor(hwnd) \
	//         (HCURSOR)SNDMSG((hwnd), LVM_GETHOTCURSOR, 0, 0)

	LVM_APPROXIMATEVIEWRECT = (LVM_FIRST + 64)
	// #define ListView_ApproximateViewRect(hwnd, iWidth, iHeight, iCount) \
	//         (DWORD)SNDMSG((hwnd), LVM_APPROXIMATEVIEWRECT, (WPARAM)(iCount), MAKELPARAM(iWidth, iHeight))

	LV_MAX_WORKAREAS = 16
	LVM_SETWORKAREAS = (LVM_FIRST + 65)
	// #define ListView_SetWorkAreas(hwnd, nWorkAreas, prc) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETWORKAREAS, (WPARAM)(int)(nWorkAreas), (LPARAM)(RECT *)(prc))

	LVM_GETWORKAREAS = (LVM_FIRST + 70)
	// #define ListView_GetWorkAreas(hwnd, nWorkAreas, prc) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETWORKAREAS, (WPARAM)(int)(nWorkAreas), (LPARAM)(RECT *)(prc))

	LVM_GETNUMBEROFWORKAREAS = (LVM_FIRST + 73)
	// #define ListView_GetNumberOfWorkAreas(hwnd, pnWorkAreas) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETNUMBEROFWORKAREAS, 0, (LPARAM)(UINT *)(pnWorkAreas))

	LVM_GETSELECTIONMARK = (LVM_FIRST + 66)
	// #define ListView_GetSelectionMark(hwnd) \
	//     (int)SNDMSG((hwnd), LVM_GETSELECTIONMARK, 0, 0)

	LVM_SETSELECTIONMARK = (LVM_FIRST + 67)
	// #define ListView_SetSelectionMark(hwnd, i) \
	//     (int)SNDMSG((hwnd), LVM_SETSELECTIONMARK, 0, (LPARAM)(i))

	LVM_SETHOVERTIME = (LVM_FIRST + 71)
	// #define ListView_SetHoverTime(hwndLV, dwHoverTimeMs)\
	//         (DWORD)SNDMSG((hwndLV), LVM_SETHOVERTIME, 0, (LPARAM)(dwHoverTimeMs))

	LVM_GETHOVERTIME = (LVM_FIRST + 72)
	// #define ListView_GetHoverTime(hwndLV)\
	//         (DWORD)SNDMSG((hwndLV), LVM_GETHOVERTIME, 0, 0)

	LVM_SETTOOLTIPS = (LVM_FIRST + 74)
	// #define ListView_SetToolTips(hwndLV, hwndNewHwnd)\
	//         (HWND)SNDMSG((hwndLV), LVM_SETTOOLTIPS, (WPARAM)(hwndNewHwnd), 0)

	LVM_GETTOOLTIPS = (LVM_FIRST + 78)
	// #define ListView_GetToolTips(hwndLV)\
	//         (HWND)SNDMSG((hwndLV), LVM_GETTOOLTIPS, 0, 0)

	LVM_SORTITEMSEX = (LVM_FIRST + 81)
	// #define ListView_SortItemsEx(hwndLV, _pfnCompare, _lPrm) \
	//   (BOOL)SNDMSG((hwndLV), LVM_SORTITEMSEX, (WPARAM)(LPARAM)(_lPrm), (LPARAM)(PFNLVCOMPARE)(_pfnCompare))

	// typedef struct tagLVBKIMAGEA
	// {
	//     ULONG ulFlags;              // LVBKIF_*
	//     HBITMAP hbm;
	//     LPSTR pszImage;
	//     UINT cchImageMax;
	//     int xOffsetPercent;
	//     int yOffsetPercent;
	// } LVBKIMAGEA, *LPLVBKIMAGEA;
	// typedef struct tagLVBKIMAGEW
	// {
	//     ULONG ulFlags;              // LVBKIF_*
	//     HBITMAP hbm;
	//     LPWSTR pszImage;
	//     UINT cchImageMax;
	//     int xOffsetPercent;
	//     int yOffsetPercent;
	// } LVBKIMAGEW, *LPLVBKIMAGEW;

	LVBKIF_SOURCE_NONE    = 0
	LVBKIF_SOURCE_HBITMAP = 0
	LVBKIF_SOURCE_URL     = 0
	LVBKIF_SOURCE_MASK    = 0
	LVBKIF_STYLE_NORMAL   = 0
	LVBKIF_STYLE_TILE     = 0
	LVBKIF_STYLE_MASK     = 0
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	LVBKIF_FLAG_TILEOFFSET = 0
	LVBKIF_TYPE_WATERMARK  = 0
	LVBKIF_FLAG_ALPHABLEND = 0
	// #endif

	LVM_SETBKIMAGEA = (LVM_FIRST + 68)
	LVM_SETBKIMAGEW = (LVM_FIRST + 138)
	LVM_GETBKIMAGEA = (LVM_FIRST + 69)
	LVM_GETBKIMAGEW = (LVM_FIRST + 139)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	LVM_SETSELECTEDCOLUMN = (LVM_FIRST + 140)
	// #define ListView_SetSelectedColumn(hwnd, iCol) \
	//     SNDMSG((hwnd), LVM_SETSELECTEDCOLUMN, (WPARAM)(iCol), 0)

	LV_VIEW_ICON      = 0
	LV_VIEW_DETAILS   = 0
	LV_VIEW_SMALLICON = 0
	LV_VIEW_LIST      = 0
	LV_VIEW_TILE      = 0
	LV_VIEW_MAX       = 0

	LVM_SETVIEW = (LVM_FIRST + 142)
	// #define ListView_SetView(hwnd, iView) \
	//     (DWORD)SNDMSG((hwnd), LVM_SETVIEW, (WPARAM)(DWORD)(iView), 0)

	LVM_GETVIEW = (LVM_FIRST + 143)
	// #define ListView_GetView(hwnd) \
	//     (DWORD)SNDMSG((hwnd), LVM_GETVIEW, 0, 0)

	LVGF_NONE    = 0
	LVGF_HEADER  = 0
	LVGF_FOOTER  = 0
	LVGF_STATE   = 0
	LVGF_ALIGN   = 0
	LVGF_GROUPID = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVGF_SUBTITLE          = 0 // pszSubtitle is valid
	LVGF_TASK              = 0 // pszTask is valid
	LVGF_DESCRIPTIONTOP    = 0 // pszDescriptionTop is valid
	LVGF_DESCRIPTIONBOTTOM = 0 // pszDescriptionBottom is valid
	LVGF_TITLEIMAGE        = 0 // iTitleImage is valid
	LVGF_EXTENDEDIMAGE     = 0 // iExtendedImage is valid
	LVGF_ITEMS             = 0 // iFirstItem and cItems are valid
	LVGF_SUBSET            = 0 // pszSubsetTitle is valid
	LVGF_SUBSETITEMS       = 0 // readonly, cItems holds count of items in visible subset, iFirstItem is valid
	// #endif

	LVGS_NORMAL            = 0
	LVGS_COLLAPSED         = 0
	LVGS_HIDDEN            = 0
	LVGS_NOHEADER          = 0
	LVGS_COLLAPSIBLE       = 0
	LVGS_FOCUSED           = 0
	LVGS_SELECTED          = 0
	LVGS_SUBSETED          = 0
	LVGS_SUBSETLINKFOCUSED = 0

	LVGA_HEADER_LEFT   = 0
	LVGA_HEADER_CENTER = 0
	LVGA_HEADER_RIGHT  = 0 // Don't forget to validate exclusivity
	LVGA_FOOTER_LEFT   = 0
	LVGA_FOOTER_CENTER = 0
	LVGA_FOOTER_RIGHT  = 0 // Don't forget to validate exclusivity

	// typedef struct tagLVGROUP
	// {
	//     UINT    cbSize;
	//     UINT    mask;
	//     LPWSTR  pszHeader;
	//     int     cchHeader;

	//     LPWSTR  pszFooter;
	//     int     cchFooter;

	//     int     iGroupId;

	//     UINT    stateMask;
	//     UINT    state;
	//     UINT    uAlign;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     LPWSTR  pszSubtitle;
	//     UINT    cchSubtitle;
	//     LPWSTR  pszTask;
	//     UINT    cchTask;
	//     LPWSTR  pszDescriptionTop;
	//     UINT    cchDescriptionTop;
	//     LPWSTR  pszDescriptionBottom;
	//     UINT    cchDescriptionBottom;
	//     int     iTitleImage;
	//     int     iExtendedImage;
	//     int     iFirstItem;         // Read only
	//     UINT    cItems;             // Read only
	//     LPWSTR  pszSubsetTitle;     // NULL if group is not subset
	//     UINT    cchSubsetTitle;

	//LVGROUP_V5_SIZE = CCSIZEOF_STRUCT(LVGROUP, uAlign)

	// #endif
	// } LVGROUP, *PLVGROUP;

	LVM_INSERTGROUP = (LVM_FIRST + 145)
	// #define ListView_InsertGroup(hwnd, index, pgrp) \
	//     SNDMSG((hwnd), LVM_INSERTGROUP, (WPARAM)(index), (LPARAM)(pgrp))

	LVM_SETGROUPINFO = (LVM_FIRST + 147)
	// #define ListView_SetGroupInfo(hwnd, iGroupId, pgrp) \
	//     SNDMSG((hwnd), LVM_SETGROUPINFO, (WPARAM)(iGroupId), (LPARAM)(pgrp))

	LVM_GETGROUPINFO = (LVM_FIRST + 149)
	// #define ListView_GetGroupInfo(hwnd, iGroupId, pgrp) \
	//     SNDMSG((hwnd), LVM_GETGROUPINFO, (WPARAM)(iGroupId), (LPARAM)(pgrp))

	LVM_REMOVEGROUP = (LVM_FIRST + 150)
	// #define ListView_RemoveGroup(hwnd, iGroupId) \
	//     SNDMSG((hwnd), LVM_REMOVEGROUP, (WPARAM)(iGroupId), 0)

	LVM_MOVEGROUP = (LVM_FIRST + 151)
	// #define ListView_MoveGroup(hwnd, iGroupId, toIndex) \
	//     SNDMSG((hwnd), LVM_MOVEGROUP, (WPARAM)(iGroupId), (LPARAM)(toIndex))

	LVM_GETGROUPCOUNT = (LVM_FIRST + 152)
	// #define ListView_GetGroupCount(hwnd) \
	//     SNDMSG((hwnd), LVM_GETGROUPCOUNT, (WPARAM)0, (LPARAM)0)

	LVM_GETGROUPINFOBYINDEX = (LVM_FIRST + 153)
	// #define ListView_GetGroupInfoByIndex(hwnd, iIndex, pgrp) \
	//     SNDMSG((hwnd), LVM_GETGROUPINFOBYINDEX, (WPARAM)(iIndex), (LPARAM)(pgrp))

	LVM_MOVEITEMTOGROUP = (LVM_FIRST + 154)
	// #define ListView_MoveItemToGroup(hwnd, idItemFrom, idGroupTo) \
	//     SNDMSG((hwnd), LVM_MOVEITEMTOGROUP, (WPARAM)(idItemFrom), (LPARAM)(idGroupTo))

	LVGGR_GROUP      = 0 // Entire expanded group
	LVGGR_HEADER     = 1 // Header only (collapsed group)
	LVGGR_LABEL      = 2 // Label only
	LVGGR_SUBSETLINK = 3 // subset link only

	LVM_GETGROUPRECT = (LVM_FIRST + 98)
	// #define ListView_GetGroupRect(hwnd, iGroupId, type, prc) \
	//     SNDMSG((hwnd), LVM_GETGROUPRECT, (WPARAM)(iGroupId), \
	//((prc) ? (((RECT*)(prc))->top = (type)), (LPARAM)(RECT*)(prc) : (LPARAM)(RECT*)NULL))

	LVGMF_NONE        = 0
	LVGMF_BORDERSIZE  = 0
	LVGMF_BORDERCOLOR = 0
	LVGMF_TEXTCOLOR   = 0

	// typedef struct tagLVGROUPMETRICS
	// {
	//     UINT cbSize;
	//     UINT mask;
	//     UINT Left;
	//     UINT Top;
	//     UINT Right;
	//     UINT Bottom;
	//     COLORREF crLeft;
	//     COLORREF crTop;
	//     COLORREF crRight;
	//     COLORREF crBottom;
	//     COLORREF crHeader;
	//     COLORREF crFooter;
	// } LVGROUPMETRICS, *PLVGROUPMETRICS;

	LVM_SETGROUPMETRICS = (LVM_FIRST + 155)
	// #define ListView_SetGroupMetrics(hwnd, pGroupMetrics) \
	//     SNDMSG((hwnd), LVM_SETGROUPMETRICS, 0, (LPARAM)(pGroupMetrics))

	LVM_GETGROUPMETRICS = (LVM_FIRST + 156)
	// #define ListView_GetGroupMetrics(hwnd, pGroupMetrics) \
	//     SNDMSG((hwnd), LVM_GETGROUPMETRICS, 0, (LPARAM)(pGroupMetrics))

	LVM_ENABLEGROUPVIEW = (LVM_FIRST + 157)
	// #define ListView_EnableGroupView(hwnd, fEnable) \
	//     SNDMSG((hwnd), LVM_ENABLEGROUPVIEW, (WPARAM)(fEnable), 0)

	// typedef int (CALLBACK *PFNLVGROUPCOMPARE)(int, int, void *);

	LVM_SORTGROUPS = (LVM_FIRST + 158)
	// #define ListView_SortGroups(hwnd, _pfnGroupCompate, _plv) \
	//     SNDMSG((hwnd), LVM_SORTGROUPS, (WPARAM)(_pfnGroupCompate), (LPARAM)(_plv))

	// typedef struct tagLVINSERTGROUPSORTED
	// {
	//     PFNLVGROUPCOMPARE pfnGroupCompare;
	//     void *pvData;
	//     LVGROUP lvGroup;
	// }LVINSERTGROUPSORTED, *PLVINSERTGROUPSORTED;

	LVM_INSERTGROUPSORTED = (LVM_FIRST + 159)
	// #define ListView_InsertGroupSorted(hwnd, structInsert) \
	//     SNDMSG((hwnd), LVM_INSERTGROUPSORTED, (WPARAM)(structInsert), 0)

	LVM_REMOVEALLGROUPS = (LVM_FIRST + 160)
	// #define ListView_RemoveAllGroups(hwnd) \
	//     SNDMSG((hwnd), LVM_REMOVEALLGROUPS, 0, 0)

	LVM_HASGROUP = (LVM_FIRST + 161)
	// #define ListView_HasGroup(hwnd, dwGroupId) \
	//     SNDMSG((hwnd), LVM_HASGROUP, dwGroupId, 0)

	// #define ListView_SetGroupState(hwnd, dwGroupId, dwMask, dwState) \
	// { LVGROUP _macro_lvg;\
	//_macro_lvg.cbSize = sizeof(_macro_lvg);\
	//_macro_lvg.mask = LVGF_STATE;\
	//_macro_lvg.stateMask = dwMask;\
	//_macro_lvg.state = dwState;\
	//   SNDMSG((hwnd), LVM_SETGROUPINFO, (WPARAM)(dwGroupId), (LPARAM)(LVGROUP *)&_macro_lvg);\
	// }
	LVM_GETGROUPSTATE = (LVM_FIRST + 92)
	// #define ListView_GetGroupState(hwnd, dwGroupId, dwMask) \
	//   (UINT) SNDMSG((hwnd), LVM_GETGROUPSTATE, (WPARAM)(dwGroupId), (LPARAM)(dwMask))

	LVM_GETFOCUSEDGROUP = (LVM_FIRST + 93)
	// #define ListView_GetFocusedGroup(hwnd) \
	//     SNDMSG((hwnd), LVM_GETFOCUSEDGROUP, 0, 0)

	LVTVIF_AUTOSIZE    = 0
	LVTVIF_FIXEDWIDTH  = 0
	LVTVIF_FIXEDHEIGHT = 0
	LVTVIF_FIXEDSIZE   = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVTVIF_EXTENDED = 0
	// #endif

	LVTVIM_TILESIZE    = 0
	LVTVIM_COLUMNS     = 0
	LVTVIM_LABELMARGIN = 0

	// typedef struct tagLVTILEVIEWINFO
	// {
	//     UINT    cbSize;
	//     DWORD   dwMask;     //LVTVIM_*
	//     DWORD   dwFlags;    //LVTVIF_*
	//     SIZE    sizeTile;
	//     int     cLines;
	//     RECT    rcLabelMargin;
	// } LVTILEVIEWINFO, *PLVTILEVIEWINFO;

	// typedef struct tagLVTILEINFO
	// {
	//     UINT    cbSize;
	//     int     iItem;
	//     UINT    cColumns;
	//     PUINT   puColumns;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//     int*    piColFmt;
	// #endif
	// } LVTILEINFO, *PLVTILEINFO;

	//LVTILEINFO_V5_SIZE = CCSIZEOF_STRUCT(LVTILEINFO, puColumns)

	LVM_SETTILEVIEWINFO = (LVM_FIRST + 162)
	// #define ListView_SetTileViewInfo(hwnd, ptvi) \
	//     SNDMSG((hwnd), LVM_SETTILEVIEWINFO, 0, (LPARAM)(ptvi))

	LVM_GETTILEVIEWINFO = (LVM_FIRST + 163)
	// #define ListView_GetTileViewInfo(hwnd, ptvi) \
	//     SNDMSG((hwnd), LVM_GETTILEVIEWINFO, 0, (LPARAM)(ptvi))

	LVM_SETTILEINFO = (LVM_FIRST + 164)
	// #define ListView_SetTileInfo(hwnd, pti) \
	//     SNDMSG((hwnd), LVM_SETTILEINFO, 0, (LPARAM)(pti))

	LVM_GETTILEINFO = (LVM_FIRST + 165)
	// #define ListView_GetTileInfo(hwnd, pti) \
	//     SNDMSG((hwnd), LVM_GETTILEINFO, 0, (LPARAM)(pti))

	// typedef struct
	// {
	//     UINT cbSize;
	//     DWORD dwFlags;
	//     int iItem;
	//     DWORD dwReserved;
	// } LVINSERTMARK, * LPLVINSERTMARK;

	LVIM_AFTER = 0 // TRUE = insert After iItem, otherwise before

	LVM_SETINSERTMARK = (LVM_FIRST + 166)
	// #define ListView_SetInsertMark(hwnd, lvim) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETINSERTMARK, (WPARAM) 0, (LPARAM) (lvim))

	LVM_GETINSERTMARK = (LVM_FIRST + 167)
	// #define ListView_GetInsertMark(hwnd, lvim) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETINSERTMARK, (WPARAM) 0, (LPARAM) (lvim))

	LVM_INSERTMARKHITTEST = (LVM_FIRST + 168)
	// #define ListView_InsertMarkHitTest(hwnd, point, lvim) \
	//     (int)SNDMSG((hwnd), LVM_INSERTMARKHITTEST, (WPARAM)(LPPOINT)(point), (LPARAM)(LPLVINSERTMARK)(lvim))

	LVM_GETINSERTMARKRECT = (LVM_FIRST + 169)
	// #define ListView_GetInsertMarkRect(hwnd, rc) \
	//     (int)SNDMSG((hwnd), LVM_GETINSERTMARKRECT, (WPARAM)0, (LPARAM)(LPRECT)(rc))

	LVM_SETINSERTMARKCOLOR = (LVM_FIRST + 170)
	// #define ListView_SetInsertMarkColor(hwnd, color) \
	//     (COLORREF)SNDMSG((hwnd), LVM_SETINSERTMARKCOLOR, (WPARAM)0, (LPARAM)(COLORREF)(color))

	LVM_GETINSERTMARKCOLOR = (LVM_FIRST + 171)
	// #define ListView_GetInsertMarkColor(hwnd) \
	//     (COLORREF)SNDMSG((hwnd), LVM_GETINSERTMARKCOLOR, (WPARAM)0, (LPARAM)0)

	// typedef struct tagLVSETINFOTIP
	// {
	//     UINT cbSize;
	//     DWORD dwFlags;
	//     LPWSTR pszText;
	//     int iItem;
	//     int iSubItem;
	// } LVSETINFOTIP, *PLVSETINFOTIP;

	LVM_SETINFOTIP = (LVM_FIRST + 173)

	// #define ListView_SetInfoTip(hwndLV, plvInfoTip)\
	//         (BOOL)SNDMSG((hwndLV), LVM_SETINFOTIP, (WPARAM)0, (LPARAM)(plvInfoTip))

	LVM_GETSELECTEDCOLUMN = (LVM_FIRST + 174)
	// #define ListView_GetSelectedColumn(hwnd) \
	//     (UINT)SNDMSG((hwnd), LVM_GETSELECTEDCOLUMN, 0, 0)

	LVM_ISGROUPVIEWENABLED = (LVM_FIRST + 175)
	// #define ListView_IsGroupViewEnabled(hwnd) \
	//     (BOOL)SNDMSG((hwnd), LVM_ISGROUPVIEWENABLED, 0, 0)

	LVM_GETOUTLINECOLOR = (LVM_FIRST + 176)
	// #define ListView_GetOutlineColor(hwnd) \
	//     (COLORREF)SNDMSG((hwnd), LVM_GETOUTLINECOLOR, 0, 0)

	LVM_SETOUTLINECOLOR = (LVM_FIRST + 177)
	// #define ListView_SetOutlineColor(hwnd, color) \
	//     (COLORREF)SNDMSG((hwnd), LVM_SETOUTLINECOLOR, (WPARAM)0, (LPARAM)(COLORREF)(color))

	LVM_CANCELEDITLABEL = (LVM_FIRST + 179)
	// #define ListView_CancelEditLabel(hwnd) \
	//     (VOID)SNDMSG((hwnd), LVM_CANCELEDITLABEL, (WPARAM)0, (LPARAM)0)

	// // These next to methods make it easy to identify an item that can be repositioned
	// // within listview. For example: Many developers use the lParam to store an identifier that is
	// // unique. Unfortunatly, in order to find this item, they have to iterate through all of the items
	// // in the listview. Listview will maintain a unique identifier.  The upper bound is the size of a DWORD.
	LVM_MAPINDEXTOID = (LVM_FIRST + 180)
	// #define ListView_MapIndexToID(hwnd, index) \
	//     (UINT)SNDMSG((hwnd), LVM_MAPINDEXTOID, (WPARAM)(index), (LPARAM)0)

	LVM_MAPIDTOINDEX = (LVM_FIRST + 181)
	// #define ListView_MapIDToIndex(hwnd, id) \
	//     (UINT)SNDMSG((hwnd), LVM_MAPIDTOINDEX, (WPARAM)(id), (LPARAM)0)

	LVM_ISITEMVISIBLE = (LVM_FIRST + 182)
	// #define ListView_IsItemVisible(hwnd, index) \
	//     (UINT)SNDMSG((hwnd), LVM_ISITEMVISIBLE, (WPARAM)(index), (LPARAM)0)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	// #define ListView_SetGroupHeaderImageList(hwnd, himl) \
	//     (HIMAGELIST)SNDMSG((hwnd), LVM_SETIMAGELIST, (WPARAM)LVSIL_GROUPHEADER, (LPARAM)(HIMAGELIST)(himl))

	// #define ListView_GetGroupHeaderImageList(hwnd) \
	//     (HIMAGELIST)SNDMSG((hwnd), LVM_GETIMAGELIST, (WPARAM)LVSIL_GROUPHEADER, 0L)

	LVM_GETEMPTYTEXT = (LVM_FIRST + 204)
	// #define ListView_GetEmptyText(hwnd, pszText, cchText) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETEMPTYTEXT, (WPARAM)(cchText), (LPARAM)(pszText))

	LVM_GETFOOTERRECT = (LVM_FIRST + 205)
	// #define ListView_GetFooterRect(hwnd, prc) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETFOOTERRECT, (WPARAM)(0), (LPARAM)(prc))

	// // footer flags
	LVFF_ITEMCOUNT = 0

	// typedef struct tagLVFOOTERINFO
	// {
	//     UINT mask;          // LVFF_*
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     UINT cItems;
	// } LVFOOTERINFO, *LPLVFOOTERINFO;

	LVM_GETFOOTERINFO = (LVM_FIRST + 206)
	// #define ListView_GetFooterInfo(hwnd, plvfi) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETFOOTERINFO, (WPARAM)(0), (LPARAM)(plvfi))

	LVM_GETFOOTERITEMRECT = (LVM_FIRST + 207)
	// #define ListView_GetFooterItemRect(hwnd, iItem, prc) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETFOOTERITEMRECT, (WPARAM)(iItem), (LPARAM)(prc))

	// // footer item flags
	LVFIF_TEXT  = 0
	LVFIF_STATE = 0

	// // footer item state
	LVFIS_FOCUSED = 0

	// typedef struct tagLVFOOTERITEM
	// {
	//     UINT mask;          // LVFIF_*
	//     int iItem;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     UINT state;         // LVFIS_*
	//     UINT stateMask;     // LVFIS_*
	// } LVFOOTERITEM, *LPLVFOOTERITEM;

	LVM_GETFOOTERITEM = (LVM_FIRST + 208)
	// #define ListView_GetFooterItem(hwnd, iItem, pfi) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETFOOTERITEM, (WPARAM)(iItem), (LPARAM)(pfi))

	// // supports a single item in multiple groups.
	// typedef struct tagLVITEMINDEX
	// {
	//     int iItem;          // listview item index
	//     int iGroup;         // group index (must be -1 if group view is not enabled)
	// } LVITEMINDEX, *PLVITEMINDEX;

	LVM_GETITEMINDEXRECT = (LVM_FIRST + 209)
	// #define ListView_GetItemIndexRect(hwnd, plvii, iSubItem, code, prc) \
	//         (BOOL)SNDMSG((hwnd), LVM_GETITEMINDEXRECT, (WPARAM)(LVITEMINDEX*)(plvii), \
	//((prc) ? ((((LPRECT)(prc))->top = (iSubItem)), (((LPRECT)(prc))->left = (code)), (LPARAM)(prc)) : (LPARAM)(LPRECT)NULL))

	LVM_SETITEMINDEXSTATE = (LVM_FIRST + 210)
	// #define ListView_SetItemIndexState(hwndLV, plvii, data, mask) \
	// { LV_ITEM _macro_lvi;\
	//_macro_lvi.stateMask = (mask);\
	//_macro_lvi.state = (data);\
	//   SNDMSG((hwndLV), LVM_SETITEMINDEXSTATE, (WPARAM)(LVITEMINDEX*)(plvii), (LPARAM)(LV_ITEM *)&_macro_lvi);\
	// }

	LVM_GETNEXTITEMINDEX = (LVM_FIRST + 211)
	// #define ListView_GetNextItemIndex(hwnd, plvii, flags) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETNEXTITEMINDEX, (WPARAM)(LVITEMINDEX*)(plvii), MAKELPARAM((flags), 0))

	// #endif

	// #endif

	//// #ifdef UNICODE
	//LVBKIMAGE = LVBKIMAGEW
	//LPLVBKIMAGE = LPLVBKIMAGEW
	//LVM_SETBKIMAGE = LVM_SETBKIMAGEW
	//LVM_GETBKIMAGE = LVM_GETBKIMAGEW
	//// #else
	//LVBKIMAGE = LVBKIMAGEA
	//LPLVBKIMAGE = LPLVBKIMAGEA
	//LVM_SETBKIMAGE = LVM_SETBKIMAGEA
	//LVM_GETBKIMAGE = LVM_GETBKIMAGEA
	//// #endif

	// #define ListView_SetBkImage(hwnd, plvbki) \
	//     (BOOL)SNDMSG((hwnd), LVM_SETBKIMAGE, 0, (LPARAM)(plvbki))

	// #define ListView_GetBkImage(hwnd, plvbki) \
	//     (BOOL)SNDMSG((hwnd), LVM_GETBKIMAGE, 0, (LPARAM)(plvbki))

	//LPNM_LISTVIEW = LPNMLISTVIEW
	//NM_LISTVIEW = NMLISTVIEW

	// typedef struct tagNMLISTVIEW
	// {
	//     NMHDR   hdr;
	//     int     iItem;
	//     int     iSubItem;
	//     UINT    uNewState;
	//     UINT    uOldState;
	//     UINT    uChanged;
	//     POINT   ptAction;
	//     LPARAM  lParam;
	// } NMLISTVIEW, *LPNMLISTVIEW;

	// NMITEMACTIVATE is used instead of NMLISTVIEW in IE >= 0x400
	// // therefore all the fields are the same except for extra uKeyFlags
	// // they are used to store key flags at the time of the single click with
	// // delayed activation - because by the time the timer goes off a user may
	// // not hold the keys (shift, ctrl) any more
	// typedef struct tagNMITEMACTIVATE
	// {
	//     NMHDR   hdr;
	//     int     iItem;
	//     int     iSubItem;
	//     UINT    uNewState;
	//     UINT    uOldState;
	//     UINT    uChanged;
	//     POINT   ptAction;
	//     LPARAM  lParam;
	//     UINT    uKeyFlags;
	// } NMITEMACTIVATE, *LPNMITEMACTIVATE;

	// // key flags stored in uKeyFlags
	LVKF_ALT     = 0
	LVKF_CONTROL = 0
	LVKF_SHIFT   = 0

	//NMLVCUSTOMDRAW_V3_SIZE = CCSIZEOF_STRUCT(NMLVCUSTOMDRAW, clrTextBk)

	// typedef struct tagNMLVCUSTOMDRAW
	// {
	//     NMCUSTOMDRAW nmcd;
	//     COLORREF clrText;
	//     COLORREF clrTextBk;
	//     int iSubItem;
	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	//     DWORD dwItemType;

	//     // Item custom draw
	//     COLORREF clrFace;
	//     int iIconEffect;
	//     int iIconPhase;
	//     int iPartId;
	//     int iStateId;

	//     // Group Custom Draw
	//     RECT rcText;
	//     UINT uAlign;      // Alignment. Use LVGA_HEADER_CENTER, LVGA_HEADER_RIGHT, LVGA_HEADER_LEFT
	// #endif
	// } NMLVCUSTOMDRAW, *LPNMLVCUSTOMDRAW;

	// // dwItemType
	LVCDI_ITEM      = 0
	LVCDI_GROUP     = 0
	LVCDI_ITEMSLIST = 0

	// // ListView custom draw return values
	LVCDRF_NOSELECT     = 0
	LVCDRF_NOGROUPFRAME = 0

	// typedef struct tagNMLVCACHEHINT
	// {
	//     NMHDR   hdr;
	//     int     iFrom;
	//     int     iTo;
	// } NMLVCACHEHINT, *LPNMLVCACHEHINT;

	//LPNM_CACHEHINT = LPNMLVCACHEHINT
	//PNM_CACHEHINT = LPNMLVCACHEHINT
	//NM_CACHEHINT = NMLVCACHEHINT

	// typedef struct tagNMLVFINDITEMA
	// {
	//     NMHDR   hdr;
	//     int     iStart;
	//     LVFINDINFOA lvfi;
	// } NMLVFINDITEMA, *LPNMLVFINDITEMA;

	// typedef struct tagNMLVFINDITEMW
	// {
	//     NMHDR   hdr;
	//     int     iStart;
	//     LVFINDINFOW lvfi;
	// } NMLVFINDITEMW, *LPNMLVFINDITEMW;

	//PNM_FINDITEMA = LPNMLVFINDITEMA
	//LPNM_FINDITEMA = LPNMLVFINDITEMA
	//NM_FINDITEMA = NMLVFINDITEMA
	//
	//PNM_FINDITEMW = LPNMLVFINDITEMW
	//LPNM_FINDITEMW = LPNMLVFINDITEMW
	//NM_FINDITEMW = NMLVFINDITEMW
	//
	//// #ifdef UNICODE
	//PNM_FINDITEM = PNM_FINDITEMW
	//LPNM_FINDITEM = LPNM_FINDITEMW
	//NM_FINDITEM = NM_FINDITEMW
	//NMLVFINDITEM = NMLVFINDITEMW
	//LPNMLVFINDITEM = LPNMLVFINDITEMW
	//// #else
	//PNM_FINDITEM = PNM_FINDITEMA
	//LPNM_FINDITEM = LPNM_FINDITEMA
	//NM_FINDITEM = NM_FINDITEMA
	//NMLVFINDITEM = NMLVFINDITEMA
	//LPNMLVFINDITEM = LPNMLVFINDITEMA
	//// #endif

	// typedef struct tagNMLVODSTATECHANGE
	// {
	//     NMHDR hdr;
	//     int iFrom;
	//     int iTo;
	//     UINT uNewState;
	//     UINT uOldState;
	// } NMLVODSTATECHANGE, *LPNMLVODSTATECHANGE;

	//PNM_ODSTATECHANGE = LPNMLVODSTATECHANGE
	//LPNM_ODSTATECHANGE = LPNMLVODSTATECHANGE
	//NM_ODSTATECHANGE = NMLVODSTATECHANGE

	LVN_ITEMCHANGING    = (LVN_FIRST - 0)
	LVN_ITEMCHANGED     = (LVN_FIRST - 1)
	LVN_INSERTITEM      = (LVN_FIRST - 2)
	LVN_DELETEITEM      = (LVN_FIRST - 3)
	LVN_DELETEALLITEMS  = (LVN_FIRST - 4)
	LVN_BEGINLABELEDITA = (LVN_FIRST - 5)
	LVN_BEGINLABELEDITW = (LVN_FIRST - 75)
	LVN_ENDLABELEDITA   = (LVN_FIRST - 6)
	LVN_ENDLABELEDITW   = (LVN_FIRST - 76)
	LVN_COLUMNCLICK     = (LVN_FIRST - 8)
	LVN_BEGINDRAG       = (LVN_FIRST - 9)
	LVN_BEGINRDRAG      = (LVN_FIRST - 11)

	LVN_ODCACHEHINT = (LVN_FIRST - 13)
	LVN_ODFINDITEMA = (LVN_FIRST - 52)
	LVN_ODFINDITEMW = (LVN_FIRST - 79)

	LVN_ITEMACTIVATE   = (LVN_FIRST - 14)
	LVN_ODSTATECHANGED = (LVN_FIRST - 15)

	//// #ifdef UNICODE
	//LVN_ODFINDITEM = LVN_ODFINDITEMW
	//// #else
	//LVN_ODFINDITEM = LVN_ODFINDITEMA
	//// #endif

	LVN_HOTTRACK = (LVN_FIRST - 21)

	LVN_GETDISPINFOA = (LVN_FIRST - 50)
	LVN_GETDISPINFOW = (LVN_FIRST - 77)
	LVN_SETDISPINFOA = (LVN_FIRST - 51)
	LVN_SETDISPINFOW = (LVN_FIRST - 78)

	//// #ifdef UNICODE
	//LVN_BEGINLABELEDIT = LVN_BEGINLABELEDITW
	//LVN_ENDLABELEDIT = LVN_ENDLABELEDITW
	//LVN_GETDISPINFO = LVN_GETDISPINFOW
	//LVN_SETDISPINFO = LVN_SETDISPINFOW
	//// #else
	//LVN_BEGINLABELEDIT = LVN_BEGINLABELEDITA
	//LVN_ENDLABELEDIT = LVN_ENDLABELEDITA
	//LVN_GETDISPINFO = LVN_GETDISPINFOA
	//LVN_SETDISPINFO = LVN_SETDISPINFOA
	//// #endif

	LVIF_DI_SETITEM = 0

	//LV_DISPINFOA = NMLVDISPINFOA
	//LV_DISPINFOW = NMLVDISPINFOW
	//
	//LV_DISPINFO = NMLVDISPINFO

	// typedef struct tagLVDISPINFO {
	//     NMHDR hdr;
	//     LVITEMA item;
	// } NMLVDISPINFOA, *LPNMLVDISPINFOA;

	// typedef struct tagLVDISPINFOW {
	//     NMHDR hdr;
	//     LVITEMW item;
	// } NMLVDISPINFOW, *LPNMLVDISPINFOW;

	//// #ifdef UNICODE
	//NMLVDISPINFO = NMLVDISPINFOW
	//// #else
	//NMLVDISPINFO = NMLVDISPINFOA
	//// #endif

	LVN_KEYDOWN = (LVN_FIRST - 55)

	//LV_KEYDOWN = NMLVKEYDOWN

	// #ifdef _WIN32
	// #include <pshpack1.h>
	// #endif

	// typedef struct tagLVKEYDOWN
	// {
	//     NMHDR hdr;
	//     WORD wVKey;
	//     UINT flags;
	// } NMLVKEYDOWN, *LPNMLVKEYDOWN;

	// #ifdef _WIN32
	// #include <poppack.h>
	// #endif

	LVN_MARQUEEBEGIN = (LVN_FIRST - 56)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	// typedef struct tagNMLVLINK
	// {
	//     NMHDR       hdr;
	//     LITEM       link;
	//     int         iItem;
	//     int         iSubItem;
	// } NMLVLINK,  *PNMLVLINK;
	// #endif

	// typedef struct tagNMLVGETINFOTIPA
	// {
	//     NMHDR hdr;
	//     DWORD dwFlags;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iItem;
	//     int iSubItem;
	//     LPARAM lParam;
	// } NMLVGETINFOTIPA, *LPNMLVGETINFOTIPA;

	// typedef struct tagNMLVGETINFOTIPW
	// {
	//     NMHDR hdr;
	//     DWORD dwFlags;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iItem;
	//     int iSubItem;
	//     LPARAM lParam;
	// } NMLVGETINFOTIPW, *LPNMLVGETINFOTIPW;

	// // NMLVGETINFOTIPA.dwFlag values

	LVGIT_UNFOLDED = 0

	LVN_GETINFOTIPA = (LVN_FIRST - 57)
	LVN_GETINFOTIPW = (LVN_FIRST - 58)

	//// #ifdef UNICODE
	//LVN_GETINFOTIP = LVN_GETINFOTIPW
	//NMLVGETINFOTIP = NMLVGETINFOTIPW
	//LPNMLVGETINFOTIP = LPNMLVGETINFOTIPW
	//// #else
	//LVN_GETINFOTIP = LVN_GETINFOTIPA
	//NMLVGETINFOTIP = NMLVGETINFOTIPA
	//LPNMLVGETINFOTIP = LPNMLVGETINFOTIPA
	//// #endif

	// //
	// //  LVN_INCREMENTALSEARCH gives the app the opportunity to customize
	// //  incremental search.  For example, if the items are numeric,
	// //  the app can do numerical search instead of string search.
	// //
	// //  ListView notifies the app with NMLVFINDITEM.
	// //  The app sets pnmfi->lvfi.lParam to the result of the incremental search,
	// //  or to LVNSCH_DEFAULT if ListView should do the default search,
	// //  or to LVNSCH_ERROR to fail the search and just beep,
	// //  or to LVNSCH_IGNORE to stop all ListView processing.
	// //
	// //  The return value is not used.

	// #define LVNSCH_DEFAULT  -1
	// #define LVNSCH_ERROR    -2
	// #define LVNSCH_IGNORE   -3

	LVN_INCREMENTALSEARCHA = (LVN_FIRST - 62)
	LVN_INCREMENTALSEARCHW = (LVN_FIRST - 63)

	//// #ifdef UNICODE
	//LVN_INCREMENTALSEARCH = LVN_INCREMENTALSEARCHW
	//// #else
	//LVN_INCREMENTALSEARCH = LVN_INCREMENTALSEARCHA
	//// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVN_COLUMNDROPDOWN = (LVN_FIRST - 64)

	LVN_COLUMNOVERFLOWCLICK = (LVN_FIRST - 66)

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	// typedef struct tagNMLVSCROLL
	// {
	//     NMHDR   hdr;
	//     int     dx;
	//     int     dy;
	// } NMLVSCROLL, *LPNMLVSCROLL;

	LVN_BEGINSCROLL = (LVN_FIRST - 80)
	LVN_ENDSCROLL   = (LVN_FIRST - 81)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	LVN_LINKCLICK = (LVN_FIRST - 84)

	EMF_CENTERED = 0 // render markup centered in the listview area

	// typedef struct tagNMLVEMPTYMARKUP
	// {
	//     NMHDR hdr;
	//     // out params from client back to listview
	//     DWORD dwFlags;                      // EMF_*
	//     WCHAR szMarkup[L_MAX_URL_LENGTH];   // markup displayed
	// } NMLVEMPTYMARKUP;

	LVN_GETEMPTYMARKUP = (LVN_FIRST - 87)

	// #endif

	// #endif // NOLISTVIEW

	//====== TREEVIEW CONTROL =====================================================

	// #ifndef NOTREEVIEW

	// #ifdef _WIN32
	// #define WC_TREEVIEWA            "SysTreeView32"
	WC_TREEVIEWW = "SysTreeView32"

	// #ifdef UNICODE
	WC_TREEVIEW = WC_TREEVIEWW
	// #else
	//WC_TREEVIEW = WC_TREEVIEWA
	// #endif

	// #else
	// #define WC_TREEVIEW             "SysTreeView"
	// #endif

	// // begin_r_commctrl

	TVS_HASBUTTONS      = 0
	TVS_HASLINES        = 0
	TVS_LINESATROOT     = 0
	TVS_EDITLABELS      = 0
	TVS_DISABLEDRAGDROP = 0
	TVS_SHOWSELALWAYS   = 0
	TVS_RTLREADING      = 0

	TVS_NOTOOLTIPS    = 0
	TVS_CHECKBOXES    = 0
	TVS_TRACKSELECT   = 0
	TVS_SINGLEEXPAND  = 0
	TVS_INFOTIP       = 0
	TVS_FULLROWSELECT = 0
	TVS_NOSCROLL      = 0
	TVS_NONEVENHEIGHT = 0
	TVS_NOHSCROLL     = 0 // TVS_NOSCROLL overrides this

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TVS_EX_NOSINGLECOLLAPSE = 0
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TVS_EX_MULTISELECT         = 0
	TVS_EX_DOUBLEBUFFER        = 0
	TVS_EX_NOINDENTSTATE       = 0
	TVS_EX_RICHTOOLTIP         = 0
	TVS_EX_AUTOHSCROLL         = 0
	TVS_EX_FADEINOUTEXPANDOS   = 0
	TVS_EX_PARTIALCHECKBOXES   = 0
	TVS_EX_EXCLUSIONCHECKBOXES = 0
	TVS_EX_DIMMEDCHECKBOXES    = 0
	TVS_EX_DRAWIMAGEASYNC      = 0
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINTHRESHOLD)
	// #endif

	// // end_r_commctrl

	// struct _TREEITEM;
	// typedef struct _TREEITEM *HTREEITEM;

	TVIF_TEXT          = 0
	TVIF_IMAGE         = 0
	TVIF_PARAM         = 0
	TVIF_STATE         = 0
	TVIF_HANDLE        = 0
	TVIF_SELECTEDIMAGE = 0
	TVIF_CHILDREN      = 0
	TVIF_INTEGRAL      = 0
	// #if (_WIN32_IE >= 0x0600)
	TVIF_STATEEX       = 0
	TVIF_EXPANDEDIMAGE = 0
	// #endif
	TVIS_SELECTED      = 0
	TVIS_CUT           = 0
	TVIS_DROPHILITED   = 0
	TVIS_BOLD          = 0
	TVIS_EXPANDED      = 0
	TVIS_EXPANDEDONCE  = 0
	TVIS_EXPANDPARTIAL = 0

	TVIS_OVERLAYMASK    = 0
	TVIS_STATEIMAGEMASK = 0
	TVIS_USERMASK       = 0

	// #if (_WIN32_IE >= 0x0600)
	TVIS_EX_FLAT = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TVIS_EX_DISABLED = 0
	// #endif
	TVIS_EX_ALL = 0

	// // Structure for TreeView's NM_TVSTATEIMAGECHANGING notification
	// typedef struct tagNMTVSTATEIMAGECHANGING
	// {
	//     NMHDR hdr;
	//     HTREEITEM hti;
	//     int iOldStateImageIndex;
	//     int iNewStateImageIndex;
	// } NMTVSTATEIMAGECHANGING, *LPNMTVSTATEIMAGECHANGING;
	// #endif

	I_CHILDRENCALLBACK = (-1)
	I_CHILDRENAUTO     = (-2)
	//
	//LPTV_ITEMW = LPTVITEMW
	//LPTV_ITEMA = LPTVITEMA
	//TV_ITEMW = TVITEMW
	//TV_ITEMA = TVITEMA
	//
	//LPTV_ITEM = LPTVITEM
	//TV_ITEM = TVITEM

	// typedef struct tagTVITEMA {
	//     UINT      mask;
	//     HTREEITEM hItem;
	//     UINT      state;
	//     UINT      stateMask;
	//     LPSTR     pszText;
	//     int       cchTextMax;
	//     int       iImage;
	//     int       iSelectedImage;
	//     int       cChildren;
	//     LPARAM    lParam;
	// } TVITEMA, *LPTVITEMA;

	// typedef struct tagTVITEMW {
	//     UINT      mask;
	//     HTREEITEM hItem;
	//     UINT      state;
	//     UINT      stateMask;
	//     LPWSTR    pszText;
	//     int       cchTextMax;
	//     int       iImage;
	//     int       iSelectedImage;
	//     int       cChildren;
	//     LPARAM    lParam;
	// } TVITEMW, *LPTVITEMW;

	// // only used for Get and Set messages.  no notifies
	// typedef struct tagTVITEMEXA {
	//     UINT      mask;
	//     HTREEITEM hItem;
	//     UINT      state;
	//     UINT      stateMask;
	//     LPSTR     pszText;
	//     int       cchTextMax;
	//     int       iImage;
	//     int       iSelectedImage;
	//     int       cChildren;
	//     LPARAM    lParam;
	//     int       iIntegral;
	// #if (_WIN32_IE >= 0x0600)
	//     UINT      uStateEx;
	//     HWND      hwnd;
	//     int       iExpandedImage;
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_WIN7)
	//     int       iReserved;
	// #endif
	// } TVITEMEXA, *LPTVITEMEXA;
	// // only used for Get and Set messages.  no notifies
	// typedef struct tagTVITEMEXW {
	//     UINT      mask;
	//     HTREEITEM hItem;
	//     UINT      state;
	//     UINT      stateMask;
	//     LPWSTR    pszText;
	//     int       cchTextMax;
	//     int       iImage;
	//     int       iSelectedImage;
	//     int       cChildren;
	//     LPARAM    lParam;
	//     int       iIntegral;
	// #if (_WIN32_IE >= 0x0600)
	//     UINT      uStateEx;
	//     HWND      hwnd;
	//     int       iExpandedImage;
	// #endif
	// #if (NTDDI_VERSION >= NTDDI_WIN7)
	//     int       iReserved;
	// #endif
	// } TVITEMEXW, *LPTVITEMEXW;
	// #ifdef UNICODE
	// typedef TVITEMEXW TVITEMEX;
	// typedef LPTVITEMEXW LPTVITEMEX;
	// #else
	// typedef TVITEMEXA TVITEMEX;
	// typedef LPTVITEMEXA LPTVITEMEX;
	// #endif // UNICODE

	//// #ifdef UNICODE
	//TVITEM = TVITEMW
	//LPTVITEM = LPTVITEMW
	//// #else
	//TVITEM = TVITEMA
	//LPTVITEM = LPTVITEMA
	//// #endif

	//TVI_ROOT = ((HTREEITEM)(ULONG_PTR)-0x10000)
	//TVI_FIRST = ((HTREEITEM)(ULONG_PTR)-0x0FFFF)
	//TVI_LAST = ((HTREEITEM)(ULONG_PTR)-0x0FFFE)
	//TVI_SORT = ((HTREEITEM)(ULONG_PTR)-0x0FFFD)
	//
	//LPTV_INSERTSTRUCTA = LPTVINSERTSTRUCTA
	//LPTV_INSERTSTRUCTW = LPTVINSERTSTRUCTW
	//TV_INSERTSTRUCTA = TVINSERTSTRUCTA
	//TV_INSERTSTRUCTW = TVINSERTSTRUCTW
	//
	//TV_INSERTSTRUCT = TVINSERTSTRUCT
	//LPTV_INSERTSTRUCT = LPTVINSERTSTRUCT
	//
	//
	//TVINSERTSTRUCTA_V1_SIZE = CCSIZEOF_STRUCT(TVINSERTSTRUCTA, item)
	//TVINSERTSTRUCTW_V1_SIZE = CCSIZEOF_STRUCT(TVINSERTSTRUCTW, item)

	// typedef struct tagTVINSERTSTRUCTA {
	//     HTREEITEM hParent;
	//     HTREEITEM hInsertAfter;
	//     union
	//     {
	//         TVITEMEXA itemex;
	//         TV_ITEMA  item;
	//     } DUMMYUNIONNAME;
	// } TVINSERTSTRUCTA, *LPTVINSERTSTRUCTA;

	// typedef struct tagTVINSERTSTRUCTW {
	//     HTREEITEM hParent;
	//     HTREEITEM hInsertAfter;
	//     union
	//     {
	//         TVITEMEXW itemex;
	//         TV_ITEMW  item;
	//     } DUMMYUNIONNAME;
	// } TVINSERTSTRUCTW, *LPTVINSERTSTRUCTW;

	//// #ifdef UNICODE
	//TVINSERTSTRUCT = TVINSERTSTRUCTW
	//LPTVINSERTSTRUCT = LPTVINSERTSTRUCTW
	//TVINSERTSTRUCT_V1_SIZE = TVINSERTSTRUCTW_V1_SIZE
	//// #else
	//TVINSERTSTRUCT = TVINSERTSTRUCTA
	//LPTVINSERTSTRUCT = LPTVINSERTSTRUCTA
	//TVINSERTSTRUCT_V1_SIZE = TVINSERTSTRUCTA_V1_SIZE
	//// #endif

	TVM_INSERTITEMA = (TV_FIRST + 0)
	TVM_INSERTITEMW = (TV_FIRST + 50)
	//// #ifdef UNICODE
	//TVM_INSERTITEM = TVM_INSERTITEMW
	//// #else
	//TVM_INSERTITEM = TVM_INSERTITEMA
	//// #endif

	// #define TreeView_InsertItem(hwnd, lpis) \
	//     (HTREEITEM)SNDMSG((hwnd), TVM_INSERTITEM, 0, (LPARAM)(LPTV_INSERTSTRUCT)(lpis))

	TVM_DELETEITEM = (TV_FIRST + 1)
	// #define TreeView_DeleteItem(hwnd, hitem) \
	//     (BOOL)SNDMSG((hwnd), TVM_DELETEITEM, 0, (LPARAM)(HTREEITEM)(hitem))

	// #define TreeView_DeleteAllItems(hwnd) \
	//     (BOOL)SNDMSG((hwnd), TVM_DELETEITEM, 0, (LPARAM)TVI_ROOT)

	TVM_EXPAND = (TV_FIRST + 2)
	// #define TreeView_Expand(hwnd, hitem, code) \
	//     (BOOL)SNDMSG((hwnd), TVM_EXPAND, (WPARAM)(code), (LPARAM)(HTREEITEM)(hitem))

	TVE_COLLAPSE      = 0
	TVE_EXPAND        = 0
	TVE_TOGGLE        = 0
	TVE_EXPANDPARTIAL = 0
	TVE_COLLAPSERESET = 0

	TVM_GETITEMRECT = (TV_FIRST + 4)
	// #define TreeView_GetItemRect(hwnd, hitem, prc, code) \
	//(*(HTREEITEM *)(prc) = (hitem), (BOOL)SNDMSG((hwnd), TVM_GETITEMRECT, (WPARAM)(code), (LPARAM)(RECT *)(prc)))

	TVM_GETCOUNT = (TV_FIRST + 5)
	// #define TreeView_GetCount(hwnd) \
	//     (UINT)SNDMSG((hwnd), TVM_GETCOUNT, 0, 0)

	TVM_GETINDENT = (TV_FIRST + 6)
	// #define TreeView_GetIndent(hwnd) \
	//     (UINT)SNDMSG((hwnd), TVM_GETINDENT, 0, 0)

	TVM_SETINDENT = (TV_FIRST + 7)
	// #define TreeView_SetIndent(hwnd, indent) \
	//     (BOOL)SNDMSG((hwnd), TVM_SETINDENT, (WPARAM)(indent), 0)

	TVM_GETIMAGELIST = (TV_FIRST + 8)
	// #define TreeView_GetImageList(hwnd, iImage) \
	//     (HIMAGELIST)SNDMSG((hwnd), TVM_GETIMAGELIST, iImage, 0)

	TVSIL_NORMAL = 0
	TVSIL_STATE  = 2

	TVM_SETIMAGELIST = (TV_FIRST + 9)
	// #define TreeView_SetImageList(hwnd, himl, iImage) \
	//     (HIMAGELIST)SNDMSG((hwnd), TVM_SETIMAGELIST, iImage, (LPARAM)(HIMAGELIST)(himl))

	TVM_GETNEXTITEM = (TV_FIRST + 10)
	// #define TreeView_GetNextItem(hwnd, hitem, code) \
	//     (HTREEITEM)SNDMSG((hwnd), TVM_GETNEXTITEM, (WPARAM)(code), (LPARAM)(HTREEITEM)(hitem))

	TVGN_ROOT            = 0
	TVGN_NEXT            = 0
	TVGN_PREVIOUS        = 0
	TVGN_PARENT          = 0
	TVGN_CHILD           = 0
	TVGN_FIRSTVISIBLE    = 0
	TVGN_NEXTVISIBLE     = 0
	TVGN_PREVIOUSVISIBLE = 0
	TVGN_DROPHILITE      = 0
	TVGN_CARET           = 0
	TVGN_LASTVISIBLE     = 0
	// #if (_WIN32_IE >= 0x0600)
	TVGN_NEXTSELECTED = 0
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TVSI_NOSINGLEEXPAND = 0 // Should not conflict with TVGN flags.
	// #endif

	// #define TreeView_GetChild(hwnd, hitem)          TreeView_GetNextItem(hwnd, hitem, TVGN_CHILD)
	// #define TreeView_GetNextSibling(hwnd, hitem)    TreeView_GetNextItem(hwnd, hitem, TVGN_NEXT)
	// #define TreeView_GetPrevSibling(hwnd, hitem)    TreeView_GetNextItem(hwnd, hitem, TVGN_PREVIOUS)
	// #define TreeView_GetParent(hwnd, hitem)         TreeView_GetNextItem(hwnd, hitem, TVGN_PARENT)
	// #define TreeView_GetFirstVisible(hwnd)          TreeView_GetNextItem(hwnd, NULL,  TVGN_FIRSTVISIBLE)
	// #define TreeView_GetNextVisible(hwnd, hitem)    TreeView_GetNextItem(hwnd, hitem, TVGN_NEXTVISIBLE)
	// #define TreeView_GetPrevVisible(hwnd, hitem)    TreeView_GetNextItem(hwnd, hitem, TVGN_PREVIOUSVISIBLE)
	// #define TreeView_GetSelection(hwnd)             TreeView_GetNextItem(hwnd, NULL,  TVGN_CARET)
	// #define TreeView_GetDropHilight(hwnd)           TreeView_GetNextItem(hwnd, NULL,  TVGN_DROPHILITE)
	// #define TreeView_GetRoot(hwnd)                  TreeView_GetNextItem(hwnd, NULL,  TVGN_ROOT)
	// #define TreeView_GetLastVisible(hwnd)           TreeView_GetNextItem(hwnd, NULL,  TVGN_LASTVISIBLE)
	// #if (_WIN32_IE >= 0x0600)
	// #define TreeView_GetNextSelected(hwnd, hitem)   TreeView_GetNextItem(hwnd, hitem,  TVGN_NEXTSELECTED)
	// #endif

	TVM_SELECTITEM = (TV_FIRST + 11)
	// #define TreeView_Select(hwnd, hitem, code) \
	//     (BOOL)SNDMSG((hwnd), TVM_SELECTITEM, (WPARAM)(code), (LPARAM)(HTREEITEM)(hitem))

	// #define TreeView_SelectItem(hwnd, hitem)            TreeView_Select(hwnd, hitem, TVGN_CARET)
	// #define TreeView_SelectDropTarget(hwnd, hitem)      TreeView_Select(hwnd, hitem, TVGN_DROPHILITE)
	// #define TreeView_SelectSetFirstVisible(hwnd, hitem) TreeView_Select(hwnd, hitem, TVGN_FIRSTVISIBLE)

	TVM_GETITEMA = (TV_FIRST + 12)
	TVM_GETITEMW = (TV_FIRST + 62)

	//// #ifdef UNICODE
	//TVM_GETITEM = TVM_GETITEMW
	//// #else
	//TVM_GETITEM = TVM_GETITEMA
	//// #endif

	// #define TreeView_GetItem(hwnd, pitem) \
	//     (BOOL)SNDMSG((hwnd), TVM_GETITEM, 0, (LPARAM)(TV_ITEM *)(pitem))

	TVM_SETITEMA = (TV_FIRST + 13)
	TVM_SETITEMW = (TV_FIRST + 63)

	//// #ifdef UNICODE
	//TVM_SETITEM = TVM_SETITEMW
	//// #else
	//TVM_SETITEM = TVM_SETITEMA
	//// #endif

	// #define TreeView_SetItem(hwnd, pitem) \
	//     (BOOL)SNDMSG((hwnd), TVM_SETITEM, 0, (LPARAM)(const TV_ITEM *)(pitem))

	TVM_EDITLABELA = (TV_FIRST + 14)
	TVM_EDITLABELW = (TV_FIRST + 65)
	//// #ifdef UNICODE
	//TVM_EDITLABEL = TVM_EDITLABELW
	//// #else
	//TVM_EDITLABEL = TVM_EDITLABELA
	//// #endif

	// #define TreeView_EditLabel(hwnd, hitem) \
	//     (HWND)SNDMSG((hwnd), TVM_EDITLABEL, 0, (LPARAM)(HTREEITEM)(hitem))

	TVM_GETEDITCONTROL = (TV_FIRST + 15)
	// #define TreeView_GetEditControl(hwnd) \
	//     (HWND)SNDMSG((hwnd), TVM_GETEDITCONTROL, 0, 0)

	TVM_GETVISIBLECOUNT = (TV_FIRST + 16)
	// #define TreeView_GetVisibleCount(hwnd) \
	//     (UINT)SNDMSG((hwnd), TVM_GETVISIBLECOUNT, 0, 0)

	TVM_HITTEST = (TV_FIRST + 17)
	// #define TreeView_HitTest(hwnd, lpht) \
	//     (HTREEITEM)SNDMSG((hwnd), TVM_HITTEST, 0, (LPARAM)(LPTV_HITTESTINFO)(lpht))

	//LPTV_HITTESTINFO = LPTVHITTESTINFO
	//TV_HITTESTINFO = TVHITTESTINFO

	// typedef struct tagTVHITTESTINFO {
	//     POINT       pt;
	//     UINT        flags;
	//     HTREEITEM   hItem;
	// } TVHITTESTINFO, *LPTVHITTESTINFO;

	TVHT_NOWHERE         = 0
	TVHT_ONITEMICON      = 0
	TVHT_ONITEMLABEL     = 0
	TVHT_ONITEM          = (TVHT_ONITEMICON | TVHT_ONITEMLABEL | TVHT_ONITEMSTATEICON)
	TVHT_ONITEMINDENT    = 0
	TVHT_ONITEMBUTTON    = 0
	TVHT_ONITEMRIGHT     = 0
	TVHT_ONITEMSTATEICON = 0

	TVHT_ABOVE   = 0
	TVHT_BELOW   = 0
	TVHT_TORIGHT = 0
	TVHT_TOLEFT  = 0

	TVM_CREATEDRAGIMAGE = (TV_FIRST + 18)
	// #define TreeView_CreateDragImage(hwnd, hitem) \
	//     (HIMAGELIST)SNDMSG((hwnd), TVM_CREATEDRAGIMAGE, 0, (LPARAM)(HTREEITEM)(hitem))

	TVM_SORTCHILDREN = (TV_FIRST + 19)
	// #define TreeView_SortChildren(hwnd, hitem, recurse) \
	//     (BOOL)SNDMSG((hwnd), TVM_SORTCHILDREN, (WPARAM)(recurse), (LPARAM)(HTREEITEM)(hitem))

	TVM_ENSUREVISIBLE = (TV_FIRST + 20)
	// #define TreeView_EnsureVisible(hwnd, hitem) \
	//     (BOOL)SNDMSG((hwnd), TVM_ENSUREVISIBLE, 0, (LPARAM)(HTREEITEM)(hitem))

	TVM_SORTCHILDRENCB = (TV_FIRST + 21)
	// #define TreeView_SortChildrenCB(hwnd, psort, recurse) \
	//     (BOOL)SNDMSG((hwnd), TVM_SORTCHILDRENCB, (WPARAM)(recurse), \
	//     (LPARAM)(LPTV_SORTCB)(psort))

	TVM_ENDEDITLABELNOW = (TV_FIRST + 22)
	// #define TreeView_EndEditLabelNow(hwnd, fCancel) \
	//     (BOOL)SNDMSG((hwnd), TVM_ENDEDITLABELNOW, (WPARAM)(fCancel), 0)

	TVM_GETISEARCHSTRINGA = (TV_FIRST + 23)
	TVM_GETISEARCHSTRINGW = (TV_FIRST + 64)

	//// #ifdef UNICODE
	//TVM_GETISEARCHSTRING = TVM_GETISEARCHSTRINGW
	//// #else
	//TVM_GETISEARCHSTRING = TVM_GETISEARCHSTRINGA
	//// #endif

	TVM_SETTOOLTIPS = (TV_FIRST + 24)
	// #define TreeView_SetToolTips(hwnd,  hwndTT) \
	//     (HWND)SNDMSG((hwnd), TVM_SETTOOLTIPS, (WPARAM)(hwndTT), 0)
	TVM_GETTOOLTIPS = (TV_FIRST + 25)
	// #define TreeView_GetToolTips(hwnd) \
	//     (HWND)SNDMSG((hwnd), TVM_GETTOOLTIPS, 0, 0)

	// #define TreeView_GetISearchString(hwndTV, lpsz) \
	//         (BOOL)SNDMSG((hwndTV), TVM_GETISEARCHSTRING, 0, (LPARAM)(LPTSTR)(lpsz))

	TVM_SETINSERTMARK = (TV_FIRST + 26)
	// #define TreeView_SetInsertMark(hwnd, hItem, fAfter) \
	//         (BOOL)SNDMSG((hwnd), TVM_SETINSERTMARK, (WPARAM) (fAfter), (LPARAM) (hItem))

	TVM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	// #define TreeView_SetUnicodeFormat(hwnd, fUnicode)  \
	//     (BOOL)SNDMSG((hwnd), TVM_SETUNICODEFORMAT, (WPARAM)(fUnicode), 0)

	TVM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	// #define TreeView_GetUnicodeFormat(hwnd)  \
	//     (BOOL)SNDMSG((hwnd), TVM_GETUNICODEFORMAT, 0, 0)

	TVM_SETITEMHEIGHT = (TV_FIRST + 27)
	// #define TreeView_SetItemHeight(hwnd,  iHeight) \
	//     (int)SNDMSG((hwnd), TVM_SETITEMHEIGHT, (WPARAM)(iHeight), 0)
	TVM_GETITEMHEIGHT = (TV_FIRST + 28)
	// #define TreeView_GetItemHeight(hwnd) \
	//     (int)SNDMSG((hwnd), TVM_GETITEMHEIGHT, 0, 0)

	TVM_SETBKCOLOR = (TV_FIRST + 29)
	// #define TreeView_SetBkColor(hwnd, clr) \
	//     (COLORREF)SNDMSG((hwnd), TVM_SETBKCOLOR, 0, (LPARAM)(clr))

	TVM_SETTEXTCOLOR = (TV_FIRST + 30)
	// #define TreeView_SetTextColor(hwnd, clr) \
	//     (COLORREF)SNDMSG((hwnd), TVM_SETTEXTCOLOR, 0, (LPARAM)(clr))

	TVM_GETBKCOLOR = (TV_FIRST + 31)
	// #define TreeView_GetBkColor(hwnd) \
	//     (COLORREF)SNDMSG((hwnd), TVM_GETBKCOLOR, 0, 0)

	TVM_GETTEXTCOLOR = (TV_FIRST + 32)
	// #define TreeView_GetTextColor(hwnd) \
	//     (COLORREF)SNDMSG((hwnd), TVM_GETTEXTCOLOR, 0, 0)

	TVM_SETSCROLLTIME = (TV_FIRST + 33)
	// #define TreeView_SetScrollTime(hwnd, uTime) \
	//     (UINT)SNDMSG((hwnd), TVM_SETSCROLLTIME, uTime, 0)

	TVM_GETSCROLLTIME = (TV_FIRST + 34)
	// #define TreeView_GetScrollTime(hwnd) \
	//     (UINT)SNDMSG((hwnd), TVM_GETSCROLLTIME, 0, 0)

	TVM_SETINSERTMARKCOLOR = (TV_FIRST + 37)
	// #define TreeView_SetInsertMarkColor(hwnd, clr) \
	//     (COLORREF)SNDMSG((hwnd), TVM_SETINSERTMARKCOLOR, 0, (LPARAM)(clr))
	TVM_GETINSERTMARKCOLOR = (TV_FIRST + 38)
	// #define TreeView_GetInsertMarkColor(hwnd) \
	//     (COLORREF)SNDMSG((hwnd), TVM_GETINSERTMARKCOLOR, 0, 0)

	TVM_SETBORDER = (TV_FIRST + 35)
	// #define TreeView_SetBorder(hwnd,  dwFlags, xBorder, yBorder) \
	//     (int)SNDMSG((hwnd), TVM_SETBORDER, (WPARAM)(dwFlags), MAKELPARAM(xBorder, yBorder))

	TVSBF_XBORDER = 0
	TVSBF_YBORDER = 0

	// // tvm_?etitemstate only uses mask, state and stateMask.
	// // so unicode or ansi is irrelevant.
	// #define TreeView_SetItemState(hwndTV, hti, data, _mask) \
	// { TVITEM _ms_TVi;\
	//_ms_TVi.mask = TVIF_STATE; \
	//_ms_TVi.hItem = (hti); \
	//_ms_TVi.stateMask = (_mask);\
	//_ms_TVi.state = (data);\
	//   SNDMSG((hwndTV), TVM_SETITEM, 0, (LPARAM)(TV_ITEM *)&_ms_TVi);\
	// }

	// #define TreeView_SetCheckState(hwndTV, hti, fCheck) \
	//   TreeView_SetItemState(hwndTV, hti, INDEXTOSTATEIMAGEMASK((fCheck)?2:1), TVIS_STATEIMAGEMASK)

	TVM_GETITEMSTATE = (TV_FIRST + 39)
	// #define TreeView_GetItemState(hwndTV, hti, mask) \
	//    (UINT)SNDMSG((hwndTV), TVM_GETITEMSTATE, (WPARAM)(hti), (LPARAM)(mask))

	// #define TreeView_GetCheckState(hwndTV, hti) \
	//    ((((UINT)(SNDMSG((hwndTV), TVM_GETITEMSTATE, (WPARAM)(hti), TVIS_STATEIMAGEMASK))) >> 12) -1)

	TVM_SETLINECOLOR = (TV_FIRST + 40)
	// #define TreeView_SetLineColor(hwnd, clr) \
	//     (COLORREF)SNDMSG((hwnd), TVM_SETLINECOLOR, 0, (LPARAM)(clr))

	TVM_GETLINECOLOR = (TV_FIRST + 41)
	// #define TreeView_GetLineColor(hwnd) \
	//     (COLORREF)SNDMSG((hwnd), TVM_GETLINECOLOR, 0, 0)

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	TVM_MAPACCIDTOHTREEITEM = (TV_FIRST + 42)
	// #define TreeView_MapAccIDToHTREEITEM(hwnd, id) \
	//     (HTREEITEM)SNDMSG((hwnd), TVM_MAPACCIDTOHTREEITEM, id, 0)

	TVM_MAPHTREEITEMTOACCID = (TV_FIRST + 43)
	// #define TreeView_MapHTREEITEMToAccID(hwnd, htreeitem) \
	//     (UINT)SNDMSG((hwnd), TVM_MAPHTREEITEMTOACCID, (WPARAM)(htreeitem), 0)

	TVM_SETEXTENDEDSTYLE = (TV_FIRST + 44)
	// #define TreeView_SetExtendedStyle(hwnd, dw, mask) \
	//     (DWORD)SNDMSG((hwnd), TVM_SETEXTENDEDSTYLE, mask, dw)

	TVM_GETEXTENDEDSTYLE = (TV_FIRST + 45)
	// #define TreeView_GetExtendedStyle(hwnd) \
	//     (DWORD)SNDMSG((hwnd), TVM_GETEXTENDEDSTYLE, 0, 0)

	TVM_SETAUTOSCROLLINFO = (TV_FIRST + 59)
	// #define TreeView_SetAutoScrollInfo(hwnd, uPixPerSec, uUpdateTime) \
	//     SNDMSG((hwnd), TVM_SETAUTOSCROLLINFO, (WPARAM)(uPixPerSec), (LPARAM)(uUpdateTime))
	// #endif

	TVM_SETHOT = (TV_FIRST + 58)
	// #define TreeView_SetHot(hwnd, hitem) \
	//     SNDMSG((hwnd), TVM_SETHOT, 0, (LPARAM)(hitem))

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	TVM_GETSELECTEDCOUNT = (TV_FIRST + 70)
	// #define TreeView_GetSelectedCount(hwnd) \
	//     (DWORD)SNDMSG((hwnd), TVM_GETSELECTEDCOUNT, 0, 0)

	TVM_SHOWINFOTIP = (TV_FIRST + 71)
	// #define TreeView_ShowInfoTip(hwnd, hitem) \
	//     (DWORD)SNDMSG((hwnd), TVM_SHOWINFOTIP, 0, (LPARAM)(hitem))

	// typedef enum _TVITEMPART
	// {
	TVGIPR_BUTTON = 0x0001
	// } TVITEMPART;

	// typedef struct tagTVGETITEMPARTRECTINFO {
	//     HTREEITEM hti;
	//     RECT*     prc;
	//     TVITEMPART partID;
	// } TVGETITEMPARTRECTINFO;

	TVM_GETITEMPARTRECT = (TV_FIRST + 72)
	// #define TreeView_GetItemPartRect(hwnd, hitem, prc, partid) \
	// { TVGETITEMPARTRECTINFO info; \
	//info.hti = (hitem); \
	//info.prc = (prc); \
	//info.partID = (partid); \
	//   (BOOL)SNDMSG((hwnd), TVM_GETITEMPARTRECT, 0, (LPARAM)&info); \
	// }

	// #endif

	// typedef int (CALLBACK *PFNTVCOMPARE)(LPARAM lParam1, LPARAM lParam2, LPARAM lParamSort);

	//LPTV_SORTCB = LPTVSORTCB
	//TV_SORTCB = TVSORTCB

	// typedef struct tagTVSORTCB
	// {
	//     HTREEITEM       hParent;
	//     PFNTVCOMPARE    lpfnCompare;
	//     LPARAM          lParam;
	// } TVSORTCB, *LPTVSORTCB;

	//LPNM_TREEVIEWA = LPNMTREEVIEWA
	//LPNM_TREEVIEWW = LPNMTREEVIEWW
	//NM_TREEVIEWW = NMTREEVIEWW
	//NM_TREEVIEWA = NMTREEVIEWA
	//
	//LPNM_TREEVIEW = LPNMTREEVIEW
	//NM_TREEVIEW = NMTREEVIEW

	// typedef struct tagNMTREEVIEWA {
	//     NMHDR       hdr;
	//     UINT        action;
	//     TVITEMA    itemOld;
	//     TVITEMA    itemNew;
	//     POINT       ptDrag;
	// } NMTREEVIEWA, *LPNMTREEVIEWA;

	// typedef struct tagNMTREEVIEWW {
	//     NMHDR       hdr;
	//     UINT        action;
	//     TVITEMW    itemOld;
	//     TVITEMW    itemNew;
	//     POINT       ptDrag;
	// } NMTREEVIEWW, *LPNMTREEVIEWW;

	//// #ifdef UNICODE
	//NMTREEVIEW = NMTREEVIEWW
	//LPNMTREEVIEW = LPNMTREEVIEWW
	//// #else
	//NMTREEVIEW = NMTREEVIEWA
	//LPNMTREEVIEW = LPNMTREEVIEWA
	//// #endif

	TVN_SELCHANGINGA = (TVN_FIRST - 1)
	TVN_SELCHANGINGW = (TVN_FIRST - 50)
	TVN_SELCHANGEDA  = (TVN_FIRST - 2)
	TVN_SELCHANGEDW  = (TVN_FIRST - 51)

	TVC_UNKNOWN    = 0
	TVC_BYMOUSE    = 0
	TVC_BYKEYBOARD = 0

	TVN_GETDISPINFOA = (TVN_FIRST - 3)
	TVN_GETDISPINFOW = (TVN_FIRST - 52)
	TVN_SETDISPINFOA = (TVN_FIRST - 4)
	TVN_SETDISPINFOW = (TVN_FIRST - 53)

	TVIF_DI_SETITEM = 0

	//TV_DISPINFOA = NMTVDISPINFOA
	//TV_DISPINFOW = NMTVDISPINFOW
	//
	//TV_DISPINFO = NMTVDISPINFO

	// typedef struct tagTVDISPINFOA {
	//     NMHDR hdr;
	//     TVITEMA item;
	// } NMTVDISPINFOA, *LPNMTVDISPINFOA;

	// typedef struct tagTVDISPINFOW {
	//     NMHDR hdr;
	//     TVITEMW item;
	// } NMTVDISPINFOW, *LPNMTVDISPINFOW;

	//// #ifdef UNICODE
	//NMTVDISPINFO = NMTVDISPINFOW
	//LPNMTVDISPINFO = LPNMTVDISPINFOW
	//// #else
	//NMTVDISPINFO = NMTVDISPINFOA
	//LPNMTVDISPINFO = LPNMTVDISPINFOA
	//// #endif

	// #if (_WIN32_IE >= 0x0600)

	// typedef struct tagTVDISPINFOEXA {
	//     NMHDR hdr;
	//     TVITEMEXA item;
	// } NMTVDISPINFOEXA, *LPNMTVDISPINFOEXA;

	// typedef struct tagTVDISPINFOEXW {
	//     NMHDR hdr;
	//     TVITEMEXW item;
	// } NMTVDISPINFOEXW, *LPNMTVDISPINFOEXW;

	//// #ifdef UNICODE
	//NMTVDISPINFOEX = NMTVDISPINFOEXW
	//LPNMTVDISPINFOEX = LPNMTVDISPINFOEXW
	//// #else
	//NMTVDISPINFOEX = NMTVDISPINFOEXA
	//LPNMTVDISPINFOEX = LPNMTVDISPINFOEXA
	//// #endif
	//
	//TV_DISPINFOEXA = NMTVDISPINFOEXA
	//TV_DISPINFOEXW = NMTVDISPINFOEXW
	//TV_DISPINFOEX = NMTVDISPINFOEX

	// #endif
	TVN_ITEMEXPANDINGA  = (TVN_FIRST - 5)
	TVN_ITEMEXPANDINGW  = (TVN_FIRST - 54)
	TVN_ITEMEXPANDEDA   = (TVN_FIRST - 6)
	TVN_ITEMEXPANDEDW   = (TVN_FIRST - 55)
	TVN_BEGINDRAGA      = (TVN_FIRST - 7)
	TVN_BEGINDRAGW      = (TVN_FIRST - 56)
	TVN_BEGINRDRAGA     = (TVN_FIRST - 8)
	TVN_BEGINRDRAGW     = (TVN_FIRST - 57)
	TVN_DELETEITEMA     = (TVN_FIRST - 9)
	TVN_DELETEITEMW     = (TVN_FIRST - 58)
	TVN_BEGINLABELEDITA = (TVN_FIRST - 10)
	TVN_BEGINLABELEDITW = (TVN_FIRST - 59)
	TVN_ENDLABELEDITA   = (TVN_FIRST - 11)
	TVN_ENDLABELEDITW   = (TVN_FIRST - 60)
	TVN_KEYDOWN         = (TVN_FIRST - 12)

	TVN_GETINFOTIPA  = (TVN_FIRST - 13)
	TVN_GETINFOTIPW  = (TVN_FIRST - 14)
	TVN_SINGLEEXPAND = (TVN_FIRST - 15)

	TVNRET_DEFAULT = 0
	TVNRET_SKIPOLD = 1
	TVNRET_SKIPNEW = 2

	// #if (_WIN32_IE >= 0x0600)
	TVN_ITEMCHANGINGA = (TVN_FIRST - 16)
	TVN_ITEMCHANGINGW = (TVN_FIRST - 17)
	TVN_ITEMCHANGEDA  = (TVN_FIRST - 18)
	TVN_ITEMCHANGEDW  = (TVN_FIRST - 19)
	TVN_ASYNCDRAW     = (TVN_FIRST - 20)
	// #endif

	//TV_KEYDOWN = NMTVKEYDOWN

	// #ifdef _WIN32
	// #include <pshpack1.h>
	// #endif

	// typedef struct tagTVKEYDOWN {
	//     NMHDR hdr;
	//     WORD wVKey;
	//     UINT flags;
	// } NMTVKEYDOWN, *LPNMTVKEYDOWN;

	// #ifdef _WIN32
	// #include <poppack.h>
	// #endif

	//// #ifdef UNICODE
	//TVN_SELCHANGING = TVN_SELCHANGINGW
	//TVN_SELCHANGED = TVN_SELCHANGEDW
	//TVN_GETDISPINFO = TVN_GETDISPINFOW
	//TVN_SETDISPINFO = TVN_SETDISPINFOW
	//TVN_ITEMEXPANDING = TVN_ITEMEXPANDINGW
	//TVN_ITEMEXPANDED = TVN_ITEMEXPANDEDW
	//TVN_BEGINDRAG = TVN_BEGINDRAGW
	//TVN_BEGINRDRAG = TVN_BEGINRDRAGW
	//TVN_DELETEITEM = TVN_DELETEITEMW
	//TVN_BEGINLABELEDIT = TVN_BEGINLABELEDITW
	//TVN_ENDLABELEDIT = TVN_ENDLABELEDITW
	//// #else
	//TVN_SELCHANGING = TVN_SELCHANGINGA
	//TVN_SELCHANGED = TVN_SELCHANGEDA
	//TVN_GETDISPINFO = TVN_GETDISPINFOA
	//TVN_SETDISPINFO = TVN_SETDISPINFOA
	//TVN_ITEMEXPANDING = TVN_ITEMEXPANDINGA
	//TVN_ITEMEXPANDED = TVN_ITEMEXPANDEDA
	//TVN_BEGINDRAG = TVN_BEGINDRAGA
	//TVN_BEGINRDRAG = TVN_BEGINRDRAGA
	//TVN_DELETEITEM = TVN_DELETEITEMA
	//TVN_BEGINLABELEDIT = TVN_BEGINLABELEDITA
	//TVN_ENDLABELEDIT = TVN_ENDLABELEDITA
	//// #endif
	//
	//NMTVCUSTOMDRAW_V3_SIZE = CCSIZEOF_STRUCT(NMTVCUSTOMDRAW, clrTextBk)

	// typedef struct tagNMTVCUSTOMDRAW
	// {
	//     NMCUSTOMDRAW nmcd;
	//     COLORREF     clrText;
	//     COLORREF     clrTextBk;
	//     int iLevel;
	// } NMTVCUSTOMDRAW, *LPNMTVCUSTOMDRAW;

	// // for tooltips

	// typedef struct tagNMTVGETINFOTIPA
	// {
	//     NMHDR hdr;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     HTREEITEM hItem;
	//     LPARAM lParam;
	// } NMTVGETINFOTIPA, *LPNMTVGETINFOTIPA;

	// typedef struct tagNMTVGETINFOTIPW
	// {
	//     NMHDR hdr;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     HTREEITEM hItem;
	//     LPARAM lParam;
	// } NMTVGETINFOTIPW, *LPNMTVGETINFOTIPW;

	//// #ifdef UNICODE
	//TVN_GETINFOTIP = TVN_GETINFOTIPW
	//NMTVGETINFOTIP = NMTVGETINFOTIPW
	//LPNMTVGETINFOTIP = LPNMTVGETINFOTIPW
	//// #else
	//TVN_GETINFOTIP = TVN_GETINFOTIPA
	//NMTVGETINFOTIP = NMTVGETINFOTIPA
	//LPNMTVGETINFOTIP = LPNMTVGETINFOTIPA
	//// #endif

	// // treeview's customdraw return meaning don't draw images.  valid on CDRF_NOTIFYITEMPREPAINT
	TVCDRF_NOIMAGES = 0

	// #if (_WIN32_IE > 0x0600)
	// typedef struct tagTVITEMCHANGE {
	//     NMHDR hdr;
	//     UINT uChanged;
	//     HTREEITEM hItem;
	//     UINT uStateNew;
	//     UINT uStateOld;
	//     LPARAM lParam;
	// } NMTVITEMCHANGE;

	// typedef struct tagNMTVASYNCDRAW
	// {
	//     NMHDR     hdr;
	//     IMAGELISTDRAWPARAMS *pimldp;    // the draw that failed
	//     HRESULT   hr;                   // why it failed
	//     HTREEITEM hItem;                // item that failed to draw icon
	//     LPARAM    lParam;               // its data
	//     // Out Params
	//     DWORD     dwRetFlags;           // What listview should do on return
	//     int       iRetImageIndex;       // used if ADRF_DRAWIMAGE is returned
	// } NMTVASYNCDRAW;

	//// #ifdef UNICODE
	//TVN_ITEMCHANGING = TVN_ITEMCHANGINGW
	//TVN_ITEMCHANGED = TVN_ITEMCHANGEDW
	//// #else
	//TVN_ITEMCHANGING = TVN_ITEMCHANGINGA
	//TVN_ITEMCHANGED = TVN_ITEMCHANGEDA
	//// #endif

	// #endif      // _WIN32_IE >= 0x0600

	// #endif      // NOTREEVIEW

	// #ifndef NOUSEREXCONTROLS

	// ////////////////////  ComboBoxEx ////////////////////////////////

	WC_COMBOBOXEXW = "ComboBoxEx32"
	// #define WC_COMBOBOXEXA         "ComboBoxEx32"

	// #ifdef UNICODE
	WC_COMBOBOXEX = WC_COMBOBOXEXW
	// #else
	//WC_COMBOBOXEX = WC_COMBOBOXEXA
	// #endif

	CBEIF_TEXT          = 0
	CBEIF_IMAGE         = 0
	CBEIF_SELECTEDIMAGE = 0
	CBEIF_OVERLAY       = 0
	CBEIF_INDENT        = 0
	CBEIF_LPARAM        = 0

	CBEIF_DI_SETITEM = 0

	// typedef struct tagCOMBOBOXEXITEMA
	// {
	//     UINT mask;
	//     INT_PTR iItem;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iImage;
	//     int iSelectedImage;
	//     int iOverlay;
	//     int iIndent;
	//     LPARAM lParam;
	// } COMBOBOXEXITEMA, *PCOMBOBOXEXITEMA;
	// typedef COMBOBOXEXITEMA CONST *PCCOMBOEXITEMA;

	// typedef struct tagCOMBOBOXEXITEMW
	// {
	//     UINT mask;
	//     INT_PTR iItem;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iImage;
	//     int iSelectedImage;
	//     int iOverlay;
	//     int iIndent;
	//     LPARAM lParam;
	// } COMBOBOXEXITEMW, *PCOMBOBOXEXITEMW;
	// typedef COMBOBOXEXITEMW CONST *PCCOMBOEXITEMW;

	//// #ifdef UNICODE
	//COMBOBOXEXITEM = COMBOBOXEXITEMW
	//PCOMBOBOXEXITEM = PCOMBOBOXEXITEMW
	//PCCOMBOBOXEXITEM = PCCOMBOBOXEXITEMW
	//// #else
	//COMBOBOXEXITEM = COMBOBOXEXITEMA
	//PCOMBOBOXEXITEM = PCOMBOBOXEXITEMA
	//PCCOMBOBOXEXITEM = PCCOMBOBOXEXITEMA
	//// #endif

	CBEM_INSERTITEMA      = (WM_USER + 1)
	CBEM_SETIMAGELIST     = (WM_USER + 2)
	CBEM_GETIMAGELIST     = (WM_USER + 3)
	CBEM_GETITEMA         = (WM_USER + 4)
	CBEM_SETITEMA         = (WM_USER + 5)
	CBEM_DELETEITEM       = win.CB_DELETESTRING
	CBEM_GETCOMBOCONTROL  = (WM_USER + 6)
	CBEM_GETEDITCONTROL   = (WM_USER + 7)
	CBEM_SETEXSTYLE       = (WM_USER + 8)  // use  SETEXTENDEDSTYLE instead
	CBEM_SETEXTENDEDSTYLE = (WM_USER + 14) // lparam == new style, wParam (optional) == mask
	CBEM_GETEXSTYLE       = (WM_USER + 9)  // use GETEXTENDEDSTYLE instead
	CBEM_GETEXTENDEDSTYLE = (WM_USER + 9)
	CBEM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	CBEM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	CBEM_HASEDITCHANGED   = (WM_USER + 10)
	CBEM_INSERTITEMW      = (WM_USER + 11)
	CBEM_SETITEMW         = (WM_USER + 12)
	CBEM_GETITEMW         = (WM_USER + 13)

	//// #ifdef UNICODE
	//CBEM_INSERTITEM = CBEM_INSERTITEMW
	//CBEM_SETITEM = CBEM_SETITEMW
	//CBEM_GETITEM = CBEM_GETITEMW
	//// #else
	//CBEM_INSERTITEM = CBEM_INSERTITEMA
	//CBEM_SETITEM = CBEM_SETITEMA
	//CBEM_GETITEM = CBEM_GETITEMA
	//// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	CBEM_SETWINDOWTHEME = CCM_SETWINDOWTHEME
	// #endif

	CBES_EX_NOEDITIMAGE       = 0
	CBES_EX_NOEDITIMAGEINDENT = 0
	CBES_EX_PATHWORDBREAKPROC = 0
	CBES_EX_NOSIZELIMIT       = 0
	CBES_EX_CASESENSITIVE     = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	CBES_EX_TEXTENDELLIPSIS = 0
	// #endif

	// typedef struct {
	//     NMHDR hdr;
	//     COMBOBOXEXITEMA ceItem;
	// } NMCOMBOBOXEXA, *PNMCOMBOBOXEXA;

	// typedef struct {
	//     NMHDR hdr;
	//     COMBOBOXEXITEMW ceItem;
	// } NMCOMBOBOXEXW, *PNMCOMBOBOXEXW;

	//// #ifdef UNICODE
	//NMCOMBOBOXEX = NMCOMBOBOXEXW
	//PNMCOMBOBOXEX = PNMCOMBOBOXEXW
	//CBEN_GETDISPINFO = CBEN_GETDISPINFOW
	//// #else
	//NMCOMBOBOXEX = NMCOMBOBOXEXA
	//PNMCOMBOBOXEX = PNMCOMBOBOXEXA
	//CBEN_GETDISPINFO = CBEN_GETDISPINFOA
	//// #endif

	CBEN_GETDISPINFOA = (CBEN_FIRST - 0)
	CBEN_INSERTITEM   = (CBEN_FIRST - 1)
	CBEN_DELETEITEM   = (CBEN_FIRST - 2)
	CBEN_BEGINEDIT    = (CBEN_FIRST - 4)
	CBEN_ENDEDITA     = (CBEN_FIRST - 5)
	CBEN_ENDEDITW     = (CBEN_FIRST - 6)

	CBEN_GETDISPINFOW = (CBEN_FIRST - 7)

	CBEN_DRAGBEGINA = (CBEN_FIRST - 8)
	CBEN_DRAGBEGINW = (CBEN_FIRST - 9)

	//// #ifdef UNICODE
	//CBEN_DRAGBEGIN = CBEN_DRAGBEGINW
	//// #else
	//CBEN_DRAGBEGIN = CBEN_DRAGBEGINA
	//// #endif

	// // lParam specifies why the endedit is happening
	//// #ifdef UNICODE
	//CBEN_ENDEDIT = CBEN_ENDEDITW
	//// #else
	//CBEN_ENDEDIT = CBEN_ENDEDITA
	//// #endif

	CBENF_KILLFOCUS = 1
	CBENF_RETURN    = 2
	CBENF_ESCAPE    = 3
	CBENF_DROPDOWN  = 4

	CBEMAXSTRLEN = 260

	// // CBEN_DRAGBEGIN sends this information ...

	// typedef struct {
	//     NMHDR hdr;
	//     int   iItemid;
	//     WCHAR szText[CBEMAXSTRLEN];
	// }NMCBEDRAGBEGINW, *LPNMCBEDRAGBEGINW, *PNMCBEDRAGBEGINW;

	// typedef struct {
	//     NMHDR hdr;
	//     int   iItemid;
	//     char szText[CBEMAXSTRLEN];
	// }NMCBEDRAGBEGINA, *LPNMCBEDRAGBEGINA, *PNMCBEDRAGBEGINA;

	//// #ifdef UNICODE
	//NMCBEDRAGBEGIN = NMCBEDRAGBEGINW
	//LPNMCBEDRAGBEGIN = LPNMCBEDRAGBEGINW
	//PNMCBEDRAGBEGIN = PNMCBEDRAGBEGINW
	//// #else
	//NMCBEDRAGBEGIN = NMCBEDRAGBEGINA
	//LPNMCBEDRAGBEGIN = LPNMCBEDRAGBEGINA
	//PNMCBEDRAGBEGIN = PNMCBEDRAGBEGINA
	//// #endif

	// // CBEN_ENDEDIT sends this information...
	// // fChanged if the user actually did anything
	// // iNewSelection gives what would be the new selection unless the notify is failed
	// //                      iNewSelection may be CB_ERR if there's no match
	// typedef struct {
	//         NMHDR hdr;
	//         BOOL fChanged;
	//         int iNewSelection;
	//         WCHAR szText[CBEMAXSTRLEN];
	//         int iWhy;
	// } NMCBEENDEDITW, *LPNMCBEENDEDITW, *PNMCBEENDEDITW;

	// typedef struct {
	//         NMHDR hdr;
	//         BOOL fChanged;
	//         int iNewSelection;
	//         char szText[CBEMAXSTRLEN];
	//         int iWhy;
	// } NMCBEENDEDITA, *LPNMCBEENDEDITA,*PNMCBEENDEDITA;

	//// #ifdef UNICODE
	//NMCBEENDEDIT = NMCBEENDEDITW
	//LPNMCBEENDEDIT = LPNMCBEENDEDITW
	//PNMCBEENDEDIT = PNMCBEENDEDITW
	//// #else
	//NMCBEENDEDIT = NMCBEENDEDITA
	//LPNMCBEENDEDIT = LPNMCBEENDEDITA
	//PNMCBEENDEDIT = PNMCBEENDEDITA
	//// #endif

	// #endif

	//====== TAB CONTROL ==========================================================

	// #ifndef NOTABCONTROL

	// #ifdef _WIN32

	// #define WC_TABCONTROLA          "SysTabControl32"
	WC_TABCONTROLW = "SysTabControl32"

	// #ifdef UNICODE
	WC_TABCONTROL = WC_TABCONTROLW
	// #else
	//WC_TABCONTROL = WC_TABCONTROLA
	// #endif

	// #else
	// #define WC_TABCONTROL           "SysTabControl"
	// #endif

	// // begin_r_commctrl

	TCS_SCROLLOPPOSITE    = 0 // assumes multiline tab
	TCS_BOTTOM            = 0
	TCS_RIGHT             = 0
	TCS_MULTISELECT       = 0 // allow multi-select in button mode
	TCS_FLATBUTTONS       = 0
	TCS_FORCEICONLEFT     = 0
	TCS_FORCELABELLEFT    = 0
	TCS_HOTTRACK          = 0
	TCS_VERTICAL          = 0
	TCS_TABS              = 0
	TCS_BUTTONS           = 0
	TCS_SINGLELINE        = 0
	TCS_MULTILINE         = 0
	TCS_RIGHTJUSTIFY      = 0
	TCS_FIXEDWIDTH        = 0
	TCS_RAGGEDRIGHT       = 0
	TCS_FOCUSONBUTTONDOWN = 0
	TCS_OWNERDRAWFIXED    = 0
	TCS_TOOLTIPS          = 0
	TCS_FOCUSNEVER        = 0

	// // end_r_commctrl

	// // EX styles for use with TCM_SETEXTENDEDSTYLE
	TCS_EX_FLATSEPARATORS = 0
	TCS_EX_REGISTERDROP   = 0

	TCM_GETIMAGELIST = (TCM_FIRST + 2)
	// #define TabCtrl_GetImageList(hwnd) \
	//     (HIMAGELIST)SNDMSG((hwnd), TCM_GETIMAGELIST, 0, 0L)

	TCM_SETIMAGELIST = (TCM_FIRST + 3)
	// #define TabCtrl_SetImageList(hwnd, himl) \
	//     (HIMAGELIST)SNDMSG((hwnd), TCM_SETIMAGELIST, 0, (LPARAM)(HIMAGELIST)(himl))

	TCM_GETITEMCOUNT = (TCM_FIRST + 4)
	// #define TabCtrl_GetItemCount(hwnd) \
	//     (int)SNDMSG((hwnd), TCM_GETITEMCOUNT, 0, 0L)

	TCIF_TEXT       = 0x0001
	TCIF_IMAGE      = 0x0002
	TCIF_RTLREADING = 0x0004
	TCIF_PARAM      = 0x0008
	TCIF_STATE      = 0x0010

	TCIS_BUTTONPRESSED = 0
	TCIS_HIGHLIGHTED   = 0

	//TC_ITEMHEADERA = TCITEMHEADERA
	//TC_ITEMHEADERW = TCITEMHEADERW
	//TC_ITEMHEADER = TCITEMHEADER

	// typedef struct tagTCITEMHEADERA
	// {
	//     UINT mask;
	//     UINT lpReserved1;
	//     UINT lpReserved2;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iImage;
	// } TCITEMHEADERA, *LPTCITEMHEADERA;

	// typedef struct tagTCITEMHEADERW
	// {
	//     UINT mask;
	//     UINT lpReserved1;
	//     UINT lpReserved2;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iImage;
	// } TCITEMHEADERW, *LPTCITEMHEADERW;

	//// #ifdef UNICODE
	//TCITEMHEADER = TCITEMHEADERW
	//LPTCITEMHEADER = LPTCITEMHEADERW
	//// #else
	//TCITEMHEADER = TCITEMHEADERA
	//LPTCITEMHEADER = LPTCITEMHEADERA
	//// #endif

	//TC_ITEMA = TCITEMA
	//TC_ITEMW = TCITEMW
	//TC_ITEM = TCITEM

	// typedef struct tagTCITEMA
	// {
	//     UINT mask;
	//     DWORD dwState;
	//     DWORD dwStateMask;
	//     LPSTR pszText;
	//     int cchTextMax;
	//     int iImage;

	//     LPARAM lParam;
	// } TCITEMA, *LPTCITEMA;

	// typedef struct tagTCITEMW
	// {
	//     UINT mask;
	//     DWORD dwState;
	//     DWORD dwStateMask;
	//     LPWSTR pszText;
	//     int cchTextMax;
	//     int iImage;

	//     LPARAM lParam;
	// } TCITEMW, *LPTCITEMW;

	//// #ifdef UNICODE
	//TCITEM = TCITEMW
	//LPTCITEM = LPTCITEMW
	//// #else
	//TCITEM = TCITEMA
	//LPTCITEM = LPTCITEMA
	//// #endif

	TCM_GETITEMA = (TCM_FIRST + 5)
	TCM_GETITEMW = (TCM_FIRST + 60)

	// #ifdef UNICODE
	TCM_GETITEM = TCM_GETITEMW
	// #else
	//TCM_GETITEM = TCM_GETITEMA
	// #endif

	// #define TabCtrl_GetItem(hwnd, iItem, pitem) \
	//     (BOOL)SNDMSG((hwnd), TCM_GETITEM, (WPARAM)(int)(iItem), (LPARAM)(TC_ITEM *)(pitem))

	TCM_SETITEMA = (TCM_FIRST + 6)
	TCM_SETITEMW = (TCM_FIRST + 61)

	// #ifdef UNICODE
	TCM_SETITEM = TCM_SETITEMW
	// #else
	//TCM_SETITEM = TCM_SETITEMA
	// #endif

	// #define TabCtrl_SetItem(hwnd, iItem, pitem) \
	//     (BOOL)SNDMSG((hwnd), TCM_SETITEM, (WPARAM)(int)(iItem), (LPARAM)(TC_ITEM *)(pitem))

	TCM_INSERTITEMA = (TCM_FIRST + 7)
	TCM_INSERTITEMW = (TCM_FIRST + 62)

	// #ifdef UNICODE
	TCM_INSERTITEM = TCM_INSERTITEMW
	// #else
	//TCM_INSERTITEM = TCM_INSERTITEMA
	// #endif

	// #define TabCtrl_InsertItem(hwnd, iItem, pitem)   \
	//     (int)SNDMSG((hwnd), TCM_INSERTITEM, (WPARAM)(int)(iItem), (LPARAM)(const TC_ITEM *)(pitem))

	TCM_DELETEITEM = (TCM_FIRST + 8)
	// #define TabCtrl_DeleteItem(hwnd, i) \
	//     (BOOL)SNDMSG((hwnd), TCM_DELETEITEM, (WPARAM)(int)(i), 0L)

	TCM_DELETEALLITEMS = (TCM_FIRST + 9)
	// #define TabCtrl_DeleteAllItems(hwnd) \
	//     (BOOL)SNDMSG((hwnd), TCM_DELETEALLITEMS, 0, 0L)

	TCM_GETITEMRECT = (TCM_FIRST + 10)
	// #define TabCtrl_GetItemRect(hwnd, i, prc) \
	//     (BOOL)SNDMSG((hwnd), TCM_GETITEMRECT, (WPARAM)(int)(i), (LPARAM)(RECT *)(prc))

	TCM_GETCURSEL = (TCM_FIRST + 11)
	// #define TabCtrl_GetCurSel(hwnd) \
	//     (int)SNDMSG((hwnd), TCM_GETCURSEL, 0, 0)

	TCM_SETCURSEL = (TCM_FIRST + 12)
	// #define TabCtrl_SetCurSel(hwnd, i) \
	//     (int)SNDMSG((hwnd), TCM_SETCURSEL, (WPARAM)(i), 0)

	TCHT_NOWHERE     = 0
	TCHT_ONITEMICON  = 0
	TCHT_ONITEMLABEL = 0
	TCHT_ONITEM      = (TCHT_ONITEMICON | TCHT_ONITEMLABEL)

	//LPTC_HITTESTINFO = LPTCHITTESTINFO
	//TC_HITTESTINFO = TCHITTESTINFO

	// typedef struct tagTCHITTESTINFO
	// {
	//     POINT pt;
	//     UINT flags;
	// } TCHITTESTINFO, *LPTCHITTESTINFO;

	TCM_HITTEST = (TCM_FIRST + 13)
	// #define TabCtrl_HitTest(hwndTC, pinfo) \
	//     (int)SNDMSG((hwndTC), TCM_HITTEST, 0, (LPARAM)(TC_HITTESTINFO *)(pinfo))

	TCM_SETITEMEXTRA = (TCM_FIRST + 14)
	// #define TabCtrl_SetItemExtra(hwndTC, cb) \
	//     (BOOL)SNDMSG((hwndTC), TCM_SETITEMEXTRA, (WPARAM)(cb), 0L)

	TCM_ADJUSTRECT = (TCM_FIRST + 40)
	// #define TabCtrl_AdjustRect(hwnd, bLarger, prc) \
	//     (int)SNDMSG(hwnd, TCM_ADJUSTRECT, (WPARAM)(BOOL)(bLarger), (LPARAM)(RECT *)(prc))

	TCM_SETITEMSIZE = (TCM_FIRST + 41)
	// #define TabCtrl_SetItemSize(hwnd, x, y) \
	//     (DWORD)SNDMSG((hwnd), TCM_SETITEMSIZE, 0, MAKELPARAM(x,y))

	TCM_REMOVEIMAGE = (TCM_FIRST + 42)
	// #define TabCtrl_RemoveImage(hwnd, i) \
	//         (void)SNDMSG((hwnd), TCM_REMOVEIMAGE, i, 0L)

	TCM_SETPADDING = (TCM_FIRST + 43)
	// #define TabCtrl_SetPadding(hwnd,  cx, cy) \
	//         (void)SNDMSG((hwnd), TCM_SETPADDING, 0, MAKELPARAM(cx, cy))

	TCM_GETROWCOUNT = (TCM_FIRST + 44)
	// #define TabCtrl_GetRowCount(hwnd) \
	//         (int)SNDMSG((hwnd), TCM_GETROWCOUNT, 0, 0L)

	TCM_GETTOOLTIPS = (TCM_FIRST + 45)
	// #define TabCtrl_GetToolTips(hwnd) \
	//         (HWND)SNDMSG((hwnd), TCM_GETTOOLTIPS, 0, 0L)

	TCM_SETTOOLTIPS = (TCM_FIRST + 46)
	// #define TabCtrl_SetToolTips(hwnd, hwndTT) \
	//         (void)SNDMSG((hwnd), TCM_SETTOOLTIPS, (WPARAM)(hwndTT), 0L)

	TCM_GETCURFOCUS = (TCM_FIRST + 47)
	// #define TabCtrl_GetCurFocus(hwnd) \
	//     (int)SNDMSG((hwnd), TCM_GETCURFOCUS, 0, 0)

	TCM_SETCURFOCUS = (TCM_FIRST + 48)
	// #define TabCtrl_SetCurFocus(hwnd, i) \
	//     SNDMSG((hwnd),TCM_SETCURFOCUS, i, 0)

	TCM_SETMINTABWIDTH = (TCM_FIRST + 49)
	// #define TabCtrl_SetMinTabWidth(hwnd, x) \
	//         (int)SNDMSG((hwnd), TCM_SETMINTABWIDTH, 0, x)

	TCM_DESELECTALL = (TCM_FIRST + 50)
	// #define TabCtrl_DeselectAll(hwnd, fExcludeFocus)\
	//         (void)SNDMSG((hwnd), TCM_DESELECTALL, fExcludeFocus, 0)

	TCM_HIGHLIGHTITEM = (TCM_FIRST + 51)
	// #define TabCtrl_HighlightItem(hwnd, i, fHighlight) \
	//     (BOOL)SNDMSG((hwnd), TCM_HIGHLIGHTITEM, (WPARAM)(i), (LPARAM)MAKELONG (fHighlight, 0))

	TCM_SETEXTENDEDSTYLE = (TCM_FIRST + 52) // optional wParam == mask
	// #define TabCtrl_SetExtendedStyle(hwnd, dw)\
	//         (DWORD)SNDMSG((hwnd), TCM_SETEXTENDEDSTYLE, 0, dw)

	TCM_GETEXTENDEDSTYLE = (TCM_FIRST + 53)
	// #define TabCtrl_GetExtendedStyle(hwnd)\
	//         (DWORD)SNDMSG((hwnd), TCM_GETEXTENDEDSTYLE, 0, 0)

	TCM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	// #define TabCtrl_SetUnicodeFormat(hwnd, fUnicode)  \
	//     (BOOL)SNDMSG((hwnd), TCM_SETUNICODEFORMAT, (WPARAM)(fUnicode), 0)

	TCM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	// #define TabCtrl_GetUnicodeFormat(hwnd)  \
	//     (BOOL)SNDMSG((hwnd), TCM_GETUNICODEFORMAT, 0, 0)

	TCN_KEYDOWN = (TCN_FIRST - 0)

	//TC_KEYDOWN = NMTCKEYDOWN

	// #ifdef _WIN32
	// #include <pshpack1.h>
	// #endif

	// typedef struct tagTCKEYDOWN
	// {
	//     NMHDR hdr;
	//     WORD wVKey;
	//     UINT flags;
	// } NMTCKEYDOWN;

	// #ifdef _WIN32
	// #include <poppack.h>
	// #endif

	TCN_SELCHANGE   = (TCN_FIRST - 1)
	TCN_SELCHANGING = (TCN_FIRST - 2)
	TCN_GETOBJECT   = (TCN_FIRST - 3)
	TCN_FOCUSCHANGE = (TCN_FIRST - 4)
	// #endif      // NOTABCONTROL

	//====== ANIMATE CONTROL ======================================================

	// #ifndef NOANIMATE

	// #ifdef _WIN32

	ANIMATE_CLASSW = "SysAnimate32"
	// #define ANIMATE_CLASSA          "SysAnimate32"

	// #ifdef UNICODE
	ANIMATE_CLASS = ANIMATE_CLASSW
	// #else
	//ANIMATE_CLASS = ANIMATE_CLASSA
	// #endif

	// // begin_r_commctrl

	ACS_CENTER      = 0
	ACS_TRANSPARENT = 0
	ACS_AUTOPLAY    = 0
	ACS_TIMER       = 0 // don't use threads... use timers

	// // end_r_commctrl

	ACM_OPENA = (WM_USER + 100)
	ACM_OPENW = (WM_USER + 103)

	// #ifdef UNICODE
	ACM_OPEN = ACM_OPENW
	// #else
	//ACM_OPEN = ACM_OPENA
	// #endif

	ACM_PLAY      = (WM_USER + 101)
	ACM_STOP      = (WM_USER + 102)
	ACM_ISPLAYING = (WM_USER + 104)

	ACN_START = 1
	ACN_STOP  = 2

	// #define Animate_Create(hwndP, id, dwStyle, hInstance)   \
	//             CreateWindow(ANIMATE_CLASS, NULL,           \
	//                 dwStyle, 0, 0, 0, 0, hwndP, (HMENU)(id), hInstance, NULL)

	// #define Animate_Open(hwnd, szName)          (BOOL)SNDMSG(hwnd, ACM_OPEN, 0, (LPARAM)(LPTSTR)(szName))
	// #define Animate_OpenEx(hwnd, hInst, szName) (BOOL)SNDMSG(hwnd, ACM_OPEN, (WPARAM)(hInst), (LPARAM)(LPTSTR)(szName))
	// #define Animate_Play(hwnd, from, to, rep)   (BOOL)SNDMSG(hwnd, ACM_PLAY, (WPARAM)(rep), (LPARAM)MAKELONG(from, to))
	// #define Animate_Stop(hwnd)                  (BOOL)SNDMSG(hwnd, ACM_STOP, 0, 0)
	// #define Animate_IsPlaying(hwnd)             (BOOL)SNDMSG(hwnd, ACM_ISPLAYING, 0, 0)
	// #define Animate_Close(hwnd)                 Animate_Open(hwnd, NULL)
	// #define Animate_Seek(hwnd, frame)           Animate_Play(hwnd, frame, frame, 1)
	// #endif

	// #endif      // NOANIMATE

	//====== MONTHCAL CONTROL ======================================================

	// #ifndef NOMONTHCAL
	// #ifdef _WIN32

	MONTHCAL_CLASSW = "SysMonthCal32"
	// #define MONTHCAL_CLASSA          "SysMonthCal32"

	// #ifdef UNICODE
	MONTHCAL_CLASS = MONTHCAL_CLASSW
	// #else
	//MONTHCAL_CLASS = MONTHCAL_CLASSA
	// #endif

	// // bit-packed array of "bold" info for a month
	// // if a bit is on, that day is drawn bold
	// typedef DWORD MONTHDAYSTATE, *LPMONTHDAYSTATE;

	MCM_FIRST = 0

	// // BOOL MonthCal_GetCurSel(HWND hmc, LPSYSTEMTIME pst)
	// //   returns FALSE if MCS_MULTISELECT
	// //   returns TRUE and sets *pst to the currently selected date otherwise
	MCM_GETCURSEL = (MCM_FIRST + 1)
	// #define MonthCal_GetCurSel(hmc, pst)    (BOOL)SNDMSG(hmc, MCM_GETCURSEL, 0, (LPARAM)(pst))

	// // BOOL MonthCal_SetCurSel(HWND hmc, LPSYSTEMTIME pst)
	// //   returns FALSE if MCS_MULTISELECT
	// //   returns TURE and sets the currently selected date to *pst otherwise
	MCM_SETCURSEL = (MCM_FIRST + 2)
	// #define MonthCal_SetCurSel(hmc, pst)    (BOOL)SNDMSG(hmc, MCM_SETCURSEL, 0, (LPARAM)(pst))

	// // DWORD MonthCal_GetMaxSelCount(HWND hmc)
	// //   returns the maximum number of selectable days allowed
	MCM_GETMAXSELCOUNT = (MCM_FIRST + 3)
	// #define MonthCal_GetMaxSelCount(hmc)    (DWORD)SNDMSG(hmc, MCM_GETMAXSELCOUNT, 0, 0L)

	// // BOOL MonthCal_SetMaxSelCount(HWND hmc, UINT n)
	// //   sets the max number days that can be selected iff MCS_MULTISELECT
	MCM_SETMAXSELCOUNT = (MCM_FIRST + 4)
	// #define MonthCal_SetMaxSelCount(hmc, n) (BOOL)SNDMSG(hmc, MCM_SETMAXSELCOUNT, (WPARAM)(n), 0L)

	// // BOOL MonthCal_GetSelRange(HWND hmc, LPSYSTEMTIME rgst)
	// //   sets rgst[0] to the first day of the selection range
	// //   sets rgst[1] to the last day of the selection range
	MCM_GETSELRANGE = (MCM_FIRST + 5)
	// #define MonthCal_GetSelRange(hmc, rgst) SNDMSG(hmc, MCM_GETSELRANGE, 0, (LPARAM)(rgst))

	// // BOOL MonthCal_SetSelRange(HWND hmc, LPSYSTEMTIME rgst)
	// //   selects the range of days from rgst[0] to rgst[1]
	MCM_SETSELRANGE = (MCM_FIRST + 6)
	// #define MonthCal_SetSelRange(hmc, rgst) SNDMSG(hmc, MCM_SETSELRANGE, 0, (LPARAM)(rgst))

	// // DWORD MonthCal_GetMonthRange(HWND hmc, DWORD gmr, LPSYSTEMTIME rgst)
	// //   if rgst specified, sets rgst[0] to the starting date and
	// //      and rgst[1] to the ending date of the the selectable (non-grayed)
	// //      days if GMR_VISIBLE or all the displayed days (including grayed)
	// //      if GMR_DAYSTATE.
	// //   returns the number of months spanned by the above range.
	MCM_GETMONTHRANGE = (MCM_FIRST + 7)
	// #define MonthCal_GetMonthRange(hmc, gmr, rgst)  (DWORD)SNDMSG(hmc, MCM_GETMONTHRANGE, (WPARAM)(gmr), (LPARAM)(rgst))

	// // BOOL MonthCal_SetDayState(HWND hmc, int cbds, DAYSTATE *rgds)
	// //   cbds is the count of DAYSTATE items in rgds and it must be equal
	// //   to the value returned from MonthCal_GetMonthRange(hmc, GMR_DAYSTATE, NULL)
	// //   This sets the DAYSTATE bits for each month (grayed and non-grayed
	// //   days) displayed in the calendar. The first bit in a month's DAYSTATE
	// //   corresponts to bolding day 1, the second bit affects day 2, etc.
	MCM_SETDAYSTATE = (MCM_FIRST + 8)
	// #define MonthCal_SetDayState(hmc, cbds, rgds)   SNDMSG(hmc, MCM_SETDAYSTATE, (WPARAM)(cbds), (LPARAM)(rgds))

	// // BOOL MonthCal_GetMinReqRect(HWND hmc, LPRECT prc)
	// //   sets *prc the minimal size needed to display one month
	// //   To display two months, undo the AdjustWindowRect calculation already done to
	// //   this rect, double the width, and redo the AdjustWindowRect calculation --
	// //   the monthcal control will display two calendars in this window (if you also
	// //   double the vertical size, you will get 4 calendars)
	// //   NOTE: if you want to gurantee that the "Today" string is not clipped,
	// //   get the MCM_GETMAXTODAYWIDTH and use the max of that width and this width
	MCM_GETMINREQRECT = (MCM_FIRST + 9)
	// #define MonthCal_GetMinReqRect(hmc, prc)        SNDMSG(hmc, MCM_GETMINREQRECT, 0, (LPARAM)(prc))

	// // set colors to draw control with -- see MCSC_ bits below
	MCM_SETCOLOR = (MCM_FIRST + 10)
	// #define MonthCal_SetColor(hmc, iColor, clr) SNDMSG(hmc, MCM_SETCOLOR, iColor, clr)

	MCM_GETCOLOR = (MCM_FIRST + 11)
	// #define MonthCal_GetColor(hmc, iColor) SNDMSG(hmc, MCM_GETCOLOR, iColor, 0)

	MCSC_BACKGROUND   = 0 // the background color (between months)
	MCSC_TEXT         = 1 // the dates
	MCSC_TITLEBK      = 2 // background of the title
	MCSC_TITLETEXT    = 3
	MCSC_MONTHBK      = 4 // background within the month cal
	MCSC_TRAILINGTEXT = 5 // the text color of header & trailing days

	// // set what day is "today"   send NULL to revert back to real date
	MCM_SETTODAY = (MCM_FIRST + 12)
	// #define MonthCal_SetToday(hmc, pst)             SNDMSG(hmc, MCM_SETTODAY, 0, (LPARAM)(pst))

	// // get what day is "today"
	// // returns BOOL for success/failure
	MCM_GETTODAY = (MCM_FIRST + 13)
	// #define MonthCal_GetToday(hmc, pst)             (BOOL)SNDMSG(hmc, MCM_GETTODAY, 0, (LPARAM)(pst))

	// // determine what pinfo->pt is over
	MCM_HITTEST = (MCM_FIRST + 14)
	// #define MonthCal_HitTest(hmc, pinfo) \
	//         SNDMSG(hmc, MCM_HITTEST, 0, (LPARAM)(PMCHITTESTINFO)(pinfo))

	// typedef struct {
	//         UINT cbSize;
	//         POINT pt;

	//         UINT uHit;   // out param
	//         SYSTEMTIME st;
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	//         RECT rc;
	//         int iOffset;
	//         int iRow;
	//         int iCol;
	// #endif
	// } MCHITTESTINFO, *PMCHITTESTINFO;

	//MCHITTESTINFO_V1_SIZE = CCSIZEOF_STRUCT(MCHITTESTINFO, st)

	MCHT_TITLE     = 0
	MCHT_CALENDAR  = 0
	MCHT_TODAYLINK = 0

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	MCHT_CALENDARCONTROL = 0
	// #endif

	MCHT_NEXT = 0 // these indicate that hitting
	MCHT_PREV = 0 // here will go to the next/prev month

	MCHT_NOWHERE = 0

	MCHT_TITLEBK      = (MCHT_TITLE)
	MCHT_TITLEMONTH   = (MCHT_TITLE | 0x0001)
	MCHT_TITLEYEAR    = (MCHT_TITLE | 0x0002)
	MCHT_TITLEBTNNEXT = (MCHT_TITLE | MCHT_NEXT | 0x0003)
	MCHT_TITLEBTNPREV = (MCHT_TITLE | MCHT_PREV | 0x0003)

	MCHT_CALENDARBK       = (MCHT_CALENDAR)
	MCHT_CALENDARDATE     = (MCHT_CALENDAR | 0x0001)
	MCHT_CALENDARDATENEXT = (MCHT_CALENDARDATE | MCHT_NEXT)
	MCHT_CALENDARDATEPREV = (MCHT_CALENDARDATE | MCHT_PREV)
	MCHT_CALENDARDAY      = (MCHT_CALENDAR | 0x0002)
	MCHT_CALENDARWEEKNUM  = (MCHT_CALENDAR | 0x0003)
	MCHT_CALENDARDATEMIN  = (MCHT_CALENDAR | 0x0004)
	MCHT_CALENDARDATEMAX  = (MCHT_CALENDAR | 0x0005)

	// // set first day of week to iDay:
	// // 0 for Monday, 1 for Tuesday, ..., 6 for Sunday
	// // -1 for means use locale info
	MCM_SETFIRSTDAYOFWEEK = (MCM_FIRST + 15)
	// #define MonthCal_SetFirstDayOfWeek(hmc, iDay) \
	//         SNDMSG(hmc, MCM_SETFIRSTDAYOFWEEK, 0, iDay)

	// // DWORD result...  low word has the day.  high word is bool if this is app set
	// or not (FALSE == using locale info)
	MCM_GETFIRSTDAYOFWEEK = (MCM_FIRST + 16)
	// #define MonthCal_GetFirstDayOfWeek(hmc) \
	//         (DWORD)SNDMSG(hmc, MCM_GETFIRSTDAYOFWEEK, 0, 0)

	// // DWORD MonthCal_GetRange(HWND hmc, LPSYSTEMTIME rgst)
	// //   modifies rgst[0] to be the minimum ALLOWABLE systemtime (or 0 if no minimum)
	// //   modifies rgst[1] to be the maximum ALLOWABLE systemtime (or 0 if no maximum)
	// //   returns GDTR_MIN|GDTR_MAX if there is a minimum|maximum limit
	MCM_GETRANGE = (MCM_FIRST + 17)
	// #define MonthCal_GetRange(hmc, rgst) \
	//         (DWORD)SNDMSG(hmc, MCM_GETRANGE, 0, (LPARAM)(rgst))

	// // BOOL MonthCal_SetRange(HWND hmc, DWORD gdtr, LPSYSTEMTIME rgst)
	// //   if GDTR_MIN, sets the minimum ALLOWABLE systemtime to rgst[0], otherwise removes minimum
	// //   if GDTR_MAX, sets the maximum ALLOWABLE systemtime to rgst[1], otherwise removes maximum
	// //   returns TRUE on success, FALSE on error (such as invalid parameters)
	MCM_SETRANGE = (MCM_FIRST + 18)
	// #define MonthCal_SetRange(hmc, gd, rgst) \
	//         (BOOL)SNDMSG(hmc, MCM_SETRANGE, (WPARAM)(gd), (LPARAM)(rgst))

	// // int MonthCal_GetMonthDelta(HWND hmc)
	// //   returns the number of months one click on a next/prev button moves by
	MCM_GETMONTHDELTA = (MCM_FIRST + 19)
	// #define MonthCal_GetMonthDelta(hmc) \
	//         (int)SNDMSG(hmc, MCM_GETMONTHDELTA, 0, 0)

	// // int MonthCal_SetMonthDelta(HWND hmc, int n)
	//   sets the month delta to n. n==0 reverts to moving by a page of months
	// //   returns the previous value of n.
	MCM_SETMONTHDELTA = (MCM_FIRST + 20)
	// #define MonthCal_SetMonthDelta(hmc, n) \
	//         (int)SNDMSG(hmc, MCM_SETMONTHDELTA, n, 0)

	// // DWORD MonthCal_GetMaxTodayWidth(HWND hmc, LPSIZE psz)
	// //   sets *psz to the maximum width/height of the "Today" string displayed
	// //   at the bottom of the calendar (as long as MCS_NOTODAY is not specified)
	MCM_GETMAXTODAYWIDTH = (MCM_FIRST + 21)
	// #define MonthCal_GetMaxTodayWidth(hmc) \
	//         (DWORD)SNDMSG(hmc, MCM_GETMAXTODAYWIDTH, 0, 0)

	MCM_SETUNICODEFORMAT = CCM_SETUNICODEFORMAT
	// #define MonthCal_SetUnicodeFormat(hwnd, fUnicode)  \
	//     (BOOL)SNDMSG((hwnd), MCM_SETUNICODEFORMAT, (WPARAM)(fUnicode), 0)

	MCM_GETUNICODEFORMAT = CCM_GETUNICODEFORMAT
	// #define MonthCal_GetUnicodeFormat(hwnd)  \
	//     (BOOL)SNDMSG((hwnd), MCM_GETUNICODEFORMAT, 0, 0)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	// // View
	MCMV_MONTH   = 0
	MCMV_YEAR    = 1
	MCMV_DECADE  = 2
	MCMV_CENTURY = 3
	MCMV_MAX     = MCMV_CENTURY

	MCM_GETCURRENTVIEW = (MCM_FIRST + 22)
	// #define MonthCal_GetCurrentView(hmc) \
	//         (DWORD)SNDMSG(hmc, MCM_GETCURRENTVIEW, 0, 0)

	MCM_GETCALENDARCOUNT = (MCM_FIRST + 23)
	// #define MonthCal_GetCalendarCount(hmc) \
	//         (DWORD)SNDMSG(hmc, MCM_GETCALENDARCOUNT, 0, 0)

	// // Part
	MCGIP_CALENDARCONTROL = 0
	MCGIP_NEXT            = 1
	MCGIP_PREV            = 2
	MCGIP_FOOTER          = 3
	MCGIP_CALENDAR        = 4
	MCGIP_CALENDARHEADER  = 5
	MCGIP_CALENDARBODY    = 6
	MCGIP_CALENDARROW     = 7
	MCGIP_CALENDARCELL    = 8

	MCGIF_DATE = 0
	MCGIF_RECT = 0
	MCGIF_NAME = 0

	// // Note: iRow of -1 refers to the row header and iCol of -1 refers to the col header.
	// typedef struct tagMCGRIDINFO {
	//     UINT cbSize;
	//     DWORD dwPart;
	//     DWORD dwFlags;
	//     int iCalendar;
	//     int iRow;
	//     int iCol;
	//     BOOL bSelected;
	//     SYSTEMTIME stStart;
	//     SYSTEMTIME stEnd;
	//     RECT rc;
	//     PWSTR pszName;
	//     size_t cchName;
	// } MCGRIDINFO, *PMCGRIDINFO;

	MCM_GETCALENDARGRIDINFO = (MCM_FIRST + 24)
	// #define MonthCal_GetCalendarGridInfo(hmc, pmcGridInfo) \
	//         (BOOL)SNDMSG(hmc, MCM_GETCALENDARGRIDINFO, 0, (LPARAM)(PMCGRIDINFO)(pmcGridInfo))

	MCM_GETCALID = (MCM_FIRST + 27)
	// #define MonthCal_GetCALID(hmc) \
	//         (CALID)SNDMSG(hmc, MCM_GETCALID, 0, 0)

	MCM_SETCALID = (MCM_FIRST + 28)
	// #define MonthCal_SetCALID(hmc, calid) \
	//         SNDMSG(hmc, MCM_SETCALID, (WPARAM)(calid), 0)

	// // Returns the min rect that will fit the max number of calendars for the passed in rect.
	MCM_SIZERECTTOMIN = (MCM_FIRST + 29)
	// #define MonthCal_SizeRectToMin(hmc, prc) \
	//         SNDMSG(hmc, MCM_SIZERECTTOMIN, 0, (LPARAM)(prc))

	MCM_SETCALENDARBORDER = (MCM_FIRST + 30)
	// #define MonthCal_SetCalendarBorder(hmc, fset, xyborder) \
	//         SNDMSG(hmc, MCM_SETCALENDARBORDER, (WPARAM)(fset), (LPARAM)(xyborder))

	MCM_GETCALENDARBORDER = (MCM_FIRST + 31)
	// #define MonthCal_GetCalendarBorder(hmc) \
	//         (int)SNDMSG(hmc, MCM_GETCALENDARBORDER, 0, 0)

	MCM_SETCURRENTVIEW = (MCM_FIRST + 32)
	// #define MonthCal_SetCurrentView(hmc, dwNewView) \
	//         (BOOL)SNDMSG(hmc, MCM_SETCURRENTVIEW, 0, (LPARAM)(dwNewView))

	// #endif

	// // MCN_SELCHANGE is sent whenever the currently displayed date changes
	// // via month change, year change, keyboard navigation, prev/next button
	// //
	// typedef struct tagNMSELCHANGE
	// {
	//     NMHDR           nmhdr;  // this must be first, so we don't break WM_NOTIFY

	//     SYSTEMTIME      stSelStart;
	//     SYSTEMTIME      stSelEnd;
	// } NMSELCHANGE, *LPNMSELCHANGE;

	MCN_SELCHANGE = (MCN_FIRST - 3) // -749

	// // MCN_GETDAYSTATE is sent for MCS_DAYSTATE controls whenever new daystate
	// // information is needed (month or year scroll) to draw bolding information.
	// // The app must fill in cDayState months worth of information starting from
	// // stStart date. The app may fill in the array at prgDayState or change
	// // prgDayState to point to a different array out of which the information
	// // will be copied. (similar to tooltips)
	// //
	// typedef struct tagNMDAYSTATE
	// {
	//     NMHDR           nmhdr;  // this must be first, so we don't break WM_NOTIFY

	//     SYSTEMTIME      stStart;
	//     int             cDayState;

	//     LPMONTHDAYSTATE prgDayState; // points to cDayState MONTHDAYSTATEs
	// } NMDAYSTATE, *LPNMDAYSTATE;

	MCN_GETDAYSTATE = (MCN_FIRST - 1) // -747

	// // MCN_SELECT is sent whenever a selection has occured (via mouse or keyboard)
	// //
	// typedef NMSELCHANGE NMSELECT, *LPNMSELECT;

	MCN_SELECT = (MCN_FIRST) // -746

	// typedef struct tagNMVIEWCHANGE
	// {
	//     NMHDR           nmhdr;  // this must be first, so we don't break WM_NOTIFY
	//     DWORD           dwOldView;
	//     DWORD           dwNewView;
	// } NMVIEWCHANGE, *LPNMVIEWCHANGE;

	MCN_VIEWCHANGE = (MCN_FIRST - 4) // -750

	// // begin_r_commctrl

	MCS_DAYSTATE      = 0
	MCS_MULTISELECT   = 0
	MCS_WEEKNUMBERS   = 0
	MCS_NOTODAYCIRCLE = 0
	MCS_NOTODAY       = 0
	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	MCS_NOTRAILINGDATES  = 0
	MCS_SHORTDAYSOFWEEK  = 0
	MCS_NOSELCHANGEONNAV = 0
	// #endif

	// // end_r_commctrl

	GMR_VISIBLE  = 0 // visible portion of display
	GMR_DAYSTATE = 1 // above plus the grayed out parts of
	//                                 // partially displayed months

	// #endif // _WIN32
	// #endif // NOMONTHCAL

	//====== DATETIMEPICK CONTROL ==================================================

	// #ifndef NODATETIMEPICK
	// #ifdef _WIN32

	DATETIMEPICK_CLASSW = "SysDateTimePick32"
	// #define DATETIMEPICK_CLASSA          "SysDateTimePick32"

	// #ifdef UNICODE
	DATETIMEPICK_CLASS = DATETIMEPICK_CLASSW
	// #else
	//DATETIMEPICK_CLASS = DATETIMEPICK_CLASSA
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)

	// typedef struct tagDATETIMEPICKERINFO
	// {
	//     DWORD cbSize;

	//     RECT rcCheck;
	//     DWORD stateCheck;

	//     RECT rcButton;
	//     DWORD stateButton;

	//     HWND hwndEdit;
	//     HWND hwndUD;
	//     HWND hwndDropDown;
	// } DATETIMEPICKERINFO, *LPDATETIMEPICKERINFO;

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	DTM_FIRST = 0

	// // DWORD DateTimePick_GetSystemtime(HWND hdp, LPSYSTEMTIME pst)
	// //   returns GDT_NONE if "none" is selected (DTS_SHOWNONE only)
	// //   returns GDT_VALID and modifies *pst to be the currently selected value
	DTM_GETSYSTEMTIME = (DTM_FIRST + 1)
	// #define DateTime_GetSystemtime(hdp, pst)    (DWORD)SNDMSG(hdp, DTM_GETSYSTEMTIME, 0, (LPARAM)(pst))

	// // BOOL DateTime_SetSystemtime(HWND hdp, DWORD gd, LPSYSTEMTIME pst)
	//   if gd==GDT_NONE, sets datetimepick to None (DTS_SHOWNONE only)
	//   if gd==GDT_VALID, sets datetimepick to *pst
	// //   returns TRUE on success, FALSE on error (such as bad params)
	DTM_SETSYSTEMTIME = (DTM_FIRST + 2)
	// #define DateTime_SetSystemtime(hdp, gd, pst)    (BOOL)SNDMSG(hdp, DTM_SETSYSTEMTIME, (WPARAM)(gd), (LPARAM)(pst))

	// // DWORD DateTime_GetRange(HWND hdp, LPSYSTEMTIME rgst)
	// //   modifies rgst[0] to be the minimum ALLOWABLE systemtime (or 0 if no minimum)
	// //   modifies rgst[1] to be the maximum ALLOWABLE systemtime (or 0 if no maximum)
	// //   returns GDTR_MIN|GDTR_MAX if there is a minimum|maximum limit
	DTM_GETRANGE = (DTM_FIRST + 3)
	// #define DateTime_GetRange(hdp, rgst)  (DWORD)SNDMSG(hdp, DTM_GETRANGE, 0, (LPARAM)(rgst))

	// // BOOL DateTime_SetRange(HWND hdp, DWORD gdtr, LPSYSTEMTIME rgst)
	// //   if GDTR_MIN, sets the minimum ALLOWABLE systemtime to rgst[0], otherwise removes minimum
	// //   if GDTR_MAX, sets the maximum ALLOWABLE systemtime to rgst[1], otherwise removes maximum
	// //   returns TRUE on success, FALSE on error (such as invalid parameters)
	DTM_SETRANGE = (DTM_FIRST + 4)
	// #define DateTime_SetRange(hdp, gd, rgst)  (BOOL)SNDMSG(hdp, DTM_SETRANGE, (WPARAM)(gd), (LPARAM)(rgst))

	// // BOOL DateTime_SetFormat(HWND hdp, LPCTSTR sz)
	// //   sets the display formatting string to sz (see GetDateFormat and GetTimeFormat for valid formatting chars)
	// //   NOTE: 'X' is a valid formatting character which indicates that the application
	// //   will determine how to display information. Such apps must support DTN_WMKEYDOWN,
	// //   DTN_FORMAT, and DTN_FORMATQUERY.
	DTM_SETFORMATA = (DTM_FIRST + 5)
	DTM_SETFORMATW = (DTM_FIRST + 50)

	// #ifdef UNICODE
	DTM_SETFORMAT = DTM_SETFORMATW
	// #else
	//DTM_SETFORMAT = DTM_SETFORMATA
	// #endif

	// #define DateTime_SetFormat(hdp, sz)  (BOOL)SNDMSG(hdp, DTM_SETFORMAT, 0, (LPARAM)(sz))

	DTM_SETMCCOLOR = (DTM_FIRST + 6)
	// #define DateTime_SetMonthCalColor(hdp, iColor, clr) SNDMSG(hdp, DTM_SETMCCOLOR, iColor, clr)

	DTM_GETMCCOLOR = (DTM_FIRST + 7)
	// #define DateTime_GetMonthCalColor(hdp, iColor) SNDMSG(hdp, DTM_GETMCCOLOR, iColor, 0)

	// // HWND DateTime_GetMonthCal(HWND hdp)
	// //   returns the HWND of the MonthCal popup window. Only valid
	// // between DTN_DROPDOWN and DTN_CLOSEUP notifications.
	DTM_GETMONTHCAL = (DTM_FIRST + 8)
	// #define DateTime_GetMonthCal(hdp) (HWND)SNDMSG(hdp, DTM_GETMONTHCAL, 0, 0)

	DTM_SETMCFONT = (DTM_FIRST + 9)
	// #define DateTime_SetMonthCalFont(hdp, hfont, fRedraw) SNDMSG(hdp, DTM_SETMCFONT, (WPARAM)(hfont), (LPARAM)(fRedraw))

	DTM_GETMCFONT = (DTM_FIRST + 10)
	// #define DateTime_GetMonthCalFont(hdp) SNDMSG(hdp, DTM_GETMCFONT, 0, 0)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)

	DTM_SETMCSTYLE = (DTM_FIRST + 11)
	// #define DateTime_SetMonthCalStyle(hdp, dwStyle) SNDMSG(hdp, DTM_SETMCSTYLE, 0, (LPARAM)dwStyle)

	DTM_GETMCSTYLE = (DTM_FIRST + 12)
	// #define DateTime_GetMonthCalStyle(hdp) SNDMSG(hdp, DTM_GETMCSTYLE, 0, 0)

	DTM_CLOSEMONTHCAL = (DTM_FIRST + 13)
	// #define DateTime_CloseMonthCal(hdp) SNDMSG(hdp, DTM_CLOSEMONTHCAL, 0, 0)

	// // DateTime_GetDateTimePickerInfo(HWND hdp, DATETIMEPICKERINFO* pdtpi)
	// // Retrieves information about the selected date time picker.
	DTM_GETDATETIMEPICKERINFO = (DTM_FIRST + 14)
	// #define DateTime_GetDateTimePickerInfo(hdp, pdtpi) SNDMSG(hdp, DTM_GETDATETIMEPICKERINFO, 0, (LPARAM)(pdtpi))

	DTM_GETIDEALSIZE = (DTM_FIRST + 15)
	// #define DateTime_GetIdealSize(hdp, psize) (BOOL)SNDMSG((hdp), DTM_GETIDEALSIZE, 0, (LPARAM)(psize))

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// // begin_r_commctrl

	DTS_UPDOWN                 = 0 // use UPDOWN instead of MONTHCAL
	DTS_SHOWNONE               = 0 // allow a NONE selection
	DTS_SHORTDATEFORMAT        = 0 // use the short date format (app must forward WM_WININICHANGE messages)
	DTS_LONGDATEFORMAT         = 0 // use the long date format (app must forward WM_WININICHANGE messages)
	DTS_SHORTDATECENTURYFORMAT = 0 // short date format with century (app must forward WM_WININICHANGE messages)
	DTS_TIMEFORMAT             = 0 // use the time format (app must forward WM_WININICHANGE messages)
	DTS_APPCANPARSE            = 0 // allow user entered strings (app MUST respond to DTN_USERSTRING)
	DTS_RIGHTALIGN             = 0 // right-align popup instead of left-align it

	// // end_r_commctrl

	DTN_DATETIMECHANGE = (DTN_FIRST2 - 6) // the systemtime has changed, -759
	// typedef struct tagNMDATETIMECHANGE
	// {
	//     NMHDR       nmhdr;
	//     DWORD       dwFlags;    // GDT_VALID or GDT_NONE
	//SYSTEMTIME  st;         // valid iff dwFlags==GDT_VALID
	// } NMDATETIMECHANGE, *LPNMDATETIMECHANGE;

	DTN_USERSTRINGA = (DTN_FIRST2 - 5) // the user has entered a string, -758
	DTN_USERSTRINGW = (DTN_FIRST - 5)  // -745
	// typedef struct tagNMDATETIMESTRINGA
	// {
	//     NMHDR      nmhdr;
	//     LPCSTR     pszUserString;  // string user entered
	//     SYSTEMTIME st;             // app fills this in
	//     DWORD      dwFlags;        // GDT_VALID or GDT_NONE
	// } NMDATETIMESTRINGA, *LPNMDATETIMESTRINGA;

	// typedef struct tagNMDATETIMESTRINGW
	// {
	//     NMHDR      nmhdr;
	//     LPCWSTR    pszUserString;  // string user entered
	//     SYSTEMTIME st;             // app fills this in
	//     DWORD      dwFlags;        // GDT_VALID or GDT_NONE
	// } NMDATETIMESTRINGW, *LPNMDATETIMESTRINGW;

	//// #ifdef UNICODE
	//DTN_USERSTRING = DTN_USERSTRINGW
	//NMDATETIMESTRING = NMDATETIMESTRINGW
	//LPNMDATETIMESTRING = LPNMDATETIMESTRINGW
	//// #else
	//DTN_USERSTRING = DTN_USERSTRINGA
	//NMDATETIMESTRING = NMDATETIMESTRINGA
	//LPNMDATETIMESTRING = LPNMDATETIMESTRINGA
	//// #endif

	DTN_WMKEYDOWNA = (DTN_FIRST2 - 4) // modify keydown on app format field (X), , -757
	DTN_WMKEYDOWNW = (DTN_FIRST - 4)  // -744
	// typedef struct tagNMDATETIMEWMKEYDOWNA
	// {
	//     NMHDR      nmhdr;
	//     int        nVirtKey;  // virtual key code of WM_KEYDOWN which MODIFIES an X field
	//     LPCSTR     pszFormat; // format substring
	//     SYSTEMTIME st;        // current systemtime, app should modify based on key
	// } NMDATETIMEWMKEYDOWNA, *LPNMDATETIMEWMKEYDOWNA;

	// typedef struct tagNMDATETIMEWMKEYDOWNW
	// {
	//     NMHDR      nmhdr;
	//     int        nVirtKey;  // virtual key code of WM_KEYDOWN which MODIFIES an X field
	//     LPCWSTR    pszFormat; // format substring
	//     SYSTEMTIME st;        // current systemtime, app should modify based on key
	// } NMDATETIMEWMKEYDOWNW, *LPNMDATETIMEWMKEYDOWNW;

	//// #ifdef UNICODE
	//DTN_WMKEYDOWN = DTN_WMKEYDOWNW
	//NMDATETIMEWMKEYDOWN = NMDATETIMEWMKEYDOWNW
	//LPNMDATETIMEWMKEYDOWN = LPNMDATETIMEWMKEYDOWNW
	//// #else
	//DTN_WMKEYDOWN = DTN_WMKEYDOWNA
	//NMDATETIMEWMKEYDOWN = NMDATETIMEWMKEYDOWNA
	//LPNMDATETIMEWMKEYDOWN = LPNMDATETIMEWMKEYDOWNA
	//// #endif

	DTN_FORMATA = (DTN_FIRST2 - 3) // query display for app format field (X), -756
	DTN_FORMATW = (DTN_FIRST - 3)  // -743
	// typedef struct tagNMDATETIMEFORMATA
	// {
	//     NMHDR nmhdr;
	//     LPCSTR  pszFormat;   // format substring
	//     SYSTEMTIME st;       // current systemtime
	//     LPCSTR pszDisplay;   // string to display
	//     CHAR szDisplay[64];  // buffer pszDisplay originally points at
	// } NMDATETIMEFORMATA, *LPNMDATETIMEFORMATA;

	// typedef struct tagNMDATETIMEFORMATW
	// {
	//     NMHDR nmhdr;
	//     LPCWSTR pszFormat;   // format substring
	//     SYSTEMTIME st;       // current systemtime
	//     LPCWSTR pszDisplay;  // string to display
	//     WCHAR szDisplay[64]; // buffer pszDisplay originally points at
	// } NMDATETIMEFORMATW, *LPNMDATETIMEFORMATW;

	//// #ifdef UNICODE
	//DTN_FORMAT = DTN_FORMATW
	//NMDATETIMEFORMAT = NMDATETIMEFORMATW
	//LPNMDATETIMEFORMAT = LPNMDATETIMEFORMATW
	//// #else
	//DTN_FORMAT = DTN_FORMATA
	//NMDATETIMEFORMAT = NMDATETIMEFORMATA
	//LPNMDATETIMEFORMAT = LPNMDATETIMEFORMATA
	//// #endif

	DTN_FORMATQUERYA = (DTN_FIRST2 - 2) // query formatting info for app format field (X), -755
	DTN_FORMATQUERYW = (DTN_FIRST - 2)  // -742
	// typedef struct tagNMDATETIMEFORMATQUERYA
	// {
	//     NMHDR nmhdr;
	//     LPCSTR pszFormat;  // format substring
	//     SIZE szMax;        // max bounding rectangle app will use for this format string
	// } NMDATETIMEFORMATQUERYA, *LPNMDATETIMEFORMATQUERYA;

	// typedef struct tagNMDATETIMEFORMATQUERYW
	// {
	//     NMHDR nmhdr;
	//     LPCWSTR pszFormat; // format substring
	//     SIZE szMax;        // max bounding rectangle app will use for this format string
	// } NMDATETIMEFORMATQUERYW, *LPNMDATETIMEFORMATQUERYW;

	//// #ifdef UNICODE
	//DTN_FORMATQUERY = DTN_FORMATQUERYW
	//NMDATETIMEFORMATQUERY = NMDATETIMEFORMATQUERYW
	//LPNMDATETIMEFORMATQUERY = LPNMDATETIMEFORMATQUERYW
	//// #else
	//DTN_FORMATQUERY = DTN_FORMATQUERYA
	//NMDATETIMEFORMATQUERY = NMDATETIMEFORMATQUERYA
	//LPNMDATETIMEFORMATQUERY = LPNMDATETIMEFORMATQUERYA
	//// #endif

	DTN_DROPDOWN = (DTN_FIRST2 - 1) // MonthCal has dropped down, -754
	DTN_CLOSEUP  = (DTN_FIRST2)     // MonthCal is popping up, -753

	GDTR_MIN = 0
	GDTR_MAX = 0

	// #define GDT_ERROR    -1
	GDT_VALID = 0
	GDT_NONE  = 1

	// #endif // _WIN32
	// #endif // NODATETIMEPICK

	// #ifndef NOIPADDRESS

	// ///////////////////////////////////////////////
	// //    IP Address edit control

	// // Messages sent to IPAddress controls

	IPM_CLEARADDRESS = (WM_USER + 100) // no parameters
	IPM_SETADDRESS   = (WM_USER + 101) // lparam = TCP/IP address
	IPM_GETADDRESS   = (WM_USER + 102) // lresult = # of non black fields.  lparam = LPDWORD for TCP/IP address
	IPM_SETRANGE     = (WM_USER + 103) // wparam = field, lparam = range
	IPM_SETFOCUS     = (WM_USER + 104) // wparam = field
	IPM_ISBLANK      = (WM_USER + 105) // no parameters

	WC_IPADDRESSW = "SysIPAddress32"
	// #define WC_IPADDRESSA           "SysIPAddress32"

	// #ifdef UNICODE
	WC_IPADDRESS = WC_IPADDRESSW
	// #else
	//WC_IPADDRESS = WC_IPADDRESSA
	// #endif

	IPN_FIELDCHANGED = (IPN_FIRST - 0)
	// typedef struct tagNMIPADDRESS
	// {
	//         NMHDR hdr;
	//         int iField;
	//         int iValue;
	// } NMIPADDRESS, *LPNMIPADDRESS;

	// // The following is a useful macro for passing the range values in the
	// // IPM_SETRANGE message.

	// #define MAKEIPRANGE(low, high)    ((LPARAM)(WORD)(((BYTE)(high) << 8) + (BYTE)(low)))

	// // And this is a useful macro for making the IP Address to be passed
	// // as a LPARAM.

	// #define MAKEIPADDRESS(b1,b2,b3,b4)  ((LPARAM)(((DWORD)(b1)<<24)+((DWORD)(b2)<<16)+((DWORD)(b3)<<8)+((DWORD)(b4))))

	// // Get individual number
	// #define FIRST_IPADDRESS(x)  (((x) >> 24) & 0xff)
	// #define SECOND_IPADDRESS(x) (((x) >> 16) & 0xff)
	// #define THIRD_IPADDRESS(x)  (((x) >> 8) & 0xff)
	// #define FOURTH_IPADDRESS(x) ((x) & 0xff)

	// #endif // NOIPADDRESS

	// //---------------------------------------------------------------------------------------
	// //---------------------------------------------------------------------------------------
	//  ====================== Pager Control =============================
	// //---------------------------------------------------------------------------------------
	// //---------------------------------------------------------------------------------------

	// #ifndef NOPAGESCROLLER

	// //Pager Class Name
	WC_PAGESCROLLERW = "SysPager"
	// #define WC_PAGESCROLLERA           "SysPager"

	// #ifdef UNICODE
	WC_PAGESCROLLER = WC_PAGESCROLLERW
	// #else
	//WC_PAGESCROLLER = WC_PAGESCROLLERA
	// #endif

	// //---------------------------------------------------------------------------------------
	// // Pager Control Styles
	// //---------------------------------------------------------------------------------------
	// // begin_r_commctrl

	PGS_VERT       = 0
	PGS_HORZ       = 0
	PGS_AUTOSCROLL = 0
	PGS_DRAGNDROP  = 0

	// // end_r_commctrl

	// //---------------------------------------------------------------------------------------
	// // Pager Button State
	// //---------------------------------------------------------------------------------------
	// //The scroll can be in one of the following control State
	PGF_INVISIBLE = 0 // Scroll button is not visible
	PGF_NORMAL    = 1 // Scroll button is in normal state
	PGF_GRAYED    = 2 // Scroll button is in grayed state
	PGF_DEPRESSED = 4 // Scroll button is in depressed state
	PGF_HOT       = 8 // Scroll button is in hot state

	// // The following identifiers specifies the button control
	PGB_TOPORLEFT     = 0
	PGB_BOTTOMORRIGHT = 1

	// //---------------------------------------------------------------------------------------
	// // Pager Control  Messages
	// //---------------------------------------------------------------------------------------
	PGM_SETCHILD = (PGM_FIRST + 1) // lParam == hwnd
	// #define Pager_SetChild(hwnd, hwndChild) \
	//         (void)SNDMSG((hwnd), PGM_SETCHILD, 0, (LPARAM)(hwndChild))

	PGM_RECALCSIZE = (PGM_FIRST + 2)
	// #define Pager_RecalcSize(hwnd) \
	//         (void)SNDMSG((hwnd), PGM_RECALCSIZE, 0, 0)

	PGM_FORWARDMOUSE = (PGM_FIRST + 3)
	// #define Pager_ForwardMouse(hwnd, bForward) \
	//         (void)SNDMSG((hwnd), PGM_FORWARDMOUSE, (WPARAM)(bForward), 0)

	PGM_SETBKCOLOR = (PGM_FIRST + 4)
	// #define Pager_SetBkColor(hwnd, clr) \
	//         (COLORREF)SNDMSG((hwnd), PGM_SETBKCOLOR, 0, (LPARAM)(clr))

	PGM_GETBKCOLOR = (PGM_FIRST + 5)
	// #define Pager_GetBkColor(hwnd) \
	//         (COLORREF)SNDMSG((hwnd), PGM_GETBKCOLOR, 0, 0)

	PGM_SETBORDER = (PGM_FIRST + 6)
	// #define Pager_SetBorder(hwnd, iBorder) \
	//         (int)SNDMSG((hwnd), PGM_SETBORDER, 0, (LPARAM)(iBorder))

	PGM_GETBORDER = (PGM_FIRST + 7)
	// #define Pager_GetBorder(hwnd) \
	//         (int)SNDMSG((hwnd), PGM_GETBORDER, 0, 0)

	PGM_SETPOS = (PGM_FIRST + 8)
	// #define Pager_SetPos(hwnd, iPos) \
	//         (int)SNDMSG((hwnd), PGM_SETPOS, 0, (LPARAM)(iPos))

	PGM_GETPOS = (PGM_FIRST + 9)
	// #define Pager_GetPos(hwnd) \
	//         (int)SNDMSG((hwnd), PGM_GETPOS, 0, 0)

	PGM_SETBUTTONSIZE = (PGM_FIRST + 10)
	// #define Pager_SetButtonSize(hwnd, iSize) \
	//         (int)SNDMSG((hwnd), PGM_SETBUTTONSIZE, 0, (LPARAM)(iSize))

	PGM_GETBUTTONSIZE = (PGM_FIRST + 11)
	// #define Pager_GetButtonSize(hwnd) \
	//         (int)SNDMSG((hwnd), PGM_GETBUTTONSIZE, 0,0)

	PGM_GETBUTTONSTATE = (PGM_FIRST + 12)
	// #define Pager_GetButtonState(hwnd, iButton) \
	//         (DWORD)SNDMSG((hwnd), PGM_GETBUTTONSTATE, 0, (LPARAM)(iButton))

	PGM_GETDROPTARGET = CCM_GETDROPTARGET
	// #define Pager_GetDropTarget(hwnd, ppdt) \
	//         (void)SNDMSG((hwnd), PGM_GETDROPTARGET, 0, (LPARAM)(ppdt))

	PGM_SETSCROLLINFO = (PGM_FIRST + 13)
	// #define Pager_SetScrollInfo(hwnd, cTimeOut, cLinesPer, cPixelsPerLine) \
	//         (void) SNDMSG((hwnd), PGM_SETSCROLLINFO, cTimeOut, MAKELONG(cLinesPer, cPixelsPerLine))

	// //---------------------------------------------------------------------------------------
	// //Pager Control Notification Messages
	// //---------------------------------------------------------------------------------------

	// // PGN_SCROLL Notification Message

	PGN_SCROLL = (PGN_FIRST - 1)

	PGF_SCROLLUP    = 1
	PGF_SCROLLDOWN  = 2
	PGF_SCROLLLEFT  = 4
	PGF_SCROLLRIGHT = 8

	// //Keys down
	PGK_SHIFT   = 1
	PGK_CONTROL = 2
	PGK_MENU    = 4

	// #ifdef _WIN32
	// #include <pshpack1.h>
	// #endif

	// // This structure is sent along with PGN_SCROLL notifications
	// typedef struct {
	//     NMHDR hdr;
	//     WORD fwKeys;            // Specifies which keys are down when this notification is send
	//     RECT rcParent;          // Contains Parent Window Rect
	//     int  iDir;              // Scrolling Direction
	//     int  iXpos;             // Horizontal scroll position
	//     int  iYpos;             // Vertical scroll position
	//     int  iScroll;           // [in/out] Amount to scroll
	// }NMPGSCROLL, *LPNMPGSCROLL;

	// #ifdef _WIN32
	// #include <poppack.h>
	// #endif

	// // PGN_CALCSIZE Notification Message

	PGN_CALCSIZE = (PGN_FIRST - 2)

	PGF_CALCWIDTH  = 1
	PGF_CALCHEIGHT = 2

	// typedef struct {
	//     NMHDR   hdr;
	//     DWORD   dwFlag;
	//     int     iWidth;
	//     int     iHeight;
	// }NMPGCALCSIZE, *LPNMPGCALCSIZE;

	// // PGN_HOTITEMCHANGE Notification Message

	PGN_HOTITEMCHANGE = (PGN_FIRST - 3)

	/*
	    The PGN_HOTITEMCHANGE notification uses these notification
	    flags defined in TOOLBAR:

	   HICF_ENTERING = 0          // idOld is invalid
	   HICF_LEAVING = 0          // idNew is invalid
	*/

	// // Structure for PGN_HOTITEMCHANGE notification
	// //
	// typedef struct tagNMPGHOTITEM
	// {
	//     NMHDR   hdr;
	//     int     idOld;
	//     int     idNew;
	//     DWORD   dwFlags;           // HICF_*
	// } NMPGHOTITEM, * LPNMPGHOTITEM;

	// #endif // NOPAGESCROLLER

	////======================  End Pager Control ==========================================

	// //
	// === Native Font Control ===
	// //
	// #ifndef NONATIVEFONTCTL
	// //NativeFont Class Name
	WC_NATIVEFONTCTLW = "NativeFontCtl"
	// #define WC_NATIVEFONTCTLA           "NativeFontCtl"

	// #ifdef UNICODE
	WC_NATIVEFONTCTL = WC_NATIVEFONTCTLW
	// #else
	//WC_NATIVEFONTCTL = WC_NATIVEFONTCTLA
	// #endif

	// // begin_r_commctrl

	// // style definition
	NFS_EDIT         = 0
	NFS_STATIC       = 0
	NFS_LISTCOMBO    = 0
	NFS_BUTTON       = 0
	NFS_ALL          = 0
	NFS_USEFONTASSOC = 0

	// // end_r_commctrl

	// #endif // NONATIVEFONTCTL
	// === End Native Font Control ===

	// ====================== Button Control =============================

	// #ifndef NOBUTTON

	// #ifdef _WIN32

	// // Button Class Name
	// #define WC_BUTTONA              "Button"
	WC_BUTTONW = "Button"

	// #ifdef UNICODE
	WC_BUTTON = WC_BUTTONW
	// #else
	//WC_BUTTON = WC_BUTTONA
	// #endif

	// #else
	// #define WC_BUTTON               "Button"
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	BUTTON_IMAGELIST_ALIGN_LEFT   = 0
	BUTTON_IMAGELIST_ALIGN_RIGHT  = 1
	BUTTON_IMAGELIST_ALIGN_TOP    = 2
	BUTTON_IMAGELIST_ALIGN_BOTTOM = 3
	BUTTON_IMAGELIST_ALIGN_CENTER = 4 // Doesn't draw text

	// typedef struct
	// {
	//     HIMAGELIST  himl;   // Images: Normal, Hot, Pushed, Disabled. If count is less than 4, we use index 1
	//     RECT        margin; // Margin around icon.
	//     UINT        uAlign;
	// } BUTTON_IMAGELIST, *PBUTTON_IMAGELIST;

	BCM_GETIDEALSIZE = (BCM_FIRST + 0x0001)
	// #define Button_GetIdealSize(hwnd, psize)\
	//     (BOOL)SNDMSG((hwnd), BCM_GETIDEALSIZE, 0, (LPARAM)(psize))

	BCM_SETIMAGELIST = (BCM_FIRST + 0x0002)
	// #define Button_SetImageList(hwnd, pbuttonImagelist)\
	//     (BOOL)SNDMSG((hwnd), BCM_SETIMAGELIST, 0, (LPARAM)(pbuttonImagelist))

	BCM_GETIMAGELIST = (BCM_FIRST + 0x0003)
	// #define Button_GetImageList(hwnd, pbuttonImagelist)\
	//     (BOOL)SNDMSG((hwnd), BCM_GETIMAGELIST, 0, (LPARAM)(pbuttonImagelist))

	BCM_SETTEXTMARGIN = (BCM_FIRST + 0x0004)
	// #define Button_SetTextMargin(hwnd, pmargin)\
	//     (BOOL)SNDMSG((hwnd), BCM_SETTEXTMARGIN, 0, (LPARAM)(pmargin))
	BCM_GETTEXTMARGIN = (BCM_FIRST + 0x0005)
	// #define Button_GetTextMargin(hwnd, pmargin)\
	//     (BOOL)SNDMSG((hwnd), BCM_GETTEXTMARGIN, 0, (LPARAM)(pmargin))

	// typedef struct tagNMBCHOTITEM
	// {
	//     NMHDR   hdr;
	//     DWORD   dwFlags;           // HICF_*
	// } NMBCHOTITEM, * LPNMBCHOTITEM;

	BCN_HOTITEMCHANGE = (BCN_FIRST + 0x0001)

	BST_HOT = 0

	// #endif // (NTDDI_VERSION >= NTDDI_WINXP)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)

	// // BUTTON STATE FLAGS
	BST_DROPDOWNPUSHED = 0

	// // begin_r_commctrl

	// // BUTTON STYLES
	BS_SPLITBUTTON    = 0
	BS_DEFSPLITBUTTON = 0
	BS_COMMANDLINK    = 0
	BS_DEFCOMMANDLINK = 0

	// // SPLIT BUTTON INFO mask flags
	BCSIF_GLYPH = 0
	BCSIF_IMAGE = 0
	BCSIF_STYLE = 0
	BCSIF_SIZE  = 0

	// // SPLIT BUTTON STYLE flags
	BCSS_NOSPLIT   = 0
	BCSS_STRETCH   = 0
	BCSS_ALIGNLEFT = 0
	BCSS_IMAGE     = 0

	// // end_r_commctrl

	// // BUTTON STRUCTURES
	// typedef struct tagBUTTON_SPLITINFO
	// {
	//     UINT        mask;
	//     HIMAGELIST  himlGlyph;         // interpreted as WCHAR if BCSIF_GLYPH is set
	//     UINT        uSplitStyle;
	//     SIZE        size;
	// } BUTTON_SPLITINFO, * PBUTTON_SPLITINFO;

	// // BUTTON MESSAGES
	BCM_SETDROPDOWNSTATE = (BCM_FIRST + 0x0006)
	// #define Button_SetDropDownState(hwnd, fDropDown) \
	//     (BOOL)SNDMSG((hwnd), BCM_SETDROPDOWNSTATE, (WPARAM)(fDropDown), 0)

	BCM_SETSPLITINFO = (BCM_FIRST + 0x0007)
	// #define Button_SetSplitInfo(hwnd, pInfo) \
	//     (BOOL)SNDMSG((hwnd), BCM_SETSPLITINFO, 0, (LPARAM)(pInfo))

	BCM_GETSPLITINFO = (BCM_FIRST + 0x0008)
	// #define Button_GetSplitInfo(hwnd, pInfo) \
	//     (BOOL)SNDMSG((hwnd), BCM_GETSPLITINFO, 0, (LPARAM)(pInfo))

	BCM_SETNOTE = (BCM_FIRST + 0x0009)
	// #define Button_SetNote(hwnd, psz) \
	//     (BOOL)SNDMSG((hwnd), BCM_SETNOTE, 0, (LPARAM)(psz))

	BCM_GETNOTE = (BCM_FIRST + 0x000A)
	// #define Button_GetNote(hwnd, psz, pcc) \
	//     (BOOL)SNDMSG((hwnd), BCM_GETNOTE, (WPARAM)pcc, (LPARAM)psz)

	BCM_GETNOTELENGTH = (BCM_FIRST + 0x000B)
	// #define Button_GetNoteLength(hwnd) \
	//     (LRESULT)SNDMSG((hwnd), BCM_GETNOTELENGTH, 0, 0)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	// // Macro to use on a button or command link to display an elevated icon
	BCM_SETSHIELD = (BCM_FIRST + 0x000C)
	// #define Button_SetElevationRequiredState(hwnd, fRequired) \
	//     (LRESULT)SNDMSG((hwnd), BCM_SETSHIELD, 0, (LPARAM)fRequired)
	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// // Value to pass to BCM_SETIMAGELIST to indicate that no glyph should be
	// // displayed
	//BCCL_NOGLYPH = (HIMAGELIST)(-1)

	// // NOTIFICATION MESSAGES
	// typedef struct tagNMBCDROPDOWN
	// {
	//     NMHDR   hdr;
	//     RECT    rcButton;
	// } NMBCDROPDOWN, * LPNMBCDROPDOWN;

	BCN_DROPDOWN = (BCN_FIRST + 0x0002)

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// #endif // NOBUTTON

	// =====================  End Button Control =========================

	// ====================== Static Control =============================

	// #ifndef NOSTATIC

	// #ifdef _WIN32

	// // Static Class Name
	// #define WC_STATICA              "Static"
	WC_STATICW = "Static"

	// #ifdef UNICODE
	WC_STATIC = WC_STATICW
	// #else
	//WC_STATIC = WC_STATICA
	// #endif

	// #else
	// #define WC_STATIC               "Static"
	// #endif

	// #endif // NOSTATIC

	// =====================  End Static Control =========================

	// ====================== Edit Control =============================

	// #ifndef NOEDIT

	// #ifdef _WIN32

	// // Edit Class Name
	// #define WC_EDITA                "Edit"
	WC_EDITW = "Edit"

	// #ifdef UNICODE
	WC_EDIT = WC_EDITW
	// #else
	//WC_EDIT = WC_EDITA
	// #endif

	// #else
	// #define WC_EDIT                 "Edit"
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	EM_SETCUEBANNER = (ECM_FIRST + 1) // Set the cue banner with the lParm = LPCWSTR
	// #define Edit_SetCueBannerText(hwnd, lpcwText) \
	//         (BOOL)SNDMSG((hwnd), EM_SETCUEBANNER, 0, (LPARAM)(lpcwText))
	// #define Edit_SetCueBannerTextFocused(hwnd, lpcwText, fDrawFocused) \
	//         (BOOL)SNDMSG((hwnd), EM_SETCUEBANNER, (WPARAM)fDrawFocused, (LPARAM)lpcwText)
	EM_GETCUEBANNER = (ECM_FIRST + 2) // Set the cue banner with the lParm = LPCWSTR
	// #define Edit_GetCueBannerText(hwnd, lpwText, cchText) \
	//         (BOOL)SNDMSG((hwnd), EM_GETCUEBANNER, (WPARAM)(lpwText), (LPARAM)(cchText))

	// typedef struct _tagEDITBALLOONTIP
	// {
	//     DWORD   cbStruct;
	//     LPCWSTR pszTitle;
	//     LPCWSTR pszText;
	//     INT     ttiIcon; // From TTI_*
	// } EDITBALLOONTIP, *PEDITBALLOONTIP;
	EM_SHOWBALLOONTIP = (ECM_FIRST + 3) // Show a balloon tip associated to the edit control
	// #define Edit_ShowBalloonTip(hwnd, peditballoontip) \
	//         (BOOL)SNDMSG((hwnd), EM_SHOWBALLOONTIP, 0, (LPARAM)(peditballoontip))
	EM_HIDEBALLOONTIP = (ECM_FIRST + 4) // Hide any balloon tip associated with the edit control
	// #define Edit_HideBalloonTip(hwnd) \
	//         (BOOL)SNDMSG((hwnd), EM_HIDEBALLOONTIP, 0, 0)
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	EM_SETHILITE = (ECM_FIRST + 5)
	// #define Edit_SetHilite(hwndCtl, ichStart, ichEnd)  ((void)SNDMSG((hwndCtl), EM_SETHILITE, (ichStart), (ichEnd)))
	EM_GETHILITE = (ECM_FIRST + 6)
	// #define Edit_GetHilite(hwndCtl)                    ((DWORD)SNDMSG((hwndCtl), EM_GETHILITE, 0L, 0L))

	// #endif

	EM_NOSETFOCUS = (ECM_FIRST + 7)
	// #define Edit_NoSetFocus(hwndCtl)                   ((DWORD)SNDMSG((hwndCtl), EM_NOSETFOCUS, 0L, 0L))

	EM_TAKEFOCUS = (ECM_FIRST + 8)
	// #define Edit_TakeFocus(hwndCtl)                   ((DWORD)SNDMSG((hwndCtl), EM_TAKEFOCUS, 0L, 0L))

	// #endif // NOEDIT

	// =====================  End Edit Control =========================

	// ====================== Listbox Control =============================

	// #ifndef NOLISTBOX

	// #ifdef _WIN32

	// // Listbox Class Name
	// #define WC_LISTBOXA             "ListBox"
	WC_LISTBOXW = "ListBox"

	// #ifdef UNICODE
	WC_LISTBOX = WC_LISTBOXW
	// #else
	//WC_LISTBOX = WC_LISTBOXA
	// #endif

	// #else
	// #define WC_LISTBOX              "ListBox"
	// #endif

	// #endif // NOLISTBOX

	// =====================  End Listbox Control =========================

	// ====================== Combobox Control =============================

	// #ifndef NOCOMBOBOX

	// #ifdef _WIN32

	// // Combobox Class Name
	// #define WC_COMBOBOXA            "ComboBox"
	WC_COMBOBOXW = "ComboBox"

	// #ifdef UNICODE
	WC_COMBOBOX = WC_COMBOBOXW
	// #else
	//WC_COMBOBOX = WC_COMBOBOXA
	// #endif

	// #else
	// #define WC_COMBOBOX             "ComboBox"
	// #endif

	// #endif // NOCOMBOBOX

	// #if (NTDDI_VERSION >= NTDDI_WINXP)

	// // custom combobox control messages
	CB_SETMINVISIBLE = (CBM_FIRST + 1)
	CB_GETMINVISIBLE = (CBM_FIRST + 2)
	CB_SETCUEBANNER  = (CBM_FIRST + 3)
	CB_GETCUEBANNER  = (CBM_FIRST + 4)

	// #define ComboBox_SetMinVisible(hwnd, iMinVisible) \
	//             (BOOL)SNDMSG((hwnd), CB_SETMINVISIBLE, (WPARAM)(iMinVisible), 0)

	// #define ComboBox_GetMinVisible(hwnd) \
	//             (int)SNDMSG((hwnd), CB_GETMINVISIBLE, 0, 0)

	// #define ComboBox_SetCueBannerText(hwnd, lpcwText)   \
	//             (BOOL)SNDMSG((hwnd), CB_SETCUEBANNER, 0, (LPARAM)(lpcwText))

	// #define ComboBox_GetCueBannerText(hwnd, lpwText, cchText) \
	//             (BOOL)SNDMSG((hwnd), CB_GETCUEBANNER, (WPARAM)(lpwText), (LPARAM)(cchText))

	// #endif

	// =====================  End Combobox Control =========================

	// ====================== Scrollbar Control ============================

	// #ifndef NOSCROLLBAR

	// #ifdef _WIN32

	// // Scrollbar Class Name
	// #define WC_SCROLLBARA            "ScrollBar"
	WC_SCROLLBARW = "ScrollBar"

	// #ifdef UNICODE
	WC_SCROLLBAR = WC_SCROLLBARW
	// #else
	//WC_SCROLLBAR = WC_SCROLLBARA
	// #endif

	// #else
	// #define WC_SCROLLBAR             "ScrollBar"
	// #endif

	// #endif // NOSCROLLBAR

	// ===================== End Scrollbar Control =========================

	// ===================== Task Dialog =========================
	// #ifndef NOTASKDIALOG
	// // Task Dialog is only available starting Windows Vista
	// #if (NTDDI_VERSION >= NTDDI_VISTA)

	// #ifdef _WIN32
	// #include <pshpack1.h>
	// #endif

	// typedef HRESULT (CALLBACK *PFTASKDIALOGCALLBACK)(_In_ HWND hwnd, _In_ UINT msg, _In_ WPARAM wParam, _In_ LPARAM lParam, _In_ LONG_PTR lpRefData);

	// enum _TASKDIALOG_FLAGS
	// {
	TDF_ENABLE_HYPERLINKS           = 0x0001
	TDF_USE_HICON_MAIN              = 0x0002
	TDF_USE_HICON_FOOTER            = 0x0004
	TDF_ALLOW_DIALOG_CANCELLATION   = 0x0008
	TDF_USE_COMMAND_LINKS           = 0x0010
	TDF_USE_COMMAND_LINKS_NO_ICON   = 0x0020
	TDF_EXPAND_FOOTER_AREA          = 0x0040
	TDF_EXPANDED_BY_DEFAULT         = 0x0080
	TDF_VERIFICATION_FLAG_CHECKED   = 0x0100
	TDF_SHOW_PROGRESS_BAR           = 0x0200
	TDF_SHOW_MARQUEE_PROGRESS_BAR   = 0x0400
	TDF_CALLBACK_TIMER              = 0x0800
	TDF_POSITION_RELATIVE_TO_WINDOW = 0x1000
	TDF_RTL_LAYOUT                  = 0x2000
	TDF_NO_DEFAULT_RADIO_BUTTON     = 0x4000
	TDF_CAN_BE_MINIMIZED            = 0x8000
	// #if (NTDDI_VERSION >= NTDDI_WIN8)
	TDF_NO_SET_FOREGROUND = 0x00010000 // Don't call SetForegroundWindow() when activating the dialog
	// #endif // (NTDDI_VERSION >= NTDDI_WIN8)
	TDF_SIZE_TO_CONTENT = 0x01000000 // used by ShellMessageBox to emulate MessageBox sizing behavior
	// };
	// typedef int TASKDIALOG_FLAGS;                         // Note: _TASKDIALOG_FLAGS is an int

	// typedef enum _TASKDIALOG_MESSAGES
	// {
	TDM_NAVIGATE_PAGE                       = WM_USER + 101
	TDM_CLICK_BUTTON                        = WM_USER + 102 // wParam = Button ID
	TDM_SET_MARQUEE_PROGRESS_BAR            = WM_USER + 103 // wParam = 0 (nonMarque) wParam != 0 (Marquee)
	TDM_SET_PROGRESS_BAR_STATE              = WM_USER + 104 // wParam = new progress state
	TDM_SET_PROGRESS_BAR_RANGE              = WM_USER + 105 // lParam = MAKELPARAM(nMinRange, nMaxRange)
	TDM_SET_PROGRESS_BAR_POS                = WM_USER + 106 // wParam = new position
	TDM_SET_PROGRESS_BAR_MARQUEE            = WM_USER + 107 // wParam = 0 (stop marquee), wParam != 0 (start marquee), lparam = speed (milliseconds between repaints)
	TDM_SET_ELEMENT_TEXT                    = WM_USER + 108 // wParam = element (TASKDIALOG_ELEMENTS), lParam = new element text (LPCWSTR)
	TDM_CLICK_RADIO_BUTTON                  = WM_USER + 110 // wParam = Radio Button ID
	TDM_ENABLE_BUTTON                       = WM_USER + 111 // lParam = 0 (disable), lParam != 0 (enable), wParam = Button ID
	TDM_ENABLE_RADIO_BUTTON                 = WM_USER + 112 // lParam = 0 (disable), lParam != 0 (enable), wParam = Radio Button ID
	TDM_CLICK_VERIFICATION                  = WM_USER + 113 // wParam = 0 (unchecked), 1 (checked), lParam = 1 (set key focus)
	TDM_UPDATE_ELEMENT_TEXT                 = WM_USER + 114 // wParam = element (TASKDIALOG_ELEMENTS), lParam = new element text (LPCWSTR)
	TDM_SET_BUTTON_ELEVATION_REQUIRED_STATE = WM_USER + 115 // wParam = Button ID, lParam = 0 (elevation not required), lParam != 0 (elevation required)
	TDM_UPDATE_ICON                         = WM_USER + 116 // wParam = icon element (TASKDIALOG_ICON_ELEMENTS), lParam = new icon (hIcon if TDF_USE_HICON_* was set, PCWSTR otherwise)
	// } TASKDIALOG_MESSAGES;

	// typedef enum _TASKDIALOG_NOTIFICATIONS
	// {
	TDN_CREATED                = 0
	TDN_NAVIGATED              = 1
	TDN_BUTTON_CLICKED         = 2 // wParam = Button ID
	TDN_HYPERLINK_CLICKED      = 3 // lParam = (LPCWSTR)pszHREF
	TDN_TIMER                  = 4 // wParam = Milliseconds since dialog created or timer reset
	TDN_DESTROYED              = 5
	TDN_RADIO_BUTTON_CLICKED   = 6 // wParam = Radio Button ID
	TDN_DIALOG_CONSTRUCTED     = 7
	TDN_VERIFICATION_CLICKED   = 8 // wParam = 1 if checkbox checked, 0 if not, lParam is unused and always 0
	TDN_HELP                   = 9
	TDN_EXPANDO_BUTTON_CLICKED = 10 // wParam = 0 (dialog is now collapsed), wParam != 0 (dialog is now expanded)
	// } TASKDIALOG_NOTIFICATIONS;

	// typedef struct _TASKDIALOG_BUTTON
	// {
	//     int     nButtonID;
	//     PCWSTR  pszButtonText;
	// } TASKDIALOG_BUTTON;

	// typedef enum _TASKDIALOG_ELEMENTS
	// {
	//     TDE_CONTENT,
	//     TDE_EXPANDED_INFORMATION,
	//     TDE_FOOTER,
	//     TDE_MAIN_INSTRUCTION
	// } TASKDIALOG_ELEMENTS;

	// typedef enum _TASKDIALOG_ICON_ELEMENTS
	// {
	//     TDIE_ICON_MAIN,
	//     TDIE_ICON_FOOTER
	// } TASKDIALOG_ICON_ELEMENTS;

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)

	// #if (NTDDI_VERSION >= NTDDI_VISTA)
	// enum _TASKDIALOG_COMMON_BUTTON_FLAGS
	// {
	TDCBF_OK_BUTTON     = 0x0001 // selected control return value IDOK
	TDCBF_YES_BUTTON    = 0x0002 // selected control return value IDYES
	TDCBF_NO_BUTTON     = 0x0004 // selected control return value IDNO
	TDCBF_CANCEL_BUTTON = 0x0008 // selected control return value IDCANCEL
	TDCBF_RETRY_BUTTON  = 0x0010 // selected control return value IDRETRY
	TDCBF_CLOSE_BUTTON  = 0x0020 // selected control return value IDCLOSE
	// };
	// typedef int TASKDIALOG_COMMON_BUTTON_FLAGS;           // Note: _TASKDIALOG_COMMON_BUTTON_FLAGS is an int

	// typedef struct _TASKDIALOGCONFIG
	// {
	//     UINT        cbSize;
	//     HWND        hwndParent;                             // incorrectly named, this is the owner window, not a parent.
	//     HINSTANCE   hInstance;                              // used for MAKEINTRESOURCE() strings
	//     TASKDIALOG_FLAGS                dwFlags;            // TASKDIALOG_FLAGS (TDF_XXX) flags
	//     TASKDIALOG_COMMON_BUTTON_FLAGS  dwCommonButtons;    // TASKDIALOG_COMMON_BUTTON (TDCBF_XXX) flags
	//     PCWSTR      pszWindowTitle;                         // string or MAKEINTRESOURCE()
	//     union
	//     {
	//         HICON   hMainIcon;
	//         PCWSTR  pszMainIcon;
	//     } DUMMYUNIONNAME;
	//     PCWSTR      pszMainInstruction;
	//     PCWSTR      pszContent;
	//     UINT        cButtons;
	//     const TASKDIALOG_BUTTON  *pButtons;
	//     int         nDefaultButton;
	//     UINT        cRadioButtons;
	//     const TASKDIALOG_BUTTON  *pRadioButtons;
	//     int         nDefaultRadioButton;
	//     PCWSTR      pszVerificationText;
	//     PCWSTR      pszExpandedInformation;
	//     PCWSTR      pszExpandedControlText;
	//     PCWSTR      pszCollapsedControlText;
	//     union
	//     {
	//         HICON   hFooterIcon;
	//         PCWSTR  pszFooterIcon;
	//     } DUMMYUNIONNAME2;
	//     PCWSTR      pszFooter;
	//     PFTASKDIALOGCALLBACK pfCallback;
	//     LONG_PTR    lpCallbackData;
	//     UINT        cxWidth;                                // width of the Task Dialog's client area in DLU's. If 0, Task Dialog will calculate the ideal width.
	// } TASKDIALOGCONFIG;

	// WINCOMMCTRLAPI HRESULT WINAPI TaskDialogIndirect(_In_ const TASKDIALOGCONFIG *pTaskConfig, _Out_opt_ int *pnButton, _Out_opt_ int *pnRadioButton, _Out_opt_ BOOL *pfVerificationFlagChecked);
	// WINCOMMCTRLAPI HRESULT WINAPI TaskDialog(_In_opt_ HWND hwndOwner, _In_opt_ HINSTANCE hInstance, _In_opt_ PCWSTR pszWindowTitle, _In_opt_ PCWSTR pszMainInstruction, _In_opt_ PCWSTR pszContent, TASKDIALOG_COMMON_BUTTON_FLAGS dwCommonButtons, _In_opt_ PCWSTR pszIcon, _Out_opt_ int *pnButton);

	// #ifdef _WIN32
	// #include <poppack.h>
	// #endif

	// #endif // (NTDDI_VERSION >= NTDDI_VISTA)
	// #endif // NOTASKDIALOG

	// ==================== End TaskDialog =======================

	// //
	// === MUI APIs ===
	// //
	// #ifndef NOMUI
	// void WINAPI InitMUILanguage(LANGID uiLang);

	// LANGID WINAPI GetMUILanguage(void);
	// #endif  // NOMUI

	// #include <dpa_dsa.h>

	// #ifdef _WIN32
	//====== TrackMouseEvent  =====================================================

	// #ifndef NOTRACKMOUSEEVENT

	// //
	// // If the messages for TrackMouseEvent have not been defined then define them
	// // now.
	// //
	// #ifndef WM_MOUSEHOVER
	WM_MOUSEHOVER = 0
	WM_MOUSELEAVE = 0
	// #endif

	// //
	// // If the TRACKMOUSEEVENT structure and associated flags havent been declared
	// // then declare them now.
	// //
	// #ifndef TME_HOVER

	TME_HOVER = 0
	TME_LEAVE = 0
	// #if (WINVER >= 0x0500)
	TME_NONCLIENT = 0
	// #endif /* WINVER >= 0x0500 */
	TME_QUERY  = 0
	TME_CANCEL = 0

	HOVER_DEFAULT = 0

	// typedef struct tagTRACKMOUSEEVENT {
	//     DWORD cbSize;
	//     DWORD dwFlags;
	//     HWND  hwndTrack;
	//     DWORD dwHoverTime;
	// } TRACKMOUSEEVENT, *LPTRACKMOUSEEVENT;

	// #endif // !TME_HOVER

	// //
	// // Declare _TrackMouseEvent.  This API tries to use the window manager's
	// // implementation of TrackMouseEvent if it is present, otherwise it emulates.
	// //
	// WINCOMMCTRLAPI
	// BOOL
	// WINAPI
	// _TrackMouseEvent(
	//     _Inout_ LPTRACKMOUSEEVENT lpEventTrack);

	// #endif // !NOTRACKMOUSEEVENT

	//====== Flat Scrollbar APIs=========================================
	// #ifndef NOFLATSBAPIS

	WSB_PROP_CYVSCROLL = 0
	WSB_PROP_CXHSCROLL = 0
	WSB_PROP_CYHSCROLL = 0
	WSB_PROP_CXVSCROLL = 0
	WSB_PROP_CXHTHUMB  = 0
	WSB_PROP_CYVTHUMB  = 0
	WSB_PROP_VBKGCOLOR = 0
	WSB_PROP_HBKGCOLOR = 0
	WSB_PROP_VSTYLE    = 0
	WSB_PROP_HSTYLE    = 0
	WSB_PROP_WINSTYLE  = 0
	WSB_PROP_PALETTE   = 0
	WSB_PROP_MASK      = 0

	FSB_FLAT_MODE    = 2
	FSB_ENCARTA_MODE = 1
	FSB_REGULAR_MODE = 0

	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_EnableScrollBar(HWND, int, UINT);
	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_ShowScrollBar(HWND, int code, BOOL);

	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_GetScrollRange(HWND, int code, LPINT, LPINT);
	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_GetScrollInfo(HWND, int code, LPSCROLLINFO);

	// WINCOMMCTRLAPI int WINAPI FlatSB_GetScrollPos(HWND, int code);

	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_GetScrollProp(HWND, int propIndex, LPINT);
	// #ifdef _WIN64
	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_GetScrollPropPtr(HWND, int propIndex, PINT_PTR);
	// #else
	//FlatSB_GetScrollPropPtr = FlatSB_GetScrollProp
	// #endif

	// WINCOMMCTRLAPI int WINAPI FlatSB_SetScrollPos(HWND, int code, int pos, BOOL fRedraw);

	// WINCOMMCTRLAPI int WINAPI FlatSB_SetScrollInfo(HWND, int code, LPSCROLLINFO psi, BOOL fRedraw);

	// WINCOMMCTRLAPI int WINAPI FlatSB_SetScrollRange(HWND, int code, int min, int max, BOOL fRedraw);
	// WINCOMMCTRLAPI BOOL WINAPI FlatSB_SetScrollProp(HWND, UINT index, INT_PTR newValue, BOOL);
	//FlatSB_SetScrollPropPtr = FlatSB_SetScrollProp

	// WINCOMMCTRLAPI BOOL WINAPI InitializeFlatSB(HWND);
	// WINCOMMCTRLAPI HRESULT WINAPI UninitializeFlatSB(HWND);

	// #endif  //  NOFLATSBAPIS

	// #endif /* _WIN32 */

	// #if (NTDDI_VERSION >= NTDDI_WINXP)
	// //
	// // subclassing stuff
	// //
	// typedef LRESULT (CALLBACK *SUBCLASSPROC)(HWND hWnd, UINT uMsg, WPARAM wParam,
	//     LPARAM lParam, UINT_PTR uIdSubclass, DWORD_PTR dwRefData);

	// _Success_(return) BOOL WINAPI SetWindowSubclass(_In_ HWND hWnd, _In_ SUBCLASSPROC pfnSubclass, _In_ UINT_PTR uIdSubclass,
	//     _In_ DWORD_PTR dwRefData);
	// _Success_(return) BOOL WINAPI GetWindowSubclass(_In_ HWND hWnd, _In_ SUBCLASSPROC pfnSubclass, _In_ UINT_PTR uIdSubclass,
	//     _Out_opt_ DWORD_PTR *pdwRefData);
	// _Success_(return) BOOL WINAPI RemoveWindowSubclass(_In_ HWND hWnd, _In_ SUBCLASSPROC pfnSubclass,
	//     _In_ UINT_PTR uIdSubclass);

	// LRESULT WINAPI DefSubclassProc(HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam);
	// #endif

	// #if (NTDDI_VERSION >= NTDDI_VISTA)

	// enum _LI_METRIC
	// {
	//    LIM_SMALL, // corresponds to SM_CXSMICON/SM_CYSMICON
	//    LIM_LARGE, // corresponds to SM_CXICON/SM_CYICON
	// };

	// WINCOMMCTRLAPI HRESULT WINAPI LoadIconMetric(HINSTANCE hinst, PCWSTR pszName, int lims, _Out_ HICON *phico);
	// WINCOMMCTRLAPI HRESULT WINAPI LoadIconWithScaleDown(HINSTANCE hinst, PCWSTR pszName, int cx, int cy, _Out_ HICON *phico);

	// #endif // NTDDI_VISTA

	// #if (NTDDI_VERSION >= NTDDI_WINXP)

	// int WINAPI DrawShadowText(_In_ HDC hdc, _In_reads_(cch) LPCWSTR pszText, _In_ UINT cch, _In_ RECT* prc, _In_ DWORD dwFlags, _In_ COLORREF crText, _In_ COLORREF crShadow,
	//     _In_ int ixOffset, _In_ int iyOffset);
	// #endif

	// #if !defined(RC_INVOKED) /* RC complains about long symbols in #ifs */
	// #if defined(ISOLATION_AWARE_ENABLED) && (ISOLATION_AWARE_ENABLED != 0)
	// #include "commctrl.inl"
	// #endif /* ISOLATION_AWARE_ENABLED */
	// #endif /* RC */

	// #ifdef __cplusplus
	// }
	// #endif

	// #endif

	// #if defined(_MSC_VER) && (_MSC_VER >= 1200)
	// #pragma warning(pop)
	// #endif

	// #endif /* WINAPI_FAMILY_PARTITION(WINAPI_PARTITION_DESKTOP) */
	// #pragma endregion

	// #endif  /* _INC_COMMCTRL */

)

var TD_WARNING_ICON = -1
var TD_ERROR_ICON = -2
var TD_INFORMATION_ICON = -3
var TD_SHIELD_ICON = -4


type PBRANGE struct
{
   Low,High int
}