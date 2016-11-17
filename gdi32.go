// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libgdi32 uintptr

	// Functions
	abortDoc                     uintptr
	abortPath                    uintptr
	addFontMemResourceEx         uintptr
	addFontResourceEx            uintptr
	addFontResource              uintptr
	angleArc                     uintptr
	animatePalette               uintptr
	arc                          uintptr
	arcTo                        uintptr
	bRUSHOBJ_hGetColorTransform  uintptr
	bRUSHOBJ_pvAllocRbrush       uintptr
	bRUSHOBJ_pvGetRbrush         uintptr
	bRUSHOBJ_ulGetBrushColor     uintptr
	beginPath                    uintptr
	bitBlt                       uintptr
	cLIPOBJ_bEnum                uintptr
	cLIPOBJ_cEnumStart           uintptr
	cancelDC                     uintptr
	checkColorsInGamut           uintptr
	choosePixelFormat            uintptr
	chord                        uintptr
	closeEnhMetaFile             uintptr
	closeFigure                  uintptr
	closeMetaFile                uintptr
	colorCorrectPalette          uintptr
	colorMatchToTarget           uintptr
	combineRgn                   uintptr
	combineTransform             uintptr
	copyEnhMetaFile              uintptr
	copyMetaFile                 uintptr
	createBitmap                 uintptr
	createBitmapIndirect         uintptr
	createBrushIndirect          uintptr
	createColorSpace             uintptr
	createCompatibleBitmap       uintptr
	createCompatibleDC           uintptr
	createDC                     uintptr
	createDIBPatternBrush        uintptr
	createDIBPatternBrushPt      uintptr
	createDIBSection             uintptr
	createDIBitmap               uintptr
	createDiscardableBitmap      uintptr
	createEllipticRgn            uintptr
	createEllipticRgnIndirect    uintptr
	createEnhMetaFile            uintptr
	createFontIndirectEx         uintptr
	createFontIndirect           uintptr
	createFont                   uintptr
	createHalftonePalette        uintptr
	createHatchBrush             uintptr
	createIC                     uintptr
	createMetaFile               uintptr
	createPalette                uintptr
	createPatternBrush           uintptr
	createPen                    uintptr
	createPenIndirect            uintptr
	createPolyPolygonRgn         uintptr
	createPolygonRgn             uintptr
	createRectRgn                uintptr
	createRectRgnIndirect        uintptr
	createRoundRectRgn           uintptr
	createScalableFontResource   uintptr
	createSolidBrush             uintptr
	dPtoLP                       uintptr
	deleteColorSpace             uintptr
	deleteDC                     uintptr
	deleteEnhMetaFile            uintptr
	deleteMetaFile               uintptr
	deleteObject                 uintptr
	describePixelFormat          uintptr
	drawEscape                   uintptr
	ellipse                      uintptr
	endDoc                       uintptr
	endPage                      uintptr
	endPath                      uintptr
	engAcquireSemaphore          uintptr
	engAlphaBlend                uintptr
	engAssociateSurface          uintptr
	engBitBlt                    uintptr
	engCheckAbort                uintptr
	engCopyBits                  uintptr
	engCreateBitmap              uintptr
	engCreateDeviceBitmap        uintptr
	engCreateDeviceSurface       uintptr
	engCreatePalette             uintptr
	engCreateSemaphore           uintptr
	engDeleteClip                uintptr
	engDeletePalette             uintptr
	engDeletePath                uintptr
	engDeleteSemaphore           uintptr
	engDeleteSurface             uintptr
	engEraseSurface              uintptr
	engFillPath                  uintptr
	engFindResource              uintptr
	engFreeModule                uintptr
	engGetCurrentCodePage        uintptr
	engGetDriverName             uintptr
	engGetPrinterDataFileName    uintptr
	engGradientFill              uintptr
	engLineTo                    uintptr
	engLoadModule                uintptr
	engMarkBandingSurface        uintptr
	engMultiByteToUnicodeN       uintptr
	engMultiByteToWideChar       uintptr
	engPaint                     uintptr
	engPlgBlt                    uintptr
	engQueryLocalTime            uintptr
	engReleaseSemaphore          uintptr
	engStretchBlt                uintptr
	engStretchBltROP             uintptr
	engStrokeAndFillPath         uintptr
	engStrokePath                uintptr
	engTextOut                   uintptr
	engTransparentBlt            uintptr
	engUnicodeToMultiByteN       uintptr
	engUnlockSurface             uintptr
	engWideCharToMultiByte       uintptr
	enumEnhMetaFile              uintptr
	enumFontFamiliesEx           uintptr
	enumFontFamilies             uintptr
	enumFonts                    uintptr
	enumICMProfiles              uintptr
	enumMetaFile                 uintptr
	enumObjects                  uintptr
	equalRgn                     uintptr
	escape                       uintptr
	excludeClipRect              uintptr
	extCreatePen                 uintptr
	extCreateRegion              uintptr
	extEscape                    uintptr
	extFloodFill                 uintptr
	extSelectClipRgn             uintptr
	extTextOut                   uintptr
	fONTOBJ_cGetAllGlyphHandles  uintptr
	fONTOBJ_cGetGlyphs           uintptr
	fONTOBJ_pQueryGlyphAttrs     uintptr
	fONTOBJ_pvTrueTypeFontFile   uintptr
	fONTOBJ_vGetInfo             uintptr
	fillPath                     uintptr
	fillRgn                      uintptr
	fixBrushOrgEx                uintptr
	flattenPath                  uintptr
	floodFill                    uintptr
	frameRgn                     uintptr
	gdiAlphaBlend                uintptr
	gdiComment                   uintptr
	gdiFlush                     uintptr
	gdiGetBatchLimit             uintptr
	gdiGradientFill              uintptr
	gdiSetBatchLimit             uintptr
	gdiTransparentBlt            uintptr
	getArcDirection              uintptr
	getAspectRatioFilterEx       uintptr
	getBitmapBits                uintptr
	getBitmapDimensionEx         uintptr
	getBkColor                   uintptr
	getBkMode                    uintptr
	getBoundsRect                uintptr
	getBrushOrgEx                uintptr
	getCharABCWidthsFloat        uintptr
	getCharABCWidthsI            uintptr
	getCharABCWidths             uintptr
	getCharWidth32               uintptr
	getCharWidthFloat            uintptr
	getCharWidthI                uintptr
	getCharWidth                 uintptr
	getCharacterPlacement        uintptr
	getClipBox                   uintptr
	getClipRgn                   uintptr
	getColorAdjustment           uintptr
	getColorSpace                uintptr
	getCurrentObject             uintptr
	getCurrentPositionEx         uintptr
	getDCBrushColor              uintptr
	getDCOrgEx                   uintptr
	getDCPenColor                uintptr
	getDIBColorTable             uintptr
	getDIBits                    uintptr
	getDeviceCaps                uintptr
	getDeviceGammaRamp           uintptr
	getEnhMetaFileBits           uintptr
	getEnhMetaFileDescription    uintptr
	getEnhMetaFileHeader         uintptr
	getEnhMetaFilePaletteEntries uintptr
	getEnhMetaFilePixelFormat    uintptr
	getEnhMetaFile               uintptr
	getFontData                  uintptr
	getFontLanguageInfo          uintptr
	getFontUnicodeRanges         uintptr
	getGlyphIndices              uintptr
	getGlyphOutline              uintptr
	getGraphicsMode              uintptr
	getICMProfile                uintptr
	getKerningPairs              uintptr
	getLayout                    uintptr
	getLogColorSpace             uintptr
	getMapMode                   uintptr
	getMetaFileBitsEx            uintptr
	getMetaFile                  uintptr
	getMetaRgn                   uintptr
	getMiterLimit                uintptr
	getNearestColor              uintptr
	getNearestPaletteIndex       uintptr
	getObjectType                uintptr
	getObject                    uintptr
	getOutlineTextMetrics        uintptr
	getPaletteEntries            uintptr
	getPath                      uintptr
	getPixel                     uintptr
	getPixelFormat               uintptr
	getPolyFillMode              uintptr
	getROP2                      uintptr
	getRandomRgn                 uintptr
	getRasterizerCaps            uintptr
	getRegionData                uintptr
	getRgnBox                    uintptr
	getStockObject               uintptr
	getStretchBltMode            uintptr
	getSystemPaletteEntries      uintptr
	getSystemPaletteUse          uintptr
	getTextAlign                 uintptr
	getTextCharacterExtra        uintptr
	getTextCharset               uintptr
	getTextCharsetInfo           uintptr
	getTextColor                 uintptr
	getTextExtentExPointI        uintptr
	getTextExtentExPoint         uintptr
	getTextExtentPoint32         uintptr
	getTextExtentPointI          uintptr
	getTextExtentPoint           uintptr
	getTextFace                  uintptr
	getTextMetrics               uintptr
	getViewportExtEx             uintptr
	getViewportOrgEx             uintptr
	getWinMetaFileBits           uintptr
	getWindowExtEx               uintptr
	getWindowOrgEx               uintptr
	getWorldTransform            uintptr
	hT_Get8BPPFormatPalette      uintptr
	hT_Get8BPPMaskPalette        uintptr
	intersectClipRect            uintptr
	invertRgn                    uintptr
	lPtoDP                       uintptr
	lineDDA                      uintptr
	lineTo                       uintptr
	maskBlt                      uintptr
	modifyWorldTransform         uintptr
	moveToEx                     uintptr
	offsetClipRgn                uintptr
	offsetRgn                    uintptr
	offsetViewportOrgEx          uintptr
	offsetWindowOrgEx            uintptr
	pATHOBJ_bEnum                uintptr
	pATHOBJ_bEnumClipLines       uintptr
	pATHOBJ_vEnumStart           uintptr
	pATHOBJ_vEnumStartClipLines  uintptr
	pATHOBJ_vGetBounds           uintptr
	paintRgn                     uintptr
	patBlt                       uintptr
	pathToRegion                 uintptr
	pie                          uintptr
	playEnhMetaFile              uintptr
	playEnhMetaFileRecord        uintptr
	playMetaFile                 uintptr
	playMetaFileRecord           uintptr
	plgBlt                       uintptr
	polyBezier                   uintptr
	polyBezierTo                 uintptr
	polyDraw                     uintptr
	polyPolygon                  uintptr
	polyPolyline                 uintptr
	polyTextOut                  uintptr
	polygon                      uintptr
	polyline                     uintptr
	polylineTo                   uintptr
	ptInRegion                   uintptr
	ptVisible                    uintptr
	realizePalette               uintptr
	rectInRegion                 uintptr
	rectVisible                  uintptr
	rectangle                    uintptr
	removeFontMemResourceEx      uintptr
	removeFontResourceEx         uintptr
	removeFontResource           uintptr
	resetDC                      uintptr
	resizePalette                uintptr
	restoreDC                    uintptr
	roundRect                    uintptr
	sTROBJ_bEnum                 uintptr
	sTROBJ_bEnumPositionsOnly    uintptr
	sTROBJ_bGetAdvanceWidths     uintptr
	sTROBJ_dwGetCodePage         uintptr
	sTROBJ_vEnumStart            uintptr
	saveDC                       uintptr
	scaleViewportExtEx           uintptr
	scaleWindowExtEx             uintptr
	selectClipPath               uintptr
	selectClipRgn                uintptr
	selectObject                 uintptr
	selectPalette                uintptr
	setAbortProc                 uintptr
	setArcDirection              uintptr
	setBitmapBits                uintptr
	setBitmapDimensionEx         uintptr
	setBkColor                   uintptr
	setBkMode                    uintptr
	setBoundsRect                uintptr
	setBrushOrgEx                uintptr
	setColorAdjustment           uintptr
	setColorSpace                uintptr
	setDCBrushColor              uintptr
	setDCPenColor                uintptr
	setDIBColorTable             uintptr
	setDIBits                    uintptr
	setDIBitsToDevice            uintptr
	setDeviceGammaRamp           uintptr
	setEnhMetaFileBits           uintptr
	setGraphicsMode              uintptr
	setICMMode                   uintptr
	setICMProfile                uintptr
	setLayout                    uintptr
	setMapMode                   uintptr
	setMapperFlags               uintptr
	setMetaFileBitsEx            uintptr
	setMetaRgn                   uintptr
	setMiterLimit                uintptr
	setPaletteEntries            uintptr
	setPixel                     uintptr
	setPixelFormat               uintptr
	setPixelV                    uintptr
	setPolyFillMode              uintptr
	setROP2                      uintptr
	setRectRgn                   uintptr
	setStretchBltMode            uintptr
	setSystemPaletteUse          uintptr
	setTextAlign                 uintptr
	setTextCharacterExtra        uintptr
	setTextColor                 uintptr
	setTextJustification         uintptr
	setViewportExtEx             uintptr
	setViewportOrgEx             uintptr
	setWinMetaFileBits           uintptr
	setWindowExtEx               uintptr
	setWindowOrgEx               uintptr
	setWorldTransform            uintptr
	startDoc                     uintptr
	startPage                    uintptr
	stretchBlt                   uintptr
	stretchDIBits                uintptr
	strokeAndFillPath            uintptr
	strokePath                   uintptr
	swapBuffers                  uintptr
	textOut                      uintptr
	translateCharsetInfo         uintptr
	unrealizeObject              uintptr
	updateColors                 uintptr
	updateICMRegKey              uintptr
	widenPath                    uintptr
	xFORMOBJ_bApplyXform         uintptr
	xFORMOBJ_iGetXform           uintptr
	xLATEOBJ_cGetPalette         uintptr
	xLATEOBJ_hGetColorTransform  uintptr
	xLATEOBJ_iXlate              uintptr
	enableEUDC                   uintptr
	fontIsLinked                 uintptr
	gdiDescribePixelFormat       uintptr
	gdiDrawStream                uintptr
	gdiGetCharDimensions         uintptr
	gdiGetCodePage               uintptr
	gdiGetSpoolMessage           uintptr
	gdiInitSpool                 uintptr
	gdiInitializeLanguagePack    uintptr
	gdiIsMetaFileDC              uintptr
	gdiIsMetaPrintDC             uintptr
	gdiIsPlayMetafileDC          uintptr
	gdiRealizationInfo           uintptr
	gdiSetPixelFormat            uintptr
	gdiSwapBuffers               uintptr
	getFontResourceInfoW         uintptr
	getRelAbs                    uintptr
	getTransform                 uintptr
	mirrorRgn                    uintptr
	namedEscape                  uintptr
	setMagicColors               uintptr
	setRelAbs                    uintptr
	setVirtualResolution         uintptr
)

func init() {
	// Library
	libgdi32 = doLoadLibrary("gdi32.dll")

	// Functions
	abortDoc = doGetProcAddress(libgdi32, "AbortDoc")
	abortPath = doGetProcAddress(libgdi32, "AbortPath")
	addFontMemResourceEx = doGetProcAddress(libgdi32, "AddFontMemResourceEx")
	addFontResourceEx = doGetProcAddress(libgdi32, "AddFontResourceExW")
	addFontResource = doGetProcAddress(libgdi32, "AddFontResourceW")
	angleArc = doGetProcAddress(libgdi32, "AngleArc")
	animatePalette = doGetProcAddress(libgdi32, "AnimatePalette")
	arc = doGetProcAddress(libgdi32, "Arc")
	arcTo = doGetProcAddress(libgdi32, "ArcTo")
	bRUSHOBJ_hGetColorTransform = doGetProcAddress(libgdi32, "BRUSHOBJ_hGetColorTransform")
	bRUSHOBJ_pvAllocRbrush = doGetProcAddress(libgdi32, "BRUSHOBJ_pvAllocRbrush")
	bRUSHOBJ_pvGetRbrush = doGetProcAddress(libgdi32, "BRUSHOBJ_pvGetRbrush")
	bRUSHOBJ_ulGetBrushColor = doGetProcAddress(libgdi32, "BRUSHOBJ_ulGetBrushColor")
	beginPath = doGetProcAddress(libgdi32, "BeginPath")
	bitBlt = doGetProcAddress(libgdi32, "BitBlt")
	cLIPOBJ_bEnum = doGetProcAddress(libgdi32, "CLIPOBJ_bEnum")
	cLIPOBJ_cEnumStart = doGetProcAddress(libgdi32, "CLIPOBJ_cEnumStart")
	cancelDC = doGetProcAddress(libgdi32, "CancelDC")
	checkColorsInGamut = doGetProcAddress(libgdi32, "CheckColorsInGamut")
	choosePixelFormat = doGetProcAddress(libgdi32, "ChoosePixelFormat")
	chord = doGetProcAddress(libgdi32, "Chord")
	closeEnhMetaFile = doGetProcAddress(libgdi32, "CloseEnhMetaFile")
	closeFigure = doGetProcAddress(libgdi32, "CloseFigure")
	closeMetaFile = doGetProcAddress(libgdi32, "CloseMetaFile")
	colorCorrectPalette = doGetProcAddress(libgdi32, "ColorCorrectPalette")
	colorMatchToTarget = doGetProcAddress(libgdi32, "ColorMatchToTarget")
	combineRgn = doGetProcAddress(libgdi32, "CombineRgn")
	combineTransform = doGetProcAddress(libgdi32, "CombineTransform")
	copyEnhMetaFile = doGetProcAddress(libgdi32, "CopyEnhMetaFileW")
	copyMetaFile = doGetProcAddress(libgdi32, "CopyMetaFileW")
	createBitmap = doGetProcAddress(libgdi32, "CreateBitmap")
	createBitmapIndirect = doGetProcAddress(libgdi32, "CreateBitmapIndirect")
	createBrushIndirect = doGetProcAddress(libgdi32, "CreateBrushIndirect")
	createColorSpace = doGetProcAddress(libgdi32, "CreateColorSpaceW")
	createCompatibleBitmap = doGetProcAddress(libgdi32, "CreateCompatibleBitmap")
	createCompatibleDC = doGetProcAddress(libgdi32, "CreateCompatibleDC")
	createDC = doGetProcAddress(libgdi32, "CreateDCW")
	createDIBPatternBrush = doGetProcAddress(libgdi32, "CreateDIBPatternBrush")
	createDIBPatternBrushPt = doGetProcAddress(libgdi32, "CreateDIBPatternBrushPt")
	createDIBSection = doGetProcAddress(libgdi32, "CreateDIBSection")
	createDIBitmap = doGetProcAddress(libgdi32, "CreateDIBitmap")
	createDiscardableBitmap = doGetProcAddress(libgdi32, "CreateDiscardableBitmap")
	createEllipticRgn = doGetProcAddress(libgdi32, "CreateEllipticRgn")
	createEllipticRgnIndirect = doGetProcAddress(libgdi32, "CreateEllipticRgnIndirect")
	createEnhMetaFile = doGetProcAddress(libgdi32, "CreateEnhMetaFileW")
	createFontIndirectEx = doGetProcAddress(libgdi32, "CreateFontIndirectExW")
	createFontIndirect = doGetProcAddress(libgdi32, "CreateFontIndirectW")
	createFont = doGetProcAddress(libgdi32, "CreateFontW")
	createHalftonePalette = doGetProcAddress(libgdi32, "CreateHalftonePalette")
	createHatchBrush = doGetProcAddress(libgdi32, "CreateHatchBrush")
	createIC = doGetProcAddress(libgdi32, "CreateICW")
	createMetaFile = doGetProcAddress(libgdi32, "CreateMetaFileW")
	createPalette = doGetProcAddress(libgdi32, "CreatePalette")
	createPatternBrush = doGetProcAddress(libgdi32, "CreatePatternBrush")
	createPen = doGetProcAddress(libgdi32, "CreatePen")
	createPenIndirect = doGetProcAddress(libgdi32, "CreatePenIndirect")
	createPolyPolygonRgn = doGetProcAddress(libgdi32, "CreatePolyPolygonRgn")
	createPolygonRgn = doGetProcAddress(libgdi32, "CreatePolygonRgn")
	createRectRgn = doGetProcAddress(libgdi32, "CreateRectRgn")
	createRectRgnIndirect = doGetProcAddress(libgdi32, "CreateRectRgnIndirect")
	createRoundRectRgn = doGetProcAddress(libgdi32, "CreateRoundRectRgn")
	createScalableFontResource = doGetProcAddress(libgdi32, "CreateScalableFontResourceW")
	createSolidBrush = doGetProcAddress(libgdi32, "CreateSolidBrush")
	dPtoLP = doGetProcAddress(libgdi32, "DPtoLP")
	deleteColorSpace = doGetProcAddress(libgdi32, "DeleteColorSpace")
	deleteDC = doGetProcAddress(libgdi32, "DeleteDC")
	deleteEnhMetaFile = doGetProcAddress(libgdi32, "DeleteEnhMetaFile")
	deleteMetaFile = doGetProcAddress(libgdi32, "DeleteMetaFile")
	deleteObject = doGetProcAddress(libgdi32, "DeleteObject")
	describePixelFormat = doGetProcAddress(libgdi32, "DescribePixelFormat")
	drawEscape = doGetProcAddress(libgdi32, "DrawEscape")
	ellipse = doGetProcAddress(libgdi32, "Ellipse")
	endDoc = doGetProcAddress(libgdi32, "EndDoc")
	endPage = doGetProcAddress(libgdi32, "EndPage")
	endPath = doGetProcAddress(libgdi32, "EndPath")
	engAcquireSemaphore = doGetProcAddress(libgdi32, "EngAcquireSemaphore")
	engAlphaBlend = doGetProcAddress(libgdi32, "EngAlphaBlend")
	engAssociateSurface = doGetProcAddress(libgdi32, "EngAssociateSurface")
	engBitBlt = doGetProcAddress(libgdi32, "EngBitBlt")
	engCheckAbort = doGetProcAddress(libgdi32, "EngCheckAbort")
	engCopyBits = doGetProcAddress(libgdi32, "EngCopyBits")
	engCreateBitmap = doGetProcAddress(libgdi32, "EngCreateBitmap")
	engCreateDeviceBitmap = doGetProcAddress(libgdi32, "EngCreateDeviceBitmap")
	engCreateDeviceSurface = doGetProcAddress(libgdi32, "EngCreateDeviceSurface")
	engCreatePalette = doGetProcAddress(libgdi32, "EngCreatePalette")
	engCreateSemaphore = doGetProcAddress(libgdi32, "EngCreateSemaphore")
	engDeleteClip = doGetProcAddress(libgdi32, "EngDeleteClip")
	engDeletePalette = doGetProcAddress(libgdi32, "EngDeletePalette")
	engDeletePath = doGetProcAddress(libgdi32, "EngDeletePath")
	engDeleteSemaphore = doGetProcAddress(libgdi32, "EngDeleteSemaphore")
	engDeleteSurface = doGetProcAddress(libgdi32, "EngDeleteSurface")
	engEraseSurface = doGetProcAddress(libgdi32, "EngEraseSurface")
	engFillPath = doGetProcAddress(libgdi32, "EngFillPath")
	engFindResource = doGetProcAddress(libgdi32, "EngFindResource")
	engFreeModule = doGetProcAddress(libgdi32, "EngFreeModule")
	engGetCurrentCodePage = doGetProcAddress(libgdi32, "EngGetCurrentCodePage")
	engGetDriverName = doGetProcAddress(libgdi32, "EngGetDriverName")
	engGetPrinterDataFileName = doGetProcAddress(libgdi32, "EngGetPrinterDataFileName")
	engGradientFill = doGetProcAddress(libgdi32, "EngGradientFill")
	engLineTo = doGetProcAddress(libgdi32, "EngLineTo")
	engLoadModule = doGetProcAddress(libgdi32, "EngLoadModule")
	engMarkBandingSurface = doGetProcAddress(libgdi32, "EngMarkBandingSurface")
	engMultiByteToUnicodeN = doGetProcAddress(libgdi32, "EngMultiByteToUnicodeN")
	engMultiByteToWideChar = doGetProcAddress(libgdi32, "EngMultiByteToWideChar")
	engPaint = doGetProcAddress(libgdi32, "EngPaint")
	engPlgBlt = doGetProcAddress(libgdi32, "EngPlgBlt")
	engQueryLocalTime = doGetProcAddress(libgdi32, "EngQueryLocalTime")
	engReleaseSemaphore = doGetProcAddress(libgdi32, "EngReleaseSemaphore")
	engStretchBlt = doGetProcAddress(libgdi32, "EngStretchBlt")
	engStretchBltROP = doGetProcAddress(libgdi32, "EngStretchBltROP")
	engStrokeAndFillPath = doGetProcAddress(libgdi32, "EngStrokeAndFillPath")
	engStrokePath = doGetProcAddress(libgdi32, "EngStrokePath")
	engTextOut = doGetProcAddress(libgdi32, "EngTextOut")
	engTransparentBlt = doGetProcAddress(libgdi32, "EngTransparentBlt")
	engUnicodeToMultiByteN = doGetProcAddress(libgdi32, "EngUnicodeToMultiByteN")
	engUnlockSurface = doGetProcAddress(libgdi32, "EngUnlockSurface")
	engWideCharToMultiByte = doGetProcAddress(libgdi32, "EngWideCharToMultiByte")
	enumEnhMetaFile = doGetProcAddress(libgdi32, "EnumEnhMetaFile")
	enumFontFamiliesEx = doGetProcAddress(libgdi32, "EnumFontFamiliesExW")
	enumFontFamilies = doGetProcAddress(libgdi32, "EnumFontFamiliesW")
	enumFonts = doGetProcAddress(libgdi32, "EnumFontsW")
	enumICMProfiles = doGetProcAddress(libgdi32, "EnumICMProfilesW")
	enumMetaFile = doGetProcAddress(libgdi32, "EnumMetaFile")
	enumObjects = doGetProcAddress(libgdi32, "EnumObjects")
	equalRgn = doGetProcAddress(libgdi32, "EqualRgn")
	escape = doGetProcAddress(libgdi32, "Escape")
	excludeClipRect = doGetProcAddress(libgdi32, "ExcludeClipRect")
	extCreatePen = doGetProcAddress(libgdi32, "ExtCreatePen")
	extCreateRegion = doGetProcAddress(libgdi32, "ExtCreateRegion")
	extEscape = doGetProcAddress(libgdi32, "ExtEscape")
	extFloodFill = doGetProcAddress(libgdi32, "ExtFloodFill")
	extSelectClipRgn = doGetProcAddress(libgdi32, "ExtSelectClipRgn")
	extTextOut = doGetProcAddress(libgdi32, "ExtTextOutW")
	fONTOBJ_cGetAllGlyphHandles = doGetProcAddress(libgdi32, "FONTOBJ_cGetAllGlyphHandles")
	fONTOBJ_cGetGlyphs = doGetProcAddress(libgdi32, "FONTOBJ_cGetGlyphs")
	fONTOBJ_pQueryGlyphAttrs = doGetProcAddress(libgdi32, "FONTOBJ_pQueryGlyphAttrs")
	fONTOBJ_pvTrueTypeFontFile = doGetProcAddress(libgdi32, "FONTOBJ_pvTrueTypeFontFile")
	fONTOBJ_vGetInfo = doGetProcAddress(libgdi32, "FONTOBJ_vGetInfo")
	fillPath = doGetProcAddress(libgdi32, "FillPath")
	fillRgn = doGetProcAddress(libgdi32, "FillRgn")
	fixBrushOrgEx = doGetProcAddress(libgdi32, "FixBrushOrgEx")
	flattenPath = doGetProcAddress(libgdi32, "FlattenPath")
	floodFill = doGetProcAddress(libgdi32, "FloodFill")
	frameRgn = doGetProcAddress(libgdi32, "FrameRgn")
	gdiAlphaBlend = doGetProcAddress(libgdi32, "GdiAlphaBlend")
	gdiComment = doGetProcAddress(libgdi32, "GdiComment")
	gdiFlush = doGetProcAddress(libgdi32, "GdiFlush")
	gdiGetBatchLimit = doGetProcAddress(libgdi32, "GdiGetBatchLimit")
	gdiGradientFill = doGetProcAddress(libgdi32, "GdiGradientFill")
	gdiSetBatchLimit = doGetProcAddress(libgdi32, "GdiSetBatchLimit")
	gdiTransparentBlt = doGetProcAddress(libgdi32, "GdiTransparentBlt")
	getArcDirection = doGetProcAddress(libgdi32, "GetArcDirection")
	getAspectRatioFilterEx = doGetProcAddress(libgdi32, "GetAspectRatioFilterEx")
	getBitmapBits = doGetProcAddress(libgdi32, "GetBitmapBits")
	getBitmapDimensionEx = doGetProcAddress(libgdi32, "GetBitmapDimensionEx")
	getBkColor = doGetProcAddress(libgdi32, "GetBkColor")
	getBkMode = doGetProcAddress(libgdi32, "GetBkMode")
	getBoundsRect = doGetProcAddress(libgdi32, "GetBoundsRect")
	getBrushOrgEx = doGetProcAddress(libgdi32, "GetBrushOrgEx")
	getCharABCWidthsFloat = doGetProcAddress(libgdi32, "GetCharABCWidthsFloatW")
	getCharABCWidthsI = doGetProcAddress(libgdi32, "GetCharABCWidthsI")
	getCharABCWidths = doGetProcAddress(libgdi32, "GetCharABCWidthsW")
	getCharWidth32 = doGetProcAddress(libgdi32, "GetCharWidth32W")
	getCharWidthFloat = doGetProcAddress(libgdi32, "GetCharWidthFloatW")
	getCharWidthI = doGetProcAddress(libgdi32, "GetCharWidthI")
	getCharWidth = doGetProcAddress(libgdi32, "GetCharWidthW")
	getCharacterPlacement = doGetProcAddress(libgdi32, "GetCharacterPlacementW")
	getClipBox = doGetProcAddress(libgdi32, "GetClipBox")
	getClipRgn = doGetProcAddress(libgdi32, "GetClipRgn")
	getColorAdjustment = doGetProcAddress(libgdi32, "GetColorAdjustment")
	getColorSpace = doGetProcAddress(libgdi32, "GetColorSpace")
	getCurrentObject = doGetProcAddress(libgdi32, "GetCurrentObject")
	getCurrentPositionEx = doGetProcAddress(libgdi32, "GetCurrentPositionEx")
	getDCBrushColor = doGetProcAddress(libgdi32, "GetDCBrushColor")
	getDCOrgEx = doGetProcAddress(libgdi32, "GetDCOrgEx")
	getDCPenColor = doGetProcAddress(libgdi32, "GetDCPenColor")
	getDIBColorTable = doGetProcAddress(libgdi32, "GetDIBColorTable")
	getDIBits = doGetProcAddress(libgdi32, "GetDIBits")
	getDeviceCaps = doGetProcAddress(libgdi32, "GetDeviceCaps")
	getDeviceGammaRamp = doGetProcAddress(libgdi32, "GetDeviceGammaRamp")
	getEnhMetaFileBits = doGetProcAddress(libgdi32, "GetEnhMetaFileBits")
	getEnhMetaFileDescription = doGetProcAddress(libgdi32, "GetEnhMetaFileDescriptionW")
	getEnhMetaFileHeader = doGetProcAddress(libgdi32, "GetEnhMetaFileHeader")
	getEnhMetaFilePaletteEntries = doGetProcAddress(libgdi32, "GetEnhMetaFilePaletteEntries")
	getEnhMetaFilePixelFormat = doGetProcAddress(libgdi32, "GetEnhMetaFilePixelFormat")
	getEnhMetaFile = doGetProcAddress(libgdi32, "GetEnhMetaFileW")
	getFontData = doGetProcAddress(libgdi32, "GetFontData")
	getFontLanguageInfo = doGetProcAddress(libgdi32, "GetFontLanguageInfo")
	getFontUnicodeRanges = doGetProcAddress(libgdi32, "GetFontUnicodeRanges")
	getGlyphIndices = doGetProcAddress(libgdi32, "GetGlyphIndicesW")
	getGlyphOutline = doGetProcAddress(libgdi32, "GetGlyphOutlineW")
	getGraphicsMode = doGetProcAddress(libgdi32, "GetGraphicsMode")
	getICMProfile = doGetProcAddress(libgdi32, "GetICMProfileW")
	getKerningPairs = doGetProcAddress(libgdi32, "GetKerningPairsW")
	getLayout = doGetProcAddress(libgdi32, "GetLayout")
	getLogColorSpace = doGetProcAddress(libgdi32, "GetLogColorSpaceW")
	getMapMode = doGetProcAddress(libgdi32, "GetMapMode")
	getMetaFileBitsEx = doGetProcAddress(libgdi32, "GetMetaFileBitsEx")
	getMetaFile = doGetProcAddress(libgdi32, "GetMetaFileW")
	getMetaRgn = doGetProcAddress(libgdi32, "GetMetaRgn")
	getMiterLimit = doGetProcAddress(libgdi32, "GetMiterLimit")
	getNearestColor = doGetProcAddress(libgdi32, "GetNearestColor")
	getNearestPaletteIndex = doGetProcAddress(libgdi32, "GetNearestPaletteIndex")
	getObjectType = doGetProcAddress(libgdi32, "GetObjectType")
	getObject = doGetProcAddress(libgdi32, "GetObjectW")
	getOutlineTextMetrics = doGetProcAddress(libgdi32, "GetOutlineTextMetricsW")
	getPaletteEntries = doGetProcAddress(libgdi32, "GetPaletteEntries")
	getPath = doGetProcAddress(libgdi32, "GetPath")
	getPixel = doGetProcAddress(libgdi32, "GetPixel")
	getPixelFormat = doGetProcAddress(libgdi32, "GetPixelFormat")
	getPolyFillMode = doGetProcAddress(libgdi32, "GetPolyFillMode")
	getROP2 = doGetProcAddress(libgdi32, "GetROP2")
	getRandomRgn = doGetProcAddress(libgdi32, "GetRandomRgn")
	getRasterizerCaps = doGetProcAddress(libgdi32, "GetRasterizerCaps")
	getRegionData = doGetProcAddress(libgdi32, "GetRegionData")
	getRgnBox = doGetProcAddress(libgdi32, "GetRgnBox")
	getStockObject = doGetProcAddress(libgdi32, "GetStockObject")
	getStretchBltMode = doGetProcAddress(libgdi32, "GetStretchBltMode")
	getSystemPaletteEntries = doGetProcAddress(libgdi32, "GetSystemPaletteEntries")
	getSystemPaletteUse = doGetProcAddress(libgdi32, "GetSystemPaletteUse")
	getTextAlign = doGetProcAddress(libgdi32, "GetTextAlign")
	getTextCharacterExtra = doGetProcAddress(libgdi32, "GetTextCharacterExtra")
	getTextCharset = doGetProcAddress(libgdi32, "GetTextCharset")
	getTextCharsetInfo = doGetProcAddress(libgdi32, "GetTextCharsetInfo")
	getTextColor = doGetProcAddress(libgdi32, "GetTextColor")
	getTextExtentExPointI = doGetProcAddress(libgdi32, "GetTextExtentExPointI")
	getTextExtentExPoint = doGetProcAddress(libgdi32, "GetTextExtentExPointW")
	getTextExtentPoint32 = doGetProcAddress(libgdi32, "GetTextExtentPoint32W")
	getTextExtentPointI = doGetProcAddress(libgdi32, "GetTextExtentPointI")
	getTextExtentPoint = doGetProcAddress(libgdi32, "GetTextExtentPointW")
	getTextFace = doGetProcAddress(libgdi32, "GetTextFaceW")
	getTextMetrics = doGetProcAddress(libgdi32, "GetTextMetricsW")
	getViewportExtEx = doGetProcAddress(libgdi32, "GetViewportExtEx")
	getViewportOrgEx = doGetProcAddress(libgdi32, "GetViewportOrgEx")
	getWinMetaFileBits = doGetProcAddress(libgdi32, "GetWinMetaFileBits")
	getWindowExtEx = doGetProcAddress(libgdi32, "GetWindowExtEx")
	getWindowOrgEx = doGetProcAddress(libgdi32, "GetWindowOrgEx")
	getWorldTransform = doGetProcAddress(libgdi32, "GetWorldTransform")
	hT_Get8BPPFormatPalette = doGetProcAddress(libgdi32, "HT_Get8BPPFormatPalette")
	hT_Get8BPPMaskPalette = doGetProcAddress(libgdi32, "HT_Get8BPPMaskPalette")
	intersectClipRect = doGetProcAddress(libgdi32, "IntersectClipRect")
	invertRgn = doGetProcAddress(libgdi32, "InvertRgn")
	lPtoDP = doGetProcAddress(libgdi32, "LPtoDP")
	lineDDA = doGetProcAddress(libgdi32, "LineDDA")
	lineTo = doGetProcAddress(libgdi32, "LineTo")
	maskBlt = doGetProcAddress(libgdi32, "MaskBlt")
	modifyWorldTransform = doGetProcAddress(libgdi32, "ModifyWorldTransform")
	moveToEx = doGetProcAddress(libgdi32, "MoveToEx")
	offsetClipRgn = doGetProcAddress(libgdi32, "OffsetClipRgn")
	offsetRgn = doGetProcAddress(libgdi32, "OffsetRgn")
	offsetViewportOrgEx = doGetProcAddress(libgdi32, "OffsetViewportOrgEx")
	offsetWindowOrgEx = doGetProcAddress(libgdi32, "OffsetWindowOrgEx")
	pATHOBJ_bEnum = doGetProcAddress(libgdi32, "PATHOBJ_bEnum")
	pATHOBJ_bEnumClipLines = doGetProcAddress(libgdi32, "PATHOBJ_bEnumClipLines")
	pATHOBJ_vEnumStart = doGetProcAddress(libgdi32, "PATHOBJ_vEnumStart")
	pATHOBJ_vEnumStartClipLines = doGetProcAddress(libgdi32, "PATHOBJ_vEnumStartClipLines")
	pATHOBJ_vGetBounds = doGetProcAddress(libgdi32, "PATHOBJ_vGetBounds")
	paintRgn = doGetProcAddress(libgdi32, "PaintRgn")
	patBlt = doGetProcAddress(libgdi32, "PatBlt")
	pathToRegion = doGetProcAddress(libgdi32, "PathToRegion")
	pie = doGetProcAddress(libgdi32, "Pie")
	playEnhMetaFile = doGetProcAddress(libgdi32, "PlayEnhMetaFile")
	playEnhMetaFileRecord = doGetProcAddress(libgdi32, "PlayEnhMetaFileRecord")
	playMetaFile = doGetProcAddress(libgdi32, "PlayMetaFile")
	playMetaFileRecord = doGetProcAddress(libgdi32, "PlayMetaFileRecord")
	plgBlt = doGetProcAddress(libgdi32, "PlgBlt")
	polyBezier = doGetProcAddress(libgdi32, "PolyBezier")
	polyBezierTo = doGetProcAddress(libgdi32, "PolyBezierTo")
	polyDraw = doGetProcAddress(libgdi32, "PolyDraw")
	polyPolygon = doGetProcAddress(libgdi32, "PolyPolygon")
	polyPolyline = doGetProcAddress(libgdi32, "PolyPolyline")
	polyTextOut = doGetProcAddress(libgdi32, "PolyTextOutW")
	polygon = doGetProcAddress(libgdi32, "Polygon")
	polyline = doGetProcAddress(libgdi32, "Polyline")
	polylineTo = doGetProcAddress(libgdi32, "PolylineTo")
	ptInRegion = doGetProcAddress(libgdi32, "PtInRegion")
	ptVisible = doGetProcAddress(libgdi32, "PtVisible")
	realizePalette = doGetProcAddress(libgdi32, "RealizePalette")
	rectInRegion = doGetProcAddress(libgdi32, "RectInRegion")
	rectVisible = doGetProcAddress(libgdi32, "RectVisible")
	rectangle = doGetProcAddress(libgdi32, "Rectangle")
	removeFontMemResourceEx = doGetProcAddress(libgdi32, "RemoveFontMemResourceEx")
	removeFontResourceEx = doGetProcAddress(libgdi32, "RemoveFontResourceExW")
	removeFontResource = doGetProcAddress(libgdi32, "RemoveFontResourceW")
	resetDC = doGetProcAddress(libgdi32, "ResetDCW")
	resizePalette = doGetProcAddress(libgdi32, "ResizePalette")
	restoreDC = doGetProcAddress(libgdi32, "RestoreDC")
	roundRect = doGetProcAddress(libgdi32, "RoundRect")
	sTROBJ_bEnum = doGetProcAddress(libgdi32, "STROBJ_bEnum")
	sTROBJ_bEnumPositionsOnly = doGetProcAddress(libgdi32, "STROBJ_bEnumPositionsOnly")
	sTROBJ_bGetAdvanceWidths = doGetProcAddress(libgdi32, "STROBJ_bGetAdvanceWidths")
	sTROBJ_dwGetCodePage = doGetProcAddress(libgdi32, "STROBJ_dwGetCodePage")
	sTROBJ_vEnumStart = doGetProcAddress(libgdi32, "STROBJ_vEnumStart")
	saveDC = doGetProcAddress(libgdi32, "SaveDC")
	scaleViewportExtEx = doGetProcAddress(libgdi32, "ScaleViewportExtEx")
	scaleWindowExtEx = doGetProcAddress(libgdi32, "ScaleWindowExtEx")
	selectClipPath = doGetProcAddress(libgdi32, "SelectClipPath")
	selectClipRgn = doGetProcAddress(libgdi32, "SelectClipRgn")
	selectObject = doGetProcAddress(libgdi32, "SelectObject")
	selectPalette = doGetProcAddress(libgdi32, "SelectPalette")
	setAbortProc = doGetProcAddress(libgdi32, "SetAbortProc")
	setArcDirection = doGetProcAddress(libgdi32, "SetArcDirection")
	setBitmapBits = doGetProcAddress(libgdi32, "SetBitmapBits")
	setBitmapDimensionEx = doGetProcAddress(libgdi32, "SetBitmapDimensionEx")
	setBkColor = doGetProcAddress(libgdi32, "SetBkColor")
	setBkMode = doGetProcAddress(libgdi32, "SetBkMode")
	setBoundsRect = doGetProcAddress(libgdi32, "SetBoundsRect")
	setBrushOrgEx = doGetProcAddress(libgdi32, "SetBrushOrgEx")
	setColorAdjustment = doGetProcAddress(libgdi32, "SetColorAdjustment")
	setColorSpace = doGetProcAddress(libgdi32, "SetColorSpace")
	setDCBrushColor = doGetProcAddress(libgdi32, "SetDCBrushColor")
	setDCPenColor = doGetProcAddress(libgdi32, "SetDCPenColor")
	setDIBColorTable = doGetProcAddress(libgdi32, "SetDIBColorTable")
	setDIBits = doGetProcAddress(libgdi32, "SetDIBits")
	setDIBitsToDevice = doGetProcAddress(libgdi32, "SetDIBitsToDevice")
	setDeviceGammaRamp = doGetProcAddress(libgdi32, "SetDeviceGammaRamp")
	setEnhMetaFileBits = doGetProcAddress(libgdi32, "SetEnhMetaFileBits")
	setGraphicsMode = doGetProcAddress(libgdi32, "SetGraphicsMode")
	setICMMode = doGetProcAddress(libgdi32, "SetICMMode")
	setICMProfile = doGetProcAddress(libgdi32, "SetICMProfileW")
	setLayout = doGetProcAddress(libgdi32, "SetLayout")
	setMapMode = doGetProcAddress(libgdi32, "SetMapMode")
	setMapperFlags = doGetProcAddress(libgdi32, "SetMapperFlags")
	setMetaFileBitsEx = doGetProcAddress(libgdi32, "SetMetaFileBitsEx")
	setMetaRgn = doGetProcAddress(libgdi32, "SetMetaRgn")
	setMiterLimit = doGetProcAddress(libgdi32, "SetMiterLimit")
	setPaletteEntries = doGetProcAddress(libgdi32, "SetPaletteEntries")
	setPixel = doGetProcAddress(libgdi32, "SetPixel")
	setPixelFormat = doGetProcAddress(libgdi32, "SetPixelFormat")
	setPixelV = doGetProcAddress(libgdi32, "SetPixelV")
	setPolyFillMode = doGetProcAddress(libgdi32, "SetPolyFillMode")
	setROP2 = doGetProcAddress(libgdi32, "SetROP2")
	setRectRgn = doGetProcAddress(libgdi32, "SetRectRgn")
	setStretchBltMode = doGetProcAddress(libgdi32, "SetStretchBltMode")
	setSystemPaletteUse = doGetProcAddress(libgdi32, "SetSystemPaletteUse")
	setTextAlign = doGetProcAddress(libgdi32, "SetTextAlign")
	setTextCharacterExtra = doGetProcAddress(libgdi32, "SetTextCharacterExtra")
	setTextColor = doGetProcAddress(libgdi32, "SetTextColor")
	setTextJustification = doGetProcAddress(libgdi32, "SetTextJustification")
	setViewportExtEx = doGetProcAddress(libgdi32, "SetViewportExtEx")
	setViewportOrgEx = doGetProcAddress(libgdi32, "SetViewportOrgEx")
	setWinMetaFileBits = doGetProcAddress(libgdi32, "SetWinMetaFileBits")
	setWindowExtEx = doGetProcAddress(libgdi32, "SetWindowExtEx")
	setWindowOrgEx = doGetProcAddress(libgdi32, "SetWindowOrgEx")
	setWorldTransform = doGetProcAddress(libgdi32, "SetWorldTransform")
	startDoc = doGetProcAddress(libgdi32, "StartDocW")
	startPage = doGetProcAddress(libgdi32, "StartPage")
	stretchBlt = doGetProcAddress(libgdi32, "StretchBlt")
	stretchDIBits = doGetProcAddress(libgdi32, "StretchDIBits")
	strokeAndFillPath = doGetProcAddress(libgdi32, "StrokeAndFillPath")
	strokePath = doGetProcAddress(libgdi32, "StrokePath")
	swapBuffers = doGetProcAddress(libgdi32, "SwapBuffers")
	textOut = doGetProcAddress(libgdi32, "TextOutW")
	translateCharsetInfo = doGetProcAddress(libgdi32, "TranslateCharsetInfo")
	unrealizeObject = doGetProcAddress(libgdi32, "UnrealizeObject")
	updateColors = doGetProcAddress(libgdi32, "UpdateColors")
	updateICMRegKey = doGetProcAddress(libgdi32, "UpdateICMRegKeyW")
	widenPath = doGetProcAddress(libgdi32, "WidenPath")
	xFORMOBJ_bApplyXform = doGetProcAddress(libgdi32, "XFORMOBJ_bApplyXform")
	xFORMOBJ_iGetXform = doGetProcAddress(libgdi32, "XFORMOBJ_iGetXform")
	xLATEOBJ_cGetPalette = doGetProcAddress(libgdi32, "XLATEOBJ_cGetPalette")
	xLATEOBJ_hGetColorTransform = doGetProcAddress(libgdi32, "XLATEOBJ_hGetColorTransform")
	xLATEOBJ_iXlate = doGetProcAddress(libgdi32, "XLATEOBJ_iXlate")
	enableEUDC = doGetProcAddress(libgdi32, "EnableEUDC")
	fontIsLinked = doGetProcAddress(libgdi32, "FontIsLinked")
	gdiDescribePixelFormat = doGetProcAddress(libgdi32, "GdiDescribePixelFormat")
	gdiDrawStream = doGetProcAddress(libgdi32, "GdiDrawStream")
	gdiGetCharDimensions = doGetProcAddress(libgdi32, "GdiGetCharDimensions")
	gdiGetCodePage = doGetProcAddress(libgdi32, "GdiGetCodePage")
	gdiGetSpoolMessage = doGetProcAddress(libgdi32, "GdiGetSpoolMessage")
	gdiInitSpool = doGetProcAddress(libgdi32, "GdiInitSpool")
	gdiInitializeLanguagePack = doGetProcAddress(libgdi32, "GdiInitializeLanguagePack")
	gdiIsMetaFileDC = doGetProcAddress(libgdi32, "GdiIsMetaFileDC")
	gdiIsMetaPrintDC = doGetProcAddress(libgdi32, "GdiIsMetaPrintDC")
	gdiIsPlayMetafileDC = doGetProcAddress(libgdi32, "GdiIsPlayMetafileDC")
	gdiRealizationInfo = doGetProcAddress(libgdi32, "GdiRealizationInfo")
	gdiSetPixelFormat = doGetProcAddress(libgdi32, "GdiSetPixelFormat")
	gdiSwapBuffers = doGetProcAddress(libgdi32, "GdiSwapBuffers")
	getFontResourceInfoW = doGetProcAddress(libgdi32, "GetFontResourceInfoW")
	getRelAbs = doGetProcAddress(libgdi32, "GetRelAbs")
	getTransform = doGetProcAddress(libgdi32, "GetTransform")
	mirrorRgn = doGetProcAddress(libgdi32, "MirrorRgn")
	namedEscape = doGetProcAddress(libgdi32, "NamedEscape")
	setMagicColors = doGetProcAddress(libgdi32, "SetMagicColors")
	setRelAbs = doGetProcAddress(libgdi32, "SetRelAbs")
	setVirtualResolution = doGetProcAddress(libgdi32, "SetVirtualResolution")
}

func AbortDoc(hdc HDC) int32 {
	ret1 := syscall3(abortDoc, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func AbortPath(hdc HDC) bool {
	ret1 := syscall3(abortPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func AddFontMemResourceEx(pFileView uintptr, cjSize uint32, pvResrved uintptr, pNumFonts *uint32) HANDLE {
	ret1 := syscall6(addFontMemResourceEx, 4,
		pFileView,
		uintptr(cjSize),
		pvResrved,
		uintptr(unsafe.Pointer(pNumFonts)),
		0,
		0)
	return HANDLE(ret1)
}

func AddFontResourceEx(name string, fl uint32, res uintptr) int32 {
	nameStr := unicode16FromString(name)
	ret1 := syscall3(addFontResourceEx, 3,
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(fl),
		res)
	return int32(ret1)
}

func AddFontResource(unnamed0 string) int32 {
	unnamed0Str := unicode16FromString(unnamed0)
	ret1 := syscall3(addFontResource, 1,
		uintptr(unsafe.Pointer(&unnamed0Str[0])),
		0,
		0)
	return int32(ret1)
}

func AngleArc(hdc HDC, x int32, y int32, r uint32, startAngle float32, sweepAngle float32) bool {
	ret1 := syscall6(angleArc, 6,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(r),
		uintptr(startAngle),
		uintptr(sweepAngle))
	return ret1 != 0
}

func AnimatePalette(hPal HPALETTE, iStartIndex uint32, cEntries uint32, ppe /*const*/ *PALETTEENTRY) bool {
	ret1 := syscall6(animatePalette, 4,
		uintptr(hPal),
		uintptr(iStartIndex),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(ppe)),
		0,
		0)
	return ret1 != 0
}

func Arc(hdc HDC, x1 int32, y1 int32, x2 int32, y2 int32, x3 int32, y3 int32, x4 int32, y4 int32) bool {
	ret1 := syscall9(arc, 9,
		uintptr(hdc),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(x3),
		uintptr(y3),
		uintptr(x4),
		uintptr(y4))
	return ret1 != 0
}

func ArcTo(hdc HDC, left int32, top int32, right int32, bottom int32, xr1 int32, yr1 int32, xr2 int32, yr2 int32) bool {
	ret1 := syscall9(arcTo, 9,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		uintptr(xr1),
		uintptr(yr1),
		uintptr(xr2),
		uintptr(yr2))
	return ret1 != 0
}

func BRUSHOBJ_hGetColorTransform(pbo *BRUSHOBJ) HANDLE {
	ret1 := syscall3(bRUSHOBJ_hGetColorTransform, 1,
		uintptr(unsafe.Pointer(pbo)),
		0,
		0)
	return HANDLE(ret1)
}

func BRUSHOBJ_pvAllocRbrush(pbo *BRUSHOBJ, cj ULONG) uintptr {
	ret1 := syscall3(bRUSHOBJ_pvAllocRbrush, 2,
		uintptr(unsafe.Pointer(pbo)),
		uintptr(cj),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func BRUSHOBJ_pvGetRbrush(pbo *BRUSHOBJ) uintptr {
	ret1 := syscall3(bRUSHOBJ_pvGetRbrush, 1,
		uintptr(unsafe.Pointer(pbo)),
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func BRUSHOBJ_ulGetBrushColor(pbo *BRUSHOBJ) ULONG {
	ret1 := syscall3(bRUSHOBJ_ulGetBrushColor, 1,
		uintptr(unsafe.Pointer(pbo)),
		0,
		0)
	return ULONG(ret1)
}

func BeginPath(hdc HDC) bool {
	ret1 := syscall3(beginPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func BitBlt(hdc HDC, x int32, y int32, cx int32, cy int32, hdcSrc HDC, x1 int32, y1 int32, rop uint32) bool {
	ret1 := syscall9(bitBlt, 9,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(hdcSrc),
		uintptr(x1),
		uintptr(y1),
		uintptr(rop))
	return ret1 != 0
}

func CLIPOBJ_bEnum(pco *CLIPOBJ, cj ULONG, pv *ULONG) bool {
	ret1 := syscall3(cLIPOBJ_bEnum, 3,
		uintptr(unsafe.Pointer(pco)),
		uintptr(cj),
		uintptr(unsafe.Pointer(pv)))
	return ret1 != 0
}

func CLIPOBJ_cEnumStart(pco *CLIPOBJ, bAll bool, iType ULONG, iDirection ULONG, cLimit ULONG) ULONG {
	ret1 := syscall6(cLIPOBJ_cEnumStart, 5,
		uintptr(unsafe.Pointer(pco)),
		getUintptrFromBool(bAll),
		uintptr(iType),
		uintptr(iDirection),
		uintptr(cLimit),
		0)
	return ULONG(ret1)
}

func CancelDC(hdc HDC) bool {
	ret1 := syscall3(cancelDC, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func CheckColorsInGamut(hdc HDC, lpRGBTriple LPVOID, dlpBuffer LPVOID, nCount uint32) bool {
	ret1 := syscall6(checkColorsInGamut, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpRGBTriple)),
		uintptr(unsafe.Pointer(dlpBuffer)),
		uintptr(nCount),
		0,
		0)
	return ret1 != 0
}

func ChoosePixelFormat(hdc HDC, ppfd /*const*/ *PIXELFORMATDESCRIPTOR) int32 {
	ret1 := syscall3(choosePixelFormat, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(ppfd)),
		0)
	return int32(ret1)
}

func Chord(hdc HDC, x1 int32, y1 int32, x2 int32, y2 int32, x3 int32, y3 int32, x4 int32, y4 int32) bool {
	ret1 := syscall9(chord, 9,
		uintptr(hdc),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(x3),
		uintptr(y3),
		uintptr(x4),
		uintptr(y4))
	return ret1 != 0
}

func CloseEnhMetaFile(hdc HDC) HENHMETAFILE {
	ret1 := syscall3(closeEnhMetaFile, 1,
		uintptr(hdc),
		0,
		0)
	return HENHMETAFILE(ret1)
}

func CloseFigure(hdc HDC) bool {
	ret1 := syscall3(closeFigure, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func CloseMetaFile(hdc HDC) HMETAFILE {
	ret1 := syscall3(closeMetaFile, 1,
		uintptr(hdc),
		0,
		0)
	return HMETAFILE(ret1)
}

func ColorCorrectPalette(hdc HDC, hPal HPALETTE, deFirst uint32, num uint32) bool {
	ret1 := syscall6(colorCorrectPalette, 4,
		uintptr(hdc),
		uintptr(hPal),
		uintptr(deFirst),
		uintptr(num),
		0,
		0)
	return ret1 != 0
}

func ColorMatchToTarget(hdc HDC, hdcTarget HDC, action uint32) bool {
	ret1 := syscall3(colorMatchToTarget, 3,
		uintptr(hdc),
		uintptr(hdcTarget),
		uintptr(action))
	return ret1 != 0
}

func CombineRgn(hrgnDst HRGN, hrgnSrc1 HRGN, hrgnSrc2 HRGN, iMode int32) int32 {
	ret1 := syscall6(combineRgn, 4,
		uintptr(hrgnDst),
		uintptr(hrgnSrc1),
		uintptr(hrgnSrc2),
		uintptr(iMode),
		0,
		0)
	return int32(ret1)
}

func CombineTransform(lpxfOut *XFORM, lpxf1 /*const*/ *XFORM, lpxf2 /*const*/ *XFORM) bool {
	ret1 := syscall3(combineTransform, 3,
		uintptr(unsafe.Pointer(lpxfOut)),
		uintptr(unsafe.Pointer(lpxf1)),
		uintptr(unsafe.Pointer(lpxf2)))
	return ret1 != 0
}

func CopyEnhMetaFile(hEnh HENHMETAFILE, lpFileName string) HENHMETAFILE {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(copyEnhMetaFile, 2,
		uintptr(hEnh),
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0)
	return HENHMETAFILE(ret1)
}

func CopyMetaFile(unnamed0 HMETAFILE, unnamed1 string) HMETAFILE {
	unnamed1Str := unicode16FromString(unnamed1)
	ret1 := syscall3(copyMetaFile, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(&unnamed1Str[0])),
		0)
	return HMETAFILE(ret1)
}

func CreateBitmap(nWidth int32, nHeight int32, nPlanes uint32, nBitCount uint32, lpBits /*const*/ uintptr) HBITMAP {
	ret1 := syscall6(createBitmap, 5,
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(nPlanes),
		uintptr(nBitCount),
		lpBits,
		0)
	return HBITMAP(ret1)
}

func CreateBitmapIndirect(pbm /*const*/ *BITMAP) HBITMAP {
	ret1 := syscall3(createBitmapIndirect, 1,
		uintptr(unsafe.Pointer(pbm)),
		0,
		0)
	return HBITMAP(ret1)
}

func CreateBrushIndirect(plbrush /*const*/ *LOGBRUSH) HBRUSH {
	ret1 := syscall3(createBrushIndirect, 1,
		uintptr(unsafe.Pointer(plbrush)),
		0,
		0)
	return HBRUSH(ret1)
}

func CreateColorSpace(lplcs LPLOGCOLORSPACE) HCOLORSPACE {
	ret1 := syscall3(createColorSpace, 1,
		uintptr(unsafe.Pointer(lplcs)),
		0,
		0)
	return HCOLORSPACE(ret1)
}

func CreateCompatibleBitmap(hdc HDC, cx int32, cy int32) HBITMAP {
	ret1 := syscall3(createCompatibleBitmap, 3,
		uintptr(hdc),
		uintptr(cx),
		uintptr(cy))
	return HBITMAP(ret1)
}

func CreateCompatibleDC(hdc HDC) HDC {
	ret1 := syscall3(createCompatibleDC, 1,
		uintptr(hdc),
		0,
		0)
	return HDC(ret1)
}

func CreateDC(pwszDriver string, pwszDevice string, pszPort string, pdm /*const*/ *DEVMODE) HDC {
	pwszDriverStr := unicode16FromString(pwszDriver)
	pwszDeviceStr := unicode16FromString(pwszDevice)
	pszPortStr := unicode16FromString(pszPort)
	ret1 := syscall6(createDC, 4,
		uintptr(unsafe.Pointer(&pwszDriverStr[0])),
		uintptr(unsafe.Pointer(&pwszDeviceStr[0])),
		uintptr(unsafe.Pointer(&pszPortStr[0])),
		uintptr(unsafe.Pointer(pdm)),
		0,
		0)
	return HDC(ret1)
}

func CreateDIBPatternBrush(h HGLOBAL, iUsage uint32) HBRUSH {
	ret1 := syscall3(createDIBPatternBrush, 2,
		uintptr(h),
		uintptr(iUsage),
		0)
	return HBRUSH(ret1)
}

func CreateDIBPatternBrushPt(lpPackedDIB /*const*/ uintptr, iUsage uint32) HBRUSH {
	ret1 := syscall3(createDIBPatternBrushPt, 2,
		lpPackedDIB,
		uintptr(iUsage),
		0)
	return HBRUSH(ret1)
}

func CreateDIBSection(hdc HDC, lpbmi /*const*/ *BITMAPINFO, usage uint32, ppvBits uintptr, hSection HANDLE, offset uint32) HBITMAP {
	ret1 := syscall6(createDIBSection, 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(usage),
		ppvBits,
		uintptr(hSection),
		uintptr(offset))
	return HBITMAP(ret1)
}

func CreateDIBitmap(hdc HDC, pbmih /*const*/ *BITMAPINFOHEADER, flInit uint32, pjBits /*const*/ uintptr, pbmi /*const*/ *BITMAPINFO, iUsage uint32) HBITMAP {
	ret1 := syscall6(createDIBitmap, 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pbmih)),
		uintptr(flInit),
		pjBits,
		uintptr(unsafe.Pointer(pbmi)),
		uintptr(iUsage))
	return HBITMAP(ret1)
}

func CreateDiscardableBitmap(hdc HDC, cx int32, cy int32) HBITMAP {
	ret1 := syscall3(createDiscardableBitmap, 3,
		uintptr(hdc),
		uintptr(cx),
		uintptr(cy))
	return HBITMAP(ret1)
}

func CreateEllipticRgn(x1 int32, y1 int32, x2 int32, y2 int32) HRGN {
	ret1 := syscall6(createEllipticRgn, 4,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		0,
		0)
	return HRGN(ret1)
}

func CreateEllipticRgnIndirect(lprect /*const*/ *RECT) HRGN {
	ret1 := syscall3(createEllipticRgnIndirect, 1,
		uintptr(unsafe.Pointer(lprect)),
		0,
		0)
	return HRGN(ret1)
}

func CreateEnhMetaFile(hdc HDC, lpFilename string, lprc /*const*/ *RECT, lpDesc string) HDC {
	lpFilenameStr := unicode16FromString(lpFilename)
	lpDescStr := unicode16FromString(lpDesc)
	ret1 := syscall6(createEnhMetaFile, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpFilenameStr[0])),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(unsafe.Pointer(&lpDescStr[0])),
		0,
		0)
	return HDC(ret1)
}

func CreateFontIndirectEx(unnamed0 /*const*/ *ENUMLOGFONTEXDV) HFONT {
	ret1 := syscall3(createFontIndirectEx, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return HFONT(ret1)
}

func CreateFontIndirect(lplf /*const*/ *LOGFONT) HFONT {
	ret1 := syscall3(createFontIndirect, 1,
		uintptr(unsafe.Pointer(lplf)),
		0,
		0)
	return HFONT(ret1)
}

func CreateFont(cHeight int32, cWidth int32, cEscapement int32, cOrientation int32, cWeight int32, bItalic uint32, bUnderline uint32, bStrikeOut uint32, iCharSet uint32, iOutPrecision uint32, iClipPrecision uint32, iQuality uint32, iPitchAndFamily uint32, pszFaceName string) HFONT {
	pszFaceNameStr := unicode16FromString(pszFaceName)
	ret1 := syscall15(createFont, 14,
		uintptr(cHeight),
		uintptr(cWidth),
		uintptr(cEscapement),
		uintptr(cOrientation),
		uintptr(cWeight),
		uintptr(bItalic),
		uintptr(bUnderline),
		uintptr(bStrikeOut),
		uintptr(iCharSet),
		uintptr(iOutPrecision),
		uintptr(iClipPrecision),
		uintptr(iQuality),
		uintptr(iPitchAndFamily),
		uintptr(unsafe.Pointer(&pszFaceNameStr[0])),
		0)
	return HFONT(ret1)
}

func CreateHalftonePalette(hdc HDC) HPALETTE {
	ret1 := syscall3(createHalftonePalette, 1,
		uintptr(hdc),
		0,
		0)
	return HPALETTE(ret1)
}

func CreateHatchBrush(iHatch int32, color COLORREF) HBRUSH {
	ret1 := syscall3(createHatchBrush, 2,
		uintptr(iHatch),
		uintptr(color),
		0)
	return HBRUSH(ret1)
}

func CreateIC(pszDriver string, pszDevice string, pszPort string, pdm /*const*/ *DEVMODE) HDC {
	pszDriverStr := unicode16FromString(pszDriver)
	pszDeviceStr := unicode16FromString(pszDevice)
	pszPortStr := unicode16FromString(pszPort)
	ret1 := syscall6(createIC, 4,
		uintptr(unsafe.Pointer(&pszDriverStr[0])),
		uintptr(unsafe.Pointer(&pszDeviceStr[0])),
		uintptr(unsafe.Pointer(&pszPortStr[0])),
		uintptr(unsafe.Pointer(pdm)),
		0,
		0)
	return HDC(ret1)
}

func CreateMetaFile(pszFile string) HDC {
	pszFileStr := unicode16FromString(pszFile)
	ret1 := syscall3(createMetaFile, 1,
		uintptr(unsafe.Pointer(&pszFileStr[0])),
		0,
		0)
	return HDC(ret1)
}

func CreatePalette(plpal /*const*/ *LOGPALETTE) HPALETTE {
	ret1 := syscall3(createPalette, 1,
		uintptr(unsafe.Pointer(plpal)),
		0,
		0)
	return HPALETTE(ret1)
}

func CreatePatternBrush(hbm HBITMAP) HBRUSH {
	ret1 := syscall3(createPatternBrush, 1,
		uintptr(hbm),
		0,
		0)
	return HBRUSH(ret1)
}

func CreatePen(iStyle int32, cWidth int32, color COLORREF) HPEN {
	ret1 := syscall3(createPen, 3,
		uintptr(iStyle),
		uintptr(cWidth),
		uintptr(color))
	return HPEN(ret1)
}

func CreatePenIndirect(plpen /*const*/ *LOGPEN) HPEN {
	ret1 := syscall3(createPenIndirect, 1,
		uintptr(unsafe.Pointer(plpen)),
		0,
		0)
	return HPEN(ret1)
}

func CreatePolyPolygonRgn(pptl /*const*/ *POINT, pc /*const*/ *int32, cPoly int32, iMode int32) HRGN {
	ret1 := syscall6(createPolyPolygonRgn, 4,
		uintptr(unsafe.Pointer(pptl)),
		uintptr(unsafe.Pointer(pc)),
		uintptr(cPoly),
		uintptr(iMode),
		0,
		0)
	return HRGN(ret1)
}

func CreatePolygonRgn(pptl /*const*/ *POINT, cPoint int32, iMode int32) HRGN {
	ret1 := syscall3(createPolygonRgn, 3,
		uintptr(unsafe.Pointer(pptl)),
		uintptr(cPoint),
		uintptr(iMode))
	return HRGN(ret1)
}

func CreateRectRgn(x1 int32, y1 int32, x2 int32, y2 int32) HRGN {
	ret1 := syscall6(createRectRgn, 4,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		0,
		0)
	return HRGN(ret1)
}

func CreateRectRgnIndirect(lprect /*const*/ *RECT) HRGN {
	ret1 := syscall3(createRectRgnIndirect, 1,
		uintptr(unsafe.Pointer(lprect)),
		0,
		0)
	return HRGN(ret1)
}

func CreateRoundRectRgn(x1 int32, y1 int32, x2 int32, y2 int32, w int32, h int32) HRGN {
	ret1 := syscall6(createRoundRectRgn, 6,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(w),
		uintptr(h))
	return HRGN(ret1)
}

func CreateScalableFontResource(fdwHidden uint32, lpszFont string, lpszFile string, lpszPath string) bool {
	lpszFontStr := unicode16FromString(lpszFont)
	lpszFileStr := unicode16FromString(lpszFile)
	lpszPathStr := unicode16FromString(lpszPath)
	ret1 := syscall6(createScalableFontResource, 4,
		uintptr(fdwHidden),
		uintptr(unsafe.Pointer(&lpszFontStr[0])),
		uintptr(unsafe.Pointer(&lpszFileStr[0])),
		uintptr(unsafe.Pointer(&lpszPathStr[0])),
		0,
		0)
	return ret1 != 0
}

func CreateSolidBrush(color COLORREF) HBRUSH {
	ret1 := syscall3(createSolidBrush, 1,
		uintptr(color),
		0,
		0)
	return HBRUSH(ret1)
}

func DPtoLP(hdc HDC, lppt *POINT, c int32) bool {
	ret1 := syscall3(dPtoLP, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppt)),
		uintptr(c))
	return ret1 != 0
}

func DeleteColorSpace(hcs HCOLORSPACE) bool {
	ret1 := syscall3(deleteColorSpace, 1,
		uintptr(hcs),
		0,
		0)
	return ret1 != 0
}

func DeleteDC(hdc HDC) bool {
	ret1 := syscall3(deleteDC, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func DeleteEnhMetaFile(hmf HENHMETAFILE) bool {
	ret1 := syscall3(deleteEnhMetaFile, 1,
		uintptr(hmf),
		0,
		0)
	return ret1 != 0
}

func DeleteMetaFile(hmf HMETAFILE) bool {
	ret1 := syscall3(deleteMetaFile, 1,
		uintptr(hmf),
		0,
		0)
	return ret1 != 0
}

func DeleteObject(ho HGDIOBJ) bool {
	ret1 := syscall3(deleteObject, 1,
		uintptr(ho),
		0,
		0)
	return ret1 != 0
}

func DescribePixelFormat(hdc HDC, iPixelFormat int32, nBytes uint32, ppfd *PIXELFORMATDESCRIPTOR) int32 {
	ret1 := syscall6(describePixelFormat, 4,
		uintptr(hdc),
		uintptr(iPixelFormat),
		uintptr(nBytes),
		uintptr(unsafe.Pointer(ppfd)),
		0,
		0)
	return int32(ret1)
}

func DrawEscape(hdc HDC, iEscape int32, cjIn int32, lpIn /*const*/ LPCSTR) int32 {
	ret1 := syscall6(drawEscape, 4,
		uintptr(hdc),
		uintptr(iEscape),
		uintptr(cjIn),
		uintptr(unsafe.Pointer(lpIn)),
		0,
		0)
	return int32(ret1)
}

func Ellipse(hdc HDC, left int32, top int32, right int32, bottom int32) bool {
	ret1 := syscall6(ellipse, 5,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		0)
	return ret1 != 0
}

func EndDoc(hdc HDC) int32 {
	ret1 := syscall3(endDoc, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func EndPage(hdc HDC) int32 {
	ret1 := syscall3(endPage, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func EndPath(hdc HDC) bool {
	ret1 := syscall3(endPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func EngAcquireSemaphore(hsem HSEMAPHORE) {
	syscall3(engAcquireSemaphore, 1,
		uintptr(hsem),
		0,
		0)
}

func EngAlphaBlend(psoDest *SURFOBJ, psoSrc *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, prclDest *RECTL, prclSrc *RECTL, pBlendObj *BLENDOBJ) bool {
	ret1 := syscall9(engAlphaBlend, 7,
		uintptr(unsafe.Pointer(psoDest)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(prclDest)),
		uintptr(unsafe.Pointer(prclSrc)),
		uintptr(unsafe.Pointer(pBlendObj)),
		0,
		0)
	return ret1 != 0
}

func EngAssociateSurface(hsurf HSURF, hdev HDEV, flHooks FLONG) bool {
	ret1 := syscall3(engAssociateSurface, 3,
		uintptr(hsurf),
		uintptr(hdev),
		uintptr(flHooks))
	return ret1 != 0
}

func EngBitBlt(psoTrg *SURFOBJ, psoSrc *SURFOBJ, psoMask *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, prclTrg *RECTL, pptlSrc *POINTL, pptlMask *POINTL, pbo *BRUSHOBJ, pptlBrush *POINTL, rop4 ROP4) bool {
	ret1 := syscall12(engBitBlt, 11,
		uintptr(unsafe.Pointer(psoTrg)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(psoMask)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(prclTrg)),
		uintptr(unsafe.Pointer(pptlSrc)),
		uintptr(unsafe.Pointer(pptlMask)),
		uintptr(unsafe.Pointer(pbo)),
		uintptr(unsafe.Pointer(pptlBrush)),
		uintptr(rop4),
		0)
	return ret1 != 0
}

func EngCheckAbort(pso *SURFOBJ) bool {
	ret1 := syscall3(engCheckAbort, 1,
		uintptr(unsafe.Pointer(pso)),
		0,
		0)
	return ret1 != 0
}

func EngCopyBits(psoDest *SURFOBJ, psoSrc *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, prclDest *RECTL, pptlSrc *POINTL) bool {
	ret1 := syscall6(engCopyBits, 6,
		uintptr(unsafe.Pointer(psoDest)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(prclDest)),
		uintptr(unsafe.Pointer(pptlSrc)))
	return ret1 != 0
}

func EngCreateBitmap(sizl SIZEL, lWidth LONG, iFormat ULONG, fl FLONG, pvBits uintptr) HBITMAP {
	ret1 := syscall6(engCreateBitmap, 6,
		uintptr(sizl.Cx),
		uintptr(sizl.Cy),
		uintptr(lWidth),
		uintptr(iFormat),
		uintptr(fl),
		pvBits)
	return HBITMAP(ret1)
}

func EngCreateDeviceBitmap(dhsurf DHSURF, sizl SIZEL, iFormatCompat ULONG) HBITMAP {
	ret1 := syscall6(engCreateDeviceBitmap, 4,
		uintptr(dhsurf),
		uintptr(sizl.Cx),
		uintptr(sizl.Cy),
		uintptr(iFormatCompat),
		0,
		0)
	return HBITMAP(ret1)
}

func EngCreateDeviceSurface(dhsurf DHSURF, sizl SIZEL, iFormatCompat ULONG) HSURF {
	ret1 := syscall6(engCreateDeviceSurface, 4,
		uintptr(dhsurf),
		uintptr(sizl.Cx),
		uintptr(sizl.Cy),
		uintptr(iFormatCompat),
		0,
		0)
	return HSURF(ret1)
}

func EngCreatePalette(iMode ULONG, cColors ULONG, pulColors *ULONG, flRed FLONG, flGreen FLONG, flBlue FLONG) HPALETTE {
	ret1 := syscall6(engCreatePalette, 6,
		uintptr(iMode),
		uintptr(cColors),
		uintptr(unsafe.Pointer(pulColors)),
		uintptr(flRed),
		uintptr(flGreen),
		uintptr(flBlue))
	return HPALETTE(ret1)
}

func EngCreateSemaphore() HSEMAPHORE {
	ret1 := syscall3(engCreateSemaphore, 0,
		0,
		0,
		0)
	return HSEMAPHORE(ret1)
}

func EngDeleteClip(pco *CLIPOBJ) {
	syscall3(engDeleteClip, 1,
		uintptr(unsafe.Pointer(pco)),
		0,
		0)
}

func EngDeletePalette(hpal HPALETTE) bool {
	ret1 := syscall3(engDeletePalette, 1,
		uintptr(hpal),
		0,
		0)
	return ret1 != 0
}

func EngDeletePath(ppo *PATHOBJ) {
	syscall3(engDeletePath, 1,
		uintptr(unsafe.Pointer(ppo)),
		0,
		0)
}

func EngDeleteSemaphore(hsem HSEMAPHORE) {
	syscall3(engDeleteSemaphore, 1,
		uintptr(hsem),
		0,
		0)
}

func EngDeleteSurface(hsurf HSURF) bool {
	ret1 := syscall3(engDeleteSurface, 1,
		uintptr(hsurf),
		0,
		0)
	return ret1 != 0
}

func EngEraseSurface(pso *SURFOBJ, prcl *RECTL, iColor ULONG) bool {
	ret1 := syscall3(engEraseSurface, 3,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(prcl)),
		uintptr(iColor))
	return ret1 != 0
}

func EngFillPath(pso *SURFOBJ, ppo *PATHOBJ, pco *CLIPOBJ, pbo *BRUSHOBJ, pptlBrushOrg *POINTL, mix MIX, flOptions FLONG) bool {
	ret1 := syscall9(engFillPath, 7,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(ppo)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pbo)),
		uintptr(unsafe.Pointer(pptlBrushOrg)),
		uintptr(mix),
		uintptr(flOptions),
		0,
		0)
	return ret1 != 0
}

func EngFindResource(h HANDLE, iName int32, iType int32, pulSize *uint32) uintptr {
	ret1 := syscall6(engFindResource, 4,
		uintptr(h),
		uintptr(iName),
		uintptr(iType),
		uintptr(unsafe.Pointer(pulSize)),
		0,
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func EngFreeModule(h HANDLE) {
	syscall3(engFreeModule, 1,
		uintptr(h),
		0,
		0)
}

func EngGetCurrentCodePage(oemCodePage PUSHORT, ansiCodePage PUSHORT) {
	syscall3(engGetCurrentCodePage, 2,
		uintptr(unsafe.Pointer(oemCodePage)),
		uintptr(unsafe.Pointer(ansiCodePage)),
		0)
}

func EngGetDriverName(hdev HDEV) LPWSTR {
	ret1 := syscall3(engGetDriverName, 1,
		uintptr(hdev),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func EngGetPrinterDataFileName(hdev HDEV) LPWSTR {
	ret1 := syscall3(engGetPrinterDataFileName, 1,
		uintptr(hdev),
		0,
		0)
	return (LPWSTR)(unsafe.Pointer(ret1))
}

func EngGradientFill(psoDest *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, pVertex *TRIVERTEX, nVertex ULONG, pMesh uintptr, nMesh ULONG, prclExtents *RECTL, pptlDitherOrg *POINTL, ulMode ULONG) bool {
	ret1 := syscall12(engGradientFill, 10,
		uintptr(unsafe.Pointer(psoDest)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(pVertex)),
		uintptr(nVertex),
		pMesh,
		uintptr(nMesh),
		uintptr(unsafe.Pointer(prclExtents)),
		uintptr(unsafe.Pointer(pptlDitherOrg)),
		uintptr(ulMode),
		0,
		0)
	return ret1 != 0
}

func EngLineTo(pso *SURFOBJ, pco *CLIPOBJ, pbo *BRUSHOBJ, x1 LONG, y1 LONG, x2 LONG, y2 LONG, prclBounds *RECTL, mix MIX) bool {
	ret1 := syscall9(engLineTo, 9,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pbo)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		uintptr(unsafe.Pointer(prclBounds)),
		uintptr(mix))
	return ret1 != 0
}

func EngLoadModule(pwsz LPWSTR) HANDLE {
	ret1 := syscall3(engLoadModule, 1,
		uintptr(unsafe.Pointer(pwsz)),
		0,
		0)
	return HANDLE(ret1)
}

func EngMarkBandingSurface(hsurf HSURF) bool {
	ret1 := syscall3(engMarkBandingSurface, 1,
		uintptr(hsurf),
		0,
		0)
	return ret1 != 0
}

func EngMultiByteToUnicodeN(unicodeString LPWSTR, maxBytesInUnicodeString ULONG, bytesInUnicodeString *uint32, multiByteString PCHAR, bytesInMultiByteString ULONG) {
	syscall6(engMultiByteToUnicodeN, 5,
		uintptr(unsafe.Pointer(unicodeString)),
		uintptr(maxBytesInUnicodeString),
		uintptr(unsafe.Pointer(bytesInUnicodeString)),
		uintptr(unsafe.Pointer(multiByteString)),
		uintptr(bytesInMultiByteString),
		0)
}

func EngMultiByteToWideChar(codePage uint32, wideCharString LPWSTR, bytesInWideCharString int32, multiByteString LPSTR, bytesInMultiByteString int32) int32 {
	ret1 := syscall6(engMultiByteToWideChar, 5,
		uintptr(codePage),
		uintptr(unsafe.Pointer(wideCharString)),
		uintptr(bytesInWideCharString),
		uintptr(unsafe.Pointer(multiByteString)),
		uintptr(bytesInMultiByteString),
		0)
	return int32(ret1)
}

func EngPaint(pso *SURFOBJ, pco *CLIPOBJ, pbo *BRUSHOBJ, pptlBrushOrg *POINTL, mix MIX) bool {
	ret1 := syscall6(engPaint, 5,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pbo)),
		uintptr(unsafe.Pointer(pptlBrushOrg)),
		uintptr(mix),
		0)
	return ret1 != 0
}

func EngPlgBlt(psoTrg *SURFOBJ, psoSrc *SURFOBJ, psoMsk *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, pca *COLORADJUSTMENT, pptlBrushOrg *POINTL, pptfx *POINTFIX, prcl *RECTL, pptl *POINTL, iMode ULONG) bool {
	ret1 := syscall12(engPlgBlt, 11,
		uintptr(unsafe.Pointer(psoTrg)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(psoMsk)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(pca)),
		uintptr(unsafe.Pointer(pptlBrushOrg)),
		uintptr(unsafe.Pointer(pptfx)),
		uintptr(unsafe.Pointer(prcl)),
		uintptr(unsafe.Pointer(pptl)),
		uintptr(iMode),
		0)
	return ret1 != 0
}

func EngQueryLocalTime(ptf PENG_TIME_FIELDS) {
	syscall3(engQueryLocalTime, 1,
		uintptr(unsafe.Pointer(ptf)),
		0,
		0)
}

func EngReleaseSemaphore(hsem HSEMAPHORE) {
	syscall3(engReleaseSemaphore, 1,
		uintptr(hsem),
		0,
		0)
}

func EngStretchBlt(psoDest *SURFOBJ, psoSrc *SURFOBJ, psoMask *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, pca *COLORADJUSTMENT, pptlHTOrg *POINTL, prclDest *RECTL, prclSrc *RECTL, pptlMask *POINTL, iMode ULONG) bool {
	ret1 := syscall12(engStretchBlt, 11,
		uintptr(unsafe.Pointer(psoDest)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(psoMask)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(pca)),
		uintptr(unsafe.Pointer(pptlHTOrg)),
		uintptr(unsafe.Pointer(prclDest)),
		uintptr(unsafe.Pointer(prclSrc)),
		uintptr(unsafe.Pointer(pptlMask)),
		uintptr(iMode),
		0)
	return ret1 != 0
}

func EngStretchBltROP(psoDest *SURFOBJ, psoSrc *SURFOBJ, psoMask *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, pca *COLORADJUSTMENT, pptlHTOrg *POINTL, prclDest *RECTL, prclSrc *RECTL, pptlMask *POINTL, iMode ULONG, pbo *BRUSHOBJ, rop4 uint32) bool {
	ret1 := syscall15(engStretchBltROP, 13,
		uintptr(unsafe.Pointer(psoDest)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(psoMask)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(pca)),
		uintptr(unsafe.Pointer(pptlHTOrg)),
		uintptr(unsafe.Pointer(prclDest)),
		uintptr(unsafe.Pointer(prclSrc)),
		uintptr(unsafe.Pointer(pptlMask)),
		uintptr(iMode),
		uintptr(unsafe.Pointer(pbo)),
		uintptr(rop4),
		0,
		0)
	return ret1 != 0
}

func EngStrokeAndFillPath(pso *SURFOBJ, ppo *PATHOBJ, pco *CLIPOBJ, pxo *XFORMOBJ, pboStroke *BRUSHOBJ, plineattrs *LINEATTRS, pboFill *BRUSHOBJ, pptlBrushOrg *POINTL, mixFill MIX, flOptions FLONG) bool {
	ret1 := syscall12(engStrokeAndFillPath, 10,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(ppo)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxo)),
		uintptr(unsafe.Pointer(pboStroke)),
		uintptr(unsafe.Pointer(plineattrs)),
		uintptr(unsafe.Pointer(pboFill)),
		uintptr(unsafe.Pointer(pptlBrushOrg)),
		uintptr(mixFill),
		uintptr(flOptions),
		0,
		0)
	return ret1 != 0
}

func EngStrokePath(pso *SURFOBJ, ppo *PATHOBJ, pco *CLIPOBJ, pxo *XFORMOBJ, pbo *BRUSHOBJ, pptlBrushOrg *POINTL, plineattrs *LINEATTRS, mix MIX) bool {
	ret1 := syscall9(engStrokePath, 8,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(ppo)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxo)),
		uintptr(unsafe.Pointer(pbo)),
		uintptr(unsafe.Pointer(pptlBrushOrg)),
		uintptr(unsafe.Pointer(plineattrs)),
		uintptr(mix),
		0)
	return ret1 != 0
}

func EngTextOut(pso *SURFOBJ, pstro *STROBJ, pfo *FONTOBJ, pco *CLIPOBJ, prclExtra *RECTL, prclOpaque *RECTL, pboFore *BRUSHOBJ, pboOpaque *BRUSHOBJ, pptlOrg *POINTL, mix MIX) bool {
	ret1 := syscall12(engTextOut, 10,
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(pstro)),
		uintptr(unsafe.Pointer(pfo)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(prclExtra)),
		uintptr(unsafe.Pointer(prclOpaque)),
		uintptr(unsafe.Pointer(pboFore)),
		uintptr(unsafe.Pointer(pboOpaque)),
		uintptr(unsafe.Pointer(pptlOrg)),
		uintptr(mix),
		0,
		0)
	return ret1 != 0
}

func EngTransparentBlt(psoDst *SURFOBJ, psoSrc *SURFOBJ, pco *CLIPOBJ, pxlo *XLATEOBJ, prclDst *RECTL, prclSrc *RECTL, iTransColor ULONG, ulReserved ULONG) bool {
	ret1 := syscall9(engTransparentBlt, 8,
		uintptr(unsafe.Pointer(psoDst)),
		uintptr(unsafe.Pointer(psoSrc)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(unsafe.Pointer(prclDst)),
		uintptr(unsafe.Pointer(prclSrc)),
		uintptr(iTransColor),
		uintptr(ulReserved),
		0)
	return ret1 != 0
}

func EngUnicodeToMultiByteN(multiByteString PCHAR, maxBytesInMultiByteString ULONG, bytesInMultiByteString *uint32, unicodeString PWSTR, bytesInUnicodeString ULONG) {
	syscall6(engUnicodeToMultiByteN, 5,
		uintptr(unsafe.Pointer(multiByteString)),
		uintptr(maxBytesInMultiByteString),
		uintptr(unsafe.Pointer(bytesInMultiByteString)),
		uintptr(unsafe.Pointer(unicodeString)),
		uintptr(bytesInUnicodeString),
		0)
}

func EngUnlockSurface(pso *SURFOBJ) {
	syscall3(engUnlockSurface, 1,
		uintptr(unsafe.Pointer(pso)),
		0,
		0)
}

func EngWideCharToMultiByte(codePage uint32, wideCharString LPWSTR, bytesInWideCharString int32, multiByteString LPSTR, bytesInMultiByteString int32) int32 {
	ret1 := syscall6(engWideCharToMultiByte, 5,
		uintptr(codePage),
		uintptr(unsafe.Pointer(wideCharString)),
		uintptr(bytesInWideCharString),
		uintptr(unsafe.Pointer(multiByteString)),
		uintptr(bytesInMultiByteString),
		0)
	return int32(ret1)
}

func EnumEnhMetaFile(hdc HDC, hmf HENHMETAFILE, lpProc ENHMFENUMPROC, lpParam LPVOID, lpRect /*const*/ *RECT) bool {
	lpProcCallback := syscall.NewCallback(func(hdcRawArg HDC, lphtRawArg *HANDLETABLE, lpmrRawArg /*const*/ *ENHMETARECORD, nHandlesRawArg int32, dataRawArg LPARAM) uintptr {
		ret := lpProc(hdcRawArg, lphtRawArg, lpmrRawArg, nHandlesRawArg, dataRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumEnhMetaFile, 5,
		uintptr(hdc),
		uintptr(hmf),
		lpProcCallback,
		uintptr(unsafe.Pointer(lpParam)),
		uintptr(unsafe.Pointer(lpRect)),
		0)
	return ret1 != 0
}

func EnumFontFamiliesEx(hdc HDC, lpLogfont LPLOGFONT, lpProc FONTENUMPROC, lParam LPARAM, dwFlags uint32) int32 {
	lpProcCallback := syscall.NewCallback(func(unnamed0RawArg /*const*/ *LOGFONT, unnamed1RawArg /*const*/ *TEXTMETRIC, unnamed2RawArg DWORD, unnamed3RawArg LPARAM) uintptr {
		ret := lpProc(unnamed0RawArg, unnamed1RawArg, unnamed2RawArg, unnamed3RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumFontFamiliesEx, 5,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpLogfont)),
		lpProcCallback,
		uintptr(lParam),
		uintptr(dwFlags),
		0)
	return int32(ret1)
}

func EnumFontFamilies(hdc HDC, lpLogfont string, lpProc FONTENUMPROC, lParam LPARAM) int32 {
	lpLogfontStr := unicode16FromString(lpLogfont)
	lpProcCallback := syscall.NewCallback(func(unnamed0RawArg /*const*/ *LOGFONT, unnamed1RawArg /*const*/ *TEXTMETRIC, unnamed2RawArg DWORD, unnamed3RawArg LPARAM) uintptr {
		ret := lpProc(unnamed0RawArg, unnamed1RawArg, unnamed2RawArg, unnamed3RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumFontFamilies, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpLogfontStr[0])),
		lpProcCallback,
		uintptr(lParam),
		0,
		0)
	return int32(ret1)
}

func EnumFonts(hdc HDC, lpLogfont string, lpProc FONTENUMPROC, lParam LPARAM) int32 {
	lpLogfontStr := unicode16FromString(lpLogfont)
	lpProcCallback := syscall.NewCallback(func(unnamed0RawArg /*const*/ *LOGFONT, unnamed1RawArg /*const*/ *TEXTMETRIC, unnamed2RawArg DWORD, unnamed3RawArg LPARAM) uintptr {
		ret := lpProc(unnamed0RawArg, unnamed1RawArg, unnamed2RawArg, unnamed3RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumFonts, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpLogfontStr[0])),
		lpProcCallback,
		uintptr(lParam),
		0,
		0)
	return int32(ret1)
}

func EnumICMProfiles(hdc HDC, lpProc ICMENUMPROC, lParam LPARAM) int32 {
	lpProcCallback := syscall.NewCallback(func(unnamed0RawArg LPWSTR, unnamed1RawArg LPARAM) uintptr {
		ret := lpProc(unnamed0RawArg, unnamed1RawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(enumICMProfiles, 3,
		uintptr(hdc),
		lpProcCallback,
		uintptr(lParam))
	return int32(ret1)
}

func EnumMetaFile(hdc HDC, hmf HMETAFILE, lpProc MFENUMPROC, lParam LPARAM) bool {
	lpProcCallback := syscall.NewCallback(func(hdcRawArg HDC, lphtRawArg *HANDLETABLE, lpMRRawArg *METARECORD, nObjRawArg int32, paramRawArg LPARAM) uintptr {
		ret := lpProc(hdcRawArg, lphtRawArg, lpMRRawArg, nObjRawArg, paramRawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumMetaFile, 4,
		uintptr(hdc),
		uintptr(hmf),
		lpProcCallback,
		uintptr(lParam),
		0,
		0)
	return ret1 != 0
}

func EnumObjects(hdc HDC, nType int32, lpFunc GOBJENUMPROC, lParam LPARAM) int32 {
	lpFuncCallback := syscall.NewCallback(func(unnamed0RawArg LPVOID, unnamed1RawArg LPARAM) uintptr {
		ret := lpFunc(unnamed0RawArg, unnamed1RawArg)
		return uintptr(ret)
	})
	ret1 := syscall6(enumObjects, 4,
		uintptr(hdc),
		uintptr(nType),
		lpFuncCallback,
		uintptr(lParam),
		0,
		0)
	return int32(ret1)
}

func EqualRgn(hrgn1 HRGN, hrgn2 HRGN) bool {
	ret1 := syscall3(equalRgn, 2,
		uintptr(hrgn1),
		uintptr(hrgn2),
		0)
	return ret1 != 0
}

func Escape(hdc HDC, iEscape int32, cjIn int32, pvIn /*const*/ LPCSTR, pvOut LPVOID) int32 {
	ret1 := syscall6(escape, 5,
		uintptr(hdc),
		uintptr(iEscape),
		uintptr(cjIn),
		uintptr(unsafe.Pointer(pvIn)),
		uintptr(unsafe.Pointer(pvOut)),
		0)
	return int32(ret1)
}

func ExcludeClipRect(hdc HDC, left int32, top int32, right int32, bottom int32) int32 {
	ret1 := syscall6(excludeClipRect, 5,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		0)
	return int32(ret1)
}

func ExtCreatePen(iPenStyle uint32, cWidth uint32, plbrush /*const*/ *LOGBRUSH, cStyle uint32, pstyle /*const*/ *uint32) HPEN {
	ret1 := syscall6(extCreatePen, 5,
		uintptr(iPenStyle),
		uintptr(cWidth),
		uintptr(unsafe.Pointer(plbrush)),
		uintptr(cStyle),
		uintptr(unsafe.Pointer(pstyle)),
		0)
	return HPEN(ret1)
}

func ExtCreateRegion(lpx /*const*/ *XFORM, nCount uint32, lpData /*const*/ *RGNDATA) HRGN {
	ret1 := syscall3(extCreateRegion, 3,
		uintptr(unsafe.Pointer(lpx)),
		uintptr(nCount),
		uintptr(unsafe.Pointer(lpData)))
	return HRGN(ret1)
}

func ExtEscape(hdc HDC, iEscape int32, cjInput int32, lpInData /*const*/ LPCSTR, cjOutput int32, lpOutData LPSTR) int32 {
	ret1 := syscall6(extEscape, 6,
		uintptr(hdc),
		uintptr(iEscape),
		uintptr(cjInput),
		uintptr(unsafe.Pointer(lpInData)),
		uintptr(cjOutput),
		uintptr(unsafe.Pointer(lpOutData)))
	return int32(ret1)
}

func ExtFloodFill(hdc HDC, x int32, y int32, color COLORREF, aType uint32) bool {
	ret1 := syscall6(extFloodFill, 5,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(color),
		uintptr(aType),
		0)
	return ret1 != 0
}

func ExtSelectClipRgn(hdc HDC, hrgn HRGN, mode int32) int32 {
	ret1 := syscall3(extSelectClipRgn, 3,
		uintptr(hdc),
		uintptr(hrgn),
		uintptr(mode))
	return int32(ret1)
}

func ExtTextOut(hdc HDC, x int32, y int32, options uint32, lprect /*const*/ *RECT, lpString string, c uint32, lpDx /*const*/ *int32) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall9(extTextOut, 8,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(options),
		uintptr(unsafe.Pointer(lprect)),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(c),
		uintptr(unsafe.Pointer(lpDx)),
		0)
	return ret1 != 0
}

func FONTOBJ_cGetAllGlyphHandles(pfo *FONTOBJ, phg *HGLYPH) ULONG {
	ret1 := syscall3(fONTOBJ_cGetAllGlyphHandles, 2,
		uintptr(unsafe.Pointer(pfo)),
		uintptr(unsafe.Pointer(phg)),
		0)
	return ULONG(ret1)
}

func FONTOBJ_cGetGlyphs(pfo *FONTOBJ, iMode ULONG, cGlyph ULONG, phg *HGLYPH, ppvGlyph *PVOID) ULONG {
	ret1 := syscall6(fONTOBJ_cGetGlyphs, 5,
		uintptr(unsafe.Pointer(pfo)),
		uintptr(iMode),
		uintptr(cGlyph),
		uintptr(unsafe.Pointer(phg)),
		uintptr(unsafe.Pointer(ppvGlyph)),
		0)
	return ULONG(ret1)
}

func FONTOBJ_pQueryGlyphAttrs(pfo *FONTOBJ, iMode ULONG) PFD_GLYPHATTR {
	ret1 := syscall3(fONTOBJ_pQueryGlyphAttrs, 2,
		uintptr(unsafe.Pointer(pfo)),
		uintptr(iMode),
		0)
	return (PFD_GLYPHATTR)(unsafe.Pointer(ret1))
}

func FONTOBJ_pvTrueTypeFontFile(pfo *FONTOBJ, pcjFile *ULONG) uintptr {
	ret1 := syscall3(fONTOBJ_pvTrueTypeFontFile, 2,
		uintptr(unsafe.Pointer(pfo)),
		uintptr(unsafe.Pointer(pcjFile)),
		0)
	return (uintptr)(unsafe.Pointer(ret1))
}

func FONTOBJ_vGetInfo(pfo *FONTOBJ, cjSize ULONG, pfi *FONTINFO) {
	syscall3(fONTOBJ_vGetInfo, 3,
		uintptr(unsafe.Pointer(pfo)),
		uintptr(cjSize),
		uintptr(unsafe.Pointer(pfi)))
}

func FillPath(hdc HDC) bool {
	ret1 := syscall3(fillPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func FillRgn(hdc HDC, hrgn HRGN, hbr HBRUSH) bool {
	ret1 := syscall3(fillRgn, 3,
		uintptr(hdc),
		uintptr(hrgn),
		uintptr(hbr))
	return ret1 != 0
}

func FixBrushOrgEx(hdc HDC, x int32, y int32, ptl *POINT) bool {
	ret1 := syscall6(fixBrushOrgEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(ptl)),
		0,
		0)
	return ret1 != 0
}

func FlattenPath(hdc HDC) bool {
	ret1 := syscall3(flattenPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func FloodFill(hdc HDC, x int32, y int32, color COLORREF) bool {
	ret1 := syscall6(floodFill, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(color),
		0,
		0)
	return ret1 != 0
}

func FrameRgn(hdc HDC, hrgn HRGN, hbr HBRUSH, w int32, h int32) bool {
	ret1 := syscall6(frameRgn, 5,
		uintptr(hdc),
		uintptr(hrgn),
		uintptr(hbr),
		uintptr(w),
		uintptr(h),
		0)
	return ret1 != 0
}

func GdiAlphaBlend(hdcDest HDC, xoriginDest int32, yoriginDest int32, wDest int32, hDest int32, hdcSrc HDC, xoriginSrc int32, yoriginSrc int32, wSrc int32, hSrc int32, ftn BLENDFUNCTION) bool {
	ret1 := syscall12(gdiAlphaBlend, 11,
		uintptr(hdcDest),
		uintptr(xoriginDest),
		uintptr(yoriginDest),
		uintptr(wDest),
		uintptr(hDest),
		uintptr(hdcSrc),
		uintptr(xoriginSrc),
		uintptr(yoriginSrc),
		uintptr(wSrc),
		uintptr(hSrc),
		getUintptrFromBLENDFUNCTION(ftn),
		0)
	return ret1 != 0
}

func GdiComment(hdc HDC, nSize uint32, lpData /*const*/ *byte) bool {
	ret1 := syscall3(gdiComment, 3,
		uintptr(hdc),
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpData)))
	return ret1 != 0
}

func GdiFlush() bool {
	ret1 := syscall3(gdiFlush, 0,
		0,
		0,
		0)
	return ret1 != 0
}

func GdiGetBatchLimit() uint32 {
	ret1 := syscall3(gdiGetBatchLimit, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GdiGradientFill(hdc HDC, pVertex PTRIVERTEX, nVertex ULONG, pMesh uintptr, nMesh ULONG, ulMode ULONG) bool {
	ret1 := syscall6(gdiGradientFill, 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pVertex)),
		uintptr(nVertex),
		pMesh,
		uintptr(nMesh),
		uintptr(ulMode))
	return ret1 != 0
}

func GdiSetBatchLimit(dw uint32) uint32 {
	ret1 := syscall3(gdiSetBatchLimit, 1,
		uintptr(dw),
		0,
		0)
	return uint32(ret1)
}

func GdiTransparentBlt(hdcDest HDC, xoriginDest int32, yoriginDest int32, wDest int32, hDest int32, hdcSrc HDC, xoriginSrc int32, yoriginSrc int32, wSrc int32, hSrc int32, crTransparent uint32) bool {
	ret1 := syscall12(gdiTransparentBlt, 11,
		uintptr(hdcDest),
		uintptr(xoriginDest),
		uintptr(yoriginDest),
		uintptr(wDest),
		uintptr(hDest),
		uintptr(hdcSrc),
		uintptr(xoriginSrc),
		uintptr(yoriginSrc),
		uintptr(wSrc),
		uintptr(hSrc),
		uintptr(crTransparent),
		0)
	return ret1 != 0
}

func GetArcDirection(hdc HDC) int32 {
	ret1 := syscall3(getArcDirection, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetAspectRatioFilterEx(hdc HDC, lpsize *SIZE) bool {
	ret1 := syscall3(getAspectRatioFilterEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpsize)),
		0)
	return ret1 != 0
}

func GetBitmapBits(hbit HBITMAP, cb LONG, lpvBits LPVOID) LONG {
	ret1 := syscall3(getBitmapBits, 3,
		uintptr(hbit),
		uintptr(cb),
		uintptr(unsafe.Pointer(lpvBits)))
	return LONG(ret1)
}

func GetBitmapDimensionEx(hbit HBITMAP, lpsize *SIZE) bool {
	ret1 := syscall3(getBitmapDimensionEx, 2,
		uintptr(hbit),
		uintptr(unsafe.Pointer(lpsize)),
		0)
	return ret1 != 0
}

func GetBkColor(hdc HDC) COLORREF {
	ret1 := syscall3(getBkColor, 1,
		uintptr(hdc),
		0,
		0)
	return COLORREF(ret1)
}

func GetBkMode(hdc HDC) int32 {
	ret1 := syscall3(getBkMode, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetBoundsRect(hdc HDC, lprect *RECT, flags uint32) uint32 {
	ret1 := syscall3(getBoundsRect, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lprect)),
		uintptr(flags))
	return uint32(ret1)
}

func GetBrushOrgEx(hdc HDC, lppt *POINT) bool {
	ret1 := syscall3(getBrushOrgEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppt)),
		0)
	return ret1 != 0
}

func GetCharABCWidthsFloat(hdc HDC, iFirst uint32, iLast uint32, lpABC *ABCFLOAT) bool {
	ret1 := syscall6(getCharABCWidthsFloat, 4,
		uintptr(hdc),
		uintptr(iFirst),
		uintptr(iLast),
		uintptr(unsafe.Pointer(lpABC)),
		0,
		0)
	return ret1 != 0
}

func GetCharABCWidthsI(hdc HDC, giFirst uint32, cgi uint32, pgi *uint16, pabc *ABC) bool {
	ret1 := syscall6(getCharABCWidthsI, 5,
		uintptr(hdc),
		uintptr(giFirst),
		uintptr(cgi),
		uintptr(unsafe.Pointer(pgi)),
		uintptr(unsafe.Pointer(pabc)),
		0)
	return ret1 != 0
}

func GetCharABCWidths(hdc HDC, wFirst uint32, wLast uint32, lpABC *ABC) bool {
	ret1 := syscall6(getCharABCWidths, 4,
		uintptr(hdc),
		uintptr(wFirst),
		uintptr(wLast),
		uintptr(unsafe.Pointer(lpABC)),
		0,
		0)
	return ret1 != 0
}

func GetCharWidth32(hdc HDC, iFirst uint32, iLast uint32, lpBuffer *int32) bool {
	ret1 := syscall6(getCharWidth32, 4,
		uintptr(hdc),
		uintptr(iFirst),
		uintptr(iLast),
		uintptr(unsafe.Pointer(lpBuffer)),
		0,
		0)
	return ret1 != 0
}

func GetCharWidthFloat(hdc HDC, iFirst uint32, iLast uint32, lpBuffer *float32) bool {
	ret1 := syscall6(getCharWidthFloat, 4,
		uintptr(hdc),
		uintptr(iFirst),
		uintptr(iLast),
		uintptr(unsafe.Pointer(lpBuffer)),
		0,
		0)
	return ret1 != 0
}

func GetCharWidthI(hdc HDC, giFirst uint32, cgi uint32, pgi *uint16, piWidths *int32) bool {
	ret1 := syscall6(getCharWidthI, 5,
		uintptr(hdc),
		uintptr(giFirst),
		uintptr(cgi),
		uintptr(unsafe.Pointer(pgi)),
		uintptr(unsafe.Pointer(piWidths)),
		0)
	return ret1 != 0
}

func GetCharWidth(hdc HDC, iFirst uint32, iLast uint32, lpBuffer *int32) bool {
	ret1 := syscall6(getCharWidth, 4,
		uintptr(hdc),
		uintptr(iFirst),
		uintptr(iLast),
		uintptr(unsafe.Pointer(lpBuffer)),
		0,
		0)
	return ret1 != 0
}

func GetCharacterPlacement(hdc HDC, lpString string, nCount int32, nMexExtent int32, lpResults LPGCP_RESULTS, dwFlags uint32) uint32 {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall6(getCharacterPlacement, 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(nCount),
		uintptr(nMexExtent),
		uintptr(unsafe.Pointer(lpResults)),
		uintptr(dwFlags))
	return uint32(ret1)
}

func GetClipBox(hdc HDC, lprect *RECT) int32 {
	ret1 := syscall3(getClipBox, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lprect)),
		0)
	return int32(ret1)
}

func GetClipRgn(hdc HDC, hrgn HRGN) int32 {
	ret1 := syscall3(getClipRgn, 2,
		uintptr(hdc),
		uintptr(hrgn),
		0)
	return int32(ret1)
}

func GetColorAdjustment(hdc HDC, lpca *COLORADJUSTMENT) bool {
	ret1 := syscall3(getColorAdjustment, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpca)),
		0)
	return ret1 != 0
}

func GetColorSpace(hdc HDC) HCOLORSPACE {
	ret1 := syscall3(getColorSpace, 1,
		uintptr(hdc),
		0,
		0)
	return HCOLORSPACE(ret1)
}

func GetCurrentObject(hdc HDC, aType uint32) HGDIOBJ {
	ret1 := syscall3(getCurrentObject, 2,
		uintptr(hdc),
		uintptr(aType),
		0)
	return HGDIOBJ(ret1)
}

func GetCurrentPositionEx(hdc HDC, lppt *POINT) bool {
	ret1 := syscall3(getCurrentPositionEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppt)),
		0)
	return ret1 != 0
}

func GetDCBrushColor(hdc HDC) COLORREF {
	ret1 := syscall3(getDCBrushColor, 1,
		uintptr(hdc),
		0,
		0)
	return COLORREF(ret1)
}

func GetDCOrgEx(hdc HDC, lppt *POINT) bool {
	ret1 := syscall3(getDCOrgEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppt)),
		0)
	return ret1 != 0
}

func GetDCPenColor(hdc HDC) COLORREF {
	ret1 := syscall3(getDCPenColor, 1,
		uintptr(hdc),
		0,
		0)
	return COLORREF(ret1)
}

func GetDIBColorTable(hdc HDC, iStart uint32, cEntries uint32, prgbq *RGBQUAD) uint32 {
	ret1 := syscall6(getDIBColorTable, 4,
		uintptr(hdc),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(prgbq)),
		0,
		0)
	return uint32(ret1)
}

func GetDIBits(hdc HDC, hbm HBITMAP, start uint32, cLines uint32, lpvBits LPVOID, lpbmi *BITMAPINFO, usage uint32) int32 {
	ret1 := syscall9(getDIBits, 7,
		uintptr(hdc),
		uintptr(hbm),
		uintptr(start),
		uintptr(cLines),
		uintptr(unsafe.Pointer(lpvBits)),
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(usage),
		0,
		0)
	return int32(ret1)
}

func GetDeviceCaps(hdc HDC, index int32) int32 {
	ret1 := syscall3(getDeviceCaps, 2,
		uintptr(hdc),
		uintptr(index),
		0)
	return int32(ret1)
}

func GetDeviceGammaRamp(hdc HDC, lpRamp LPVOID) bool {
	ret1 := syscall3(getDeviceGammaRamp, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpRamp)),
		0)
	return ret1 != 0
}

func GetEnhMetaFileBits(hEMF HENHMETAFILE, nSize uint32, lpData *byte) uint32 {
	ret1 := syscall3(getEnhMetaFileBits, 3,
		uintptr(hEMF),
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpData)))
	return uint32(ret1)
}

func GetEnhMetaFileDescription(hemf HENHMETAFILE, cchBuffer uint32, lpDescription LPWSTR) uint32 {
	ret1 := syscall3(getEnhMetaFileDescription, 3,
		uintptr(hemf),
		uintptr(cchBuffer),
		uintptr(unsafe.Pointer(lpDescription)))
	return uint32(ret1)
}

func GetEnhMetaFileHeader(hemf HENHMETAFILE, nSize uint32, lpEnhMetaHeader *ENHMETAHEADER) uint32 {
	ret1 := syscall3(getEnhMetaFileHeader, 3,
		uintptr(hemf),
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpEnhMetaHeader)))
	return uint32(ret1)
}

func GetEnhMetaFilePaletteEntries(hemf HENHMETAFILE, nNumEntries uint32, lpPaletteEntries *PALETTEENTRY) uint32 {
	ret1 := syscall3(getEnhMetaFilePaletteEntries, 3,
		uintptr(hemf),
		uintptr(nNumEntries),
		uintptr(unsafe.Pointer(lpPaletteEntries)))
	return uint32(ret1)
}

func GetEnhMetaFilePixelFormat(hemf HENHMETAFILE, cbBuffer uint32, ppfd *PIXELFORMATDESCRIPTOR) uint32 {
	ret1 := syscall3(getEnhMetaFilePixelFormat, 3,
		uintptr(hemf),
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(ppfd)))
	return uint32(ret1)
}

func GetEnhMetaFile(lpName string) HENHMETAFILE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(getEnhMetaFile, 1,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		0,
		0)
	return HENHMETAFILE(ret1)
}

func GetFontData(hdc HDC, dwTable uint32, dwOffset uint32, pvBuffer uintptr, cjBuffer uint32) uint32 {
	ret1 := syscall6(getFontData, 5,
		uintptr(hdc),
		uintptr(dwTable),
		uintptr(dwOffset),
		pvBuffer,
		uintptr(cjBuffer),
		0)
	return uint32(ret1)
}

func GetFontLanguageInfo(hdc HDC) uint32 {
	ret1 := syscall3(getFontLanguageInfo, 1,
		uintptr(hdc),
		0,
		0)
	return uint32(ret1)
}

func GetFontUnicodeRanges(hdc HDC, lpgs *GLYPHSET) uint32 {
	ret1 := syscall3(getFontUnicodeRanges, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpgs)),
		0)
	return uint32(ret1)
}

func GetGlyphIndices(hdc HDC, lpstr string, c int32, pgi *uint16, fl uint32) uint32 {
	lpstrStr := unicode16FromString(lpstr)
	ret1 := syscall6(getGlyphIndices, 5,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpstrStr[0])),
		uintptr(c),
		uintptr(unsafe.Pointer(pgi)),
		uintptr(fl),
		0)
	return uint32(ret1)
}

func GetGlyphOutline(hdc HDC, uChar uint32, fuFormat uint32, lpgm *GLYPHMETRICS, cjBuffer uint32, pvBuffer LPVOID, lpmat2 /*const*/ *MAT2) uint32 {
	ret1 := syscall9(getGlyphOutline, 7,
		uintptr(hdc),
		uintptr(uChar),
		uintptr(fuFormat),
		uintptr(unsafe.Pointer(lpgm)),
		uintptr(cjBuffer),
		uintptr(unsafe.Pointer(pvBuffer)),
		uintptr(unsafe.Pointer(lpmat2)),
		0,
		0)
	return uint32(ret1)
}

func GetGraphicsMode(hdc HDC) int32 {
	ret1 := syscall3(getGraphicsMode, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetICMProfile(hdc HDC, pBufSize *uint32, pszFilename LPWSTR) bool {
	ret1 := syscall3(getICMProfile, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pBufSize)),
		uintptr(unsafe.Pointer(pszFilename)))
	return ret1 != 0
}

func GetKerningPairs(hdc HDC, nPairs uint32, lpKernPair *KERNINGPAIR) uint32 {
	ret1 := syscall3(getKerningPairs, 3,
		uintptr(hdc),
		uintptr(nPairs),
		uintptr(unsafe.Pointer(lpKernPair)))
	return uint32(ret1)
}

func GetLayout(hdc HDC) uint32 {
	ret1 := syscall3(getLayout, 1,
		uintptr(hdc),
		0,
		0)
	return uint32(ret1)
}

func GetLogColorSpace(hColorSpace HCOLORSPACE, lpBuffer LPLOGCOLORSPACE, nSize uint32) bool {
	ret1 := syscall3(getLogColorSpace, 3,
		uintptr(hColorSpace),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize))
	return ret1 != 0
}

func GetMapMode(hdc HDC) int32 {
	ret1 := syscall3(getMapMode, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetMetaFileBitsEx(hMF HMETAFILE, cbBuffer uint32, lpData LPVOID) uint32 {
	ret1 := syscall3(getMetaFileBitsEx, 3,
		uintptr(hMF),
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(lpData)))
	return uint32(ret1)
}

func GetMetaFile(lpName string) HMETAFILE {
	lpNameStr := unicode16FromString(lpName)
	ret1 := syscall3(getMetaFile, 1,
		uintptr(unsafe.Pointer(&lpNameStr[0])),
		0,
		0)
	return HMETAFILE(ret1)
}

func GetMetaRgn(hdc HDC, hrgn HRGN) int32 {
	ret1 := syscall3(getMetaRgn, 2,
		uintptr(hdc),
		uintptr(hrgn),
		0)
	return int32(ret1)
}

func GetMiterLimit(hdc HDC, plimit *float32) bool {
	ret1 := syscall3(getMiterLimit, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(plimit)),
		0)
	return ret1 != 0
}

func GetNearestColor(hdc HDC, color COLORREF) COLORREF {
	ret1 := syscall3(getNearestColor, 2,
		uintptr(hdc),
		uintptr(color),
		0)
	return COLORREF(ret1)
}

func GetNearestPaletteIndex(h HPALETTE, color COLORREF) uint32 {
	ret1 := syscall3(getNearestPaletteIndex, 2,
		uintptr(h),
		uintptr(color),
		0)
	return uint32(ret1)
}

func GetObjectType(h HGDIOBJ) uint32 {
	ret1 := syscall3(getObjectType, 1,
		uintptr(h),
		0,
		0)
	return uint32(ret1)
}

func GetObject(h HANDLE, c int32, pv LPVOID) int32 {
	ret1 := syscall3(getObject, 3,
		uintptr(h),
		uintptr(c),
		uintptr(unsafe.Pointer(pv)))
	return int32(ret1)
}

func GetOutlineTextMetrics(hdc HDC, cjCopy uint32, potm LPOUTLINETEXTMETRIC) uint32 {
	ret1 := syscall3(getOutlineTextMetrics, 3,
		uintptr(hdc),
		uintptr(cjCopy),
		uintptr(unsafe.Pointer(potm)))
	return uint32(ret1)
}

func GetPaletteEntries(hpal HPALETTE, iStart uint32, cEntries uint32, pPalEntries *PALETTEENTRY) uint32 {
	ret1 := syscall6(getPaletteEntries, 4,
		uintptr(hpal),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pPalEntries)),
		0,
		0)
	return uint32(ret1)
}

func GetPath(hdc HDC, apt *POINT, aj *byte, cpt int32) int32 {
	ret1 := syscall6(getPath, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(unsafe.Pointer(aj)),
		uintptr(cpt),
		0,
		0)
	return int32(ret1)
}

func GetPixel(hdc HDC, x int32, y int32) COLORREF {
	ret1 := syscall3(getPixel, 3,
		uintptr(hdc),
		uintptr(x),
		uintptr(y))
	return COLORREF(ret1)
}

func GetPixelFormat(hdc HDC) int32 {
	ret1 := syscall3(getPixelFormat, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetPolyFillMode(hdc HDC) int32 {
	ret1 := syscall3(getPolyFillMode, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetROP2(hdc HDC) int32 {
	ret1 := syscall3(getROP2, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetRandomRgn(hdc HDC, hrgn HRGN, i int32) int32 {
	ret1 := syscall3(getRandomRgn, 3,
		uintptr(hdc),
		uintptr(hrgn),
		uintptr(i))
	return int32(ret1)
}

func GetRasterizerCaps(lpraststat LPRASTERIZER_STATUS, cjBytes uint32) bool {
	ret1 := syscall3(getRasterizerCaps, 2,
		uintptr(unsafe.Pointer(lpraststat)),
		uintptr(cjBytes),
		0)
	return ret1 != 0
}

func GetRegionData(hrgn HRGN, nCount uint32, lpRgnData *RGNDATA) uint32 {
	ret1 := syscall3(getRegionData, 3,
		uintptr(hrgn),
		uintptr(nCount),
		uintptr(unsafe.Pointer(lpRgnData)))
	return uint32(ret1)
}

func GetRgnBox(hrgn HRGN, lprc *RECT) int32 {
	ret1 := syscall3(getRgnBox, 2,
		uintptr(hrgn),
		uintptr(unsafe.Pointer(lprc)),
		0)
	return int32(ret1)
}

func GetStockObject(i int32) HGDIOBJ {
	ret1 := syscall3(getStockObject, 1,
		uintptr(i),
		0,
		0)
	return HGDIOBJ(ret1)
}

func GetStretchBltMode(hdc HDC) int32 {
	ret1 := syscall3(getStretchBltMode, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetSystemPaletteEntries(hdc HDC, iStart uint32, cEntries uint32, pPalEntries *PALETTEENTRY) uint32 {
	ret1 := syscall6(getSystemPaletteEntries, 4,
		uintptr(hdc),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pPalEntries)),
		0,
		0)
	return uint32(ret1)
}

func GetSystemPaletteUse(hdc HDC) uint32 {
	ret1 := syscall3(getSystemPaletteUse, 1,
		uintptr(hdc),
		0,
		0)
	return uint32(ret1)
}

func GetTextAlign(hdc HDC) uint32 {
	ret1 := syscall3(getTextAlign, 1,
		uintptr(hdc),
		0,
		0)
	return uint32(ret1)
}

func GetTextCharacterExtra(hdc HDC) int32 {
	ret1 := syscall3(getTextCharacterExtra, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetTextCharset(hdc HDC) int32 {
	ret1 := syscall3(getTextCharset, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func GetTextCharsetInfo(hdc HDC, lpSig *FONTSIGNATURE, dwFlags uint32) int32 {
	ret1 := syscall3(getTextCharsetInfo, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpSig)),
		uintptr(dwFlags))
	return int32(ret1)
}

func GetTextColor(hdc HDC) COLORREF {
	ret1 := syscall3(getTextColor, 1,
		uintptr(hdc),
		0,
		0)
	return COLORREF(ret1)
}

func GetTextExtentExPointI(hdc HDC, lpwszString *uint16, cwchString int32, nMaxExtent int32, lpnFit *int32, lpnDx *int32, lpSize *SIZE) bool {
	ret1 := syscall9(getTextExtentExPointI, 7,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpwszString)),
		uintptr(cwchString),
		uintptr(nMaxExtent),
		uintptr(unsafe.Pointer(lpnFit)),
		uintptr(unsafe.Pointer(lpnDx)),
		uintptr(unsafe.Pointer(lpSize)),
		0,
		0)
	return ret1 != 0
}

func GetTextExtentExPoint(hdc HDC, lpszString string, cchString int32, nMaxExtent int32, lpnFit *int32, lpnDx *int32, lpSize *SIZE) bool {
	lpszStringStr := unicode16FromString(lpszString)
	ret1 := syscall9(getTextExtentExPoint, 7,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpszStringStr[0])),
		uintptr(cchString),
		uintptr(nMaxExtent),
		uintptr(unsafe.Pointer(lpnFit)),
		uintptr(unsafe.Pointer(lpnDx)),
		uintptr(unsafe.Pointer(lpSize)),
		0,
		0)
	return ret1 != 0
}

func GetTextExtentPoint32(hdc HDC, lpString string, c int32, psizl *SIZE) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall6(getTextExtentPoint32, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(c),
		uintptr(unsafe.Pointer(psizl)),
		0,
		0)
	return ret1 != 0
}

func GetTextExtentPointI(hdc HDC, pgiIn *uint16, cgi int32, psize *SIZE) bool {
	ret1 := syscall6(getTextExtentPointI, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pgiIn)),
		uintptr(cgi),
		uintptr(unsafe.Pointer(psize)),
		0,
		0)
	return ret1 != 0
}

func GetTextExtentPoint(hdc HDC, lpString string, c int32, lpsz *SIZE) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall6(getTextExtentPoint, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(c),
		uintptr(unsafe.Pointer(lpsz)),
		0,
		0)
	return ret1 != 0
}

func GetTextFace(hdc HDC, c int32, lpName LPWSTR) int32 {
	ret1 := syscall3(getTextFace, 3,
		uintptr(hdc),
		uintptr(c),
		uintptr(unsafe.Pointer(lpName)))
	return int32(ret1)
}

func GetTextMetrics(hdc HDC, lptm LPTEXTMETRIC) bool {
	ret1 := syscall3(getTextMetrics, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lptm)),
		0)
	return ret1 != 0
}

func GetViewportExtEx(hdc HDC, lpsize *SIZE) bool {
	ret1 := syscall3(getViewportExtEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpsize)),
		0)
	return ret1 != 0
}

func GetViewportOrgEx(hdc HDC, lppoint *POINT) bool {
	ret1 := syscall3(getViewportOrgEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppoint)),
		0)
	return ret1 != 0
}

func GetWinMetaFileBits(hemf HENHMETAFILE, cbData16 uint32, pData16 *byte, iMapMode int32, hdcRef HDC) uint32 {
	ret1 := syscall6(getWinMetaFileBits, 5,
		uintptr(hemf),
		uintptr(cbData16),
		uintptr(unsafe.Pointer(pData16)),
		uintptr(iMapMode),
		uintptr(hdcRef),
		0)
	return uint32(ret1)
}

func GetWindowExtEx(hdc HDC, lpsize *SIZE) bool {
	ret1 := syscall3(getWindowExtEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpsize)),
		0)
	return ret1 != 0
}

func GetWindowOrgEx(hdc HDC, lppoint *POINT) bool {
	ret1 := syscall3(getWindowOrgEx, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppoint)),
		0)
	return ret1 != 0
}

func GetWorldTransform(hdc HDC, lpxf *XFORM) bool {
	ret1 := syscall3(getWorldTransform, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpxf)),
		0)
	return ret1 != 0
}

func HT_Get8BPPFormatPalette(pPaletteEntry *PALETTEENTRY, redGamma USHORT, greenGamma USHORT, blueGamma USHORT) LONG {
	ret1 := syscall6(hT_Get8BPPFormatPalette, 4,
		uintptr(unsafe.Pointer(pPaletteEntry)),
		uintptr(redGamma),
		uintptr(greenGamma),
		uintptr(blueGamma),
		0,
		0)
	return LONG(ret1)
}

func HT_Get8BPPMaskPalette(pPaletteEntry *PALETTEENTRY, use8BPPMaskPal bool, cMYMask byte, redGamma USHORT, greenGamma USHORT, blueGamma USHORT) LONG {
	ret1 := syscall6(hT_Get8BPPMaskPalette, 6,
		uintptr(unsafe.Pointer(pPaletteEntry)),
		getUintptrFromBool(use8BPPMaskPal),
		uintptr(cMYMask),
		uintptr(redGamma),
		uintptr(greenGamma),
		uintptr(blueGamma))
	return LONG(ret1)
}

func IntersectClipRect(hdc HDC, left int32, top int32, right int32, bottom int32) int32 {
	ret1 := syscall6(intersectClipRect, 5,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		0)
	return int32(ret1)
}

func InvertRgn(hdc HDC, hrgn HRGN) bool {
	ret1 := syscall3(invertRgn, 2,
		uintptr(hdc),
		uintptr(hrgn),
		0)
	return ret1 != 0
}

func LPtoDP(hdc HDC, lppt *POINT, c int32) bool {
	ret1 := syscall3(lPtoDP, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lppt)),
		uintptr(c))
	return ret1 != 0
}

func LineDDA(xStart int32, yStart int32, xEnd int32, yEnd int32, lpProc LINEDDAPROC, data LPARAM) bool {
	lpProcCallback := syscall.NewCallback(lpProc)
	ret1 := syscall6(lineDDA, 6,
		uintptr(xStart),
		uintptr(yStart),
		uintptr(xEnd),
		uintptr(yEnd),
		lpProcCallback,
		uintptr(data))
	return ret1 != 0
}

func LineTo(hdc HDC, x int32, y int32) bool {
	ret1 := syscall3(lineTo, 3,
		uintptr(hdc),
		uintptr(x),
		uintptr(y))
	return ret1 != 0
}

func MaskBlt(hdcDest HDC, xDest int32, yDest int32, width int32, height int32, hdcSrc HDC, xSrc int32, ySrc int32, hbmMask HBITMAP, xMask int32, yMask int32, rop uint32) bool {
	ret1 := syscall12(maskBlt, 12,
		uintptr(hdcDest),
		uintptr(xDest),
		uintptr(yDest),
		uintptr(width),
		uintptr(height),
		uintptr(hdcSrc),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(hbmMask),
		uintptr(xMask),
		uintptr(yMask),
		uintptr(rop))
	return ret1 != 0
}

func ModifyWorldTransform(hdc HDC, lpxf /*const*/ *XFORM, mode uint32) bool {
	ret1 := syscall3(modifyWorldTransform, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpxf)),
		uintptr(mode))
	return ret1 != 0
}

func MoveToEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	ret1 := syscall6(moveToEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)
	return ret1 != 0
}

func OffsetClipRgn(hdc HDC, x int32, y int32) int32 {
	ret1 := syscall3(offsetClipRgn, 3,
		uintptr(hdc),
		uintptr(x),
		uintptr(y))
	return int32(ret1)
}

func OffsetRgn(hrgn HRGN, x int32, y int32) int32 {
	ret1 := syscall3(offsetRgn, 3,
		uintptr(hrgn),
		uintptr(x),
		uintptr(y))
	return int32(ret1)
}

func OffsetViewportOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	ret1 := syscall6(offsetViewportOrgEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)
	return ret1 != 0
}

func OffsetWindowOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	ret1 := syscall6(offsetWindowOrgEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)
	return ret1 != 0
}

func PATHOBJ_bEnum(ppo *PATHOBJ, ppd *PATHDATA) bool {
	ret1 := syscall3(pATHOBJ_bEnum, 2,
		uintptr(unsafe.Pointer(ppo)),
		uintptr(unsafe.Pointer(ppd)),
		0)
	return ret1 != 0
}

func PATHOBJ_bEnumClipLines(ppo *PATHOBJ, cb ULONG, pcl *CLIPLINE) bool {
	ret1 := syscall3(pATHOBJ_bEnumClipLines, 3,
		uintptr(unsafe.Pointer(ppo)),
		uintptr(cb),
		uintptr(unsafe.Pointer(pcl)))
	return ret1 != 0
}

func PATHOBJ_vEnumStart(ppo *PATHOBJ) {
	syscall3(pATHOBJ_vEnumStart, 1,
		uintptr(unsafe.Pointer(ppo)),
		0,
		0)
}

func PATHOBJ_vEnumStartClipLines(ppo *PATHOBJ, pco *CLIPOBJ, pso *SURFOBJ, pla *LINEATTRS) {
	syscall6(pATHOBJ_vEnumStartClipLines, 4,
		uintptr(unsafe.Pointer(ppo)),
		uintptr(unsafe.Pointer(pco)),
		uintptr(unsafe.Pointer(pso)),
		uintptr(unsafe.Pointer(pla)),
		0,
		0)
}

func PATHOBJ_vGetBounds(ppo *PATHOBJ, prectfx PRECTFX) {
	syscall3(pATHOBJ_vGetBounds, 2,
		uintptr(unsafe.Pointer(ppo)),
		uintptr(unsafe.Pointer(prectfx)),
		0)
}

func PaintRgn(hdc HDC, hrgn HRGN) bool {
	ret1 := syscall3(paintRgn, 2,
		uintptr(hdc),
		uintptr(hrgn),
		0)
	return ret1 != 0
}

func PatBlt(hdc HDC, x int32, y int32, w int32, h int32, rop uint32) bool {
	ret1 := syscall6(patBlt, 6,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(rop))
	return ret1 != 0
}

func PathToRegion(hdc HDC) HRGN {
	ret1 := syscall3(pathToRegion, 1,
		uintptr(hdc),
		0,
		0)
	return HRGN(ret1)
}

func Pie(hdc HDC, left int32, top int32, right int32, bottom int32, xr1 int32, yr1 int32, xr2 int32, yr2 int32) bool {
	ret1 := syscall9(pie, 9,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		uintptr(xr1),
		uintptr(yr1),
		uintptr(xr2),
		uintptr(yr2))
	return ret1 != 0
}

func PlayEnhMetaFile(hdc HDC, hmf HENHMETAFILE, lprect /*const*/ *RECT) bool {
	ret1 := syscall3(playEnhMetaFile, 3,
		uintptr(hdc),
		uintptr(hmf),
		uintptr(unsafe.Pointer(lprect)))
	return ret1 != 0
}

func PlayEnhMetaFileRecord(hdc HDC, pht *HANDLETABLE, pmr /*const*/ *ENHMETARECORD, cht uint32) bool {
	ret1 := syscall6(playEnhMetaFileRecord, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pht)),
		uintptr(unsafe.Pointer(pmr)),
		uintptr(cht),
		0,
		0)
	return ret1 != 0
}

func PlayMetaFile(hdc HDC, hmf HMETAFILE) bool {
	ret1 := syscall3(playMetaFile, 2,
		uintptr(hdc),
		uintptr(hmf),
		0)
	return ret1 != 0
}

func PlayMetaFileRecord(hdc HDC, lpHandleTable *HANDLETABLE, lpMR *METARECORD, noObjs uint32) bool {
	ret1 := syscall6(playMetaFileRecord, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpHandleTable)),
		uintptr(unsafe.Pointer(lpMR)),
		uintptr(noObjs),
		0,
		0)
	return ret1 != 0
}

func PlgBlt(hdcDest HDC, lpPoint /*const*/ *POINT, hdcSrc HDC, xSrc int32, ySrc int32, width int32, height int32, hbmMask HBITMAP, xMask int32, yMask int32) bool {
	ret1 := syscall12(plgBlt, 10,
		uintptr(hdcDest),
		uintptr(unsafe.Pointer(lpPoint)),
		uintptr(hdcSrc),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(width),
		uintptr(height),
		uintptr(hbmMask),
		uintptr(xMask),
		uintptr(yMask),
		0,
		0)
	return ret1 != 0
}

func PolyBezier(hdc HDC, apt /*const*/ *POINT, cpt uint32) bool {
	ret1 := syscall3(polyBezier, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(cpt))
	return ret1 != 0
}

func PolyBezierTo(hdc HDC, apt /*const*/ *POINT, cpt uint32) bool {
	ret1 := syscall3(polyBezierTo, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(cpt))
	return ret1 != 0
}

func PolyDraw(hdc HDC, apt /*const*/ *POINT, aj /*const*/ *byte, cpt int32) bool {
	ret1 := syscall6(polyDraw, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(unsafe.Pointer(aj)),
		uintptr(cpt),
		0,
		0)
	return ret1 != 0
}

func PolyPolygon(hdc HDC, apt /*const*/ *POINT, asz /*const*/ *int32, csz int32) bool {
	ret1 := syscall6(polyPolygon, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(unsafe.Pointer(asz)),
		uintptr(csz),
		0,
		0)
	return ret1 != 0
}

func PolyPolyline(hdc HDC, apt /*const*/ *POINT, asz /*const*/ *uint32, csz uint32) bool {
	ret1 := syscall6(polyPolyline, 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(unsafe.Pointer(asz)),
		uintptr(csz),
		0,
		0)
	return ret1 != 0
}

func PolyTextOut(hdc HDC, ppt /*const*/ *POLYTEXT, nstrings int32) bool {
	ret1 := syscall3(polyTextOut, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(ppt)),
		uintptr(nstrings))
	return ret1 != 0
}

func Polygon(hdc HDC, apt /*const*/ *POINT, cpt int32) bool {
	ret1 := syscall3(polygon, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(cpt))
	return ret1 != 0
}

func Polyline(hdc HDC, apt /*const*/ *POINT, cpt int32) bool {
	ret1 := syscall3(polyline, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(cpt))
	return ret1 != 0
}

func PolylineTo(hdc HDC, apt /*const*/ *POINT, cpt uint32) bool {
	ret1 := syscall3(polylineTo, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(apt)),
		uintptr(cpt))
	return ret1 != 0
}

func PtInRegion(hrgn HRGN, x int32, y int32) bool {
	ret1 := syscall3(ptInRegion, 3,
		uintptr(hrgn),
		uintptr(x),
		uintptr(y))
	return ret1 != 0
}

func PtVisible(hdc HDC, x int32, y int32) bool {
	ret1 := syscall3(ptVisible, 3,
		uintptr(hdc),
		uintptr(x),
		uintptr(y))
	return ret1 != 0
}

func RealizePalette(hdc HDC) uint32 {
	ret1 := syscall3(realizePalette, 1,
		uintptr(hdc),
		0,
		0)
	return uint32(ret1)
}

func RectInRegion(hrgn HRGN, lprect /*const*/ *RECT) bool {
	ret1 := syscall3(rectInRegion, 2,
		uintptr(hrgn),
		uintptr(unsafe.Pointer(lprect)),
		0)
	return ret1 != 0
}

func RectVisible(hdc HDC, lprect /*const*/ *RECT) bool {
	ret1 := syscall3(rectVisible, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lprect)),
		0)
	return ret1 != 0
}

func Rectangle(hdc HDC, left int32, top int32, right int32, bottom int32) bool {
	ret1 := syscall6(rectangle, 5,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		0)
	return ret1 != 0
}

func RemoveFontMemResourceEx(h HANDLE) bool {
	ret1 := syscall3(removeFontMemResourceEx, 1,
		uintptr(h),
		0,
		0)
	return ret1 != 0
}

func RemoveFontResourceEx(name string, fl uint32, pdv uintptr) bool {
	nameStr := unicode16FromString(name)
	ret1 := syscall3(removeFontResourceEx, 3,
		uintptr(unsafe.Pointer(&nameStr[0])),
		uintptr(fl),
		pdv)
	return ret1 != 0
}

func RemoveFontResource(lpFileName string) bool {
	lpFileNameStr := unicode16FromString(lpFileName)
	ret1 := syscall3(removeFontResource, 1,
		uintptr(unsafe.Pointer(&lpFileNameStr[0])),
		0,
		0)
	return ret1 != 0
}

func ResetDC(hdc HDC, lpdm /*const*/ *DEVMODE) HDC {
	ret1 := syscall3(resetDC, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpdm)),
		0)
	return HDC(ret1)
}

func ResizePalette(hpal HPALETTE, n uint32) bool {
	ret1 := syscall3(resizePalette, 2,
		uintptr(hpal),
		uintptr(n),
		0)
	return ret1 != 0
}

func RestoreDC(hdc HDC, nSavedDC int32) bool {
	ret1 := syscall3(restoreDC, 2,
		uintptr(hdc),
		uintptr(nSavedDC),
		0)
	return ret1 != 0
}

func RoundRect(hdc HDC, left int32, top int32, right int32, bottom int32, width int32, height int32) bool {
	ret1 := syscall9(roundRect, 7,
		uintptr(hdc),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		uintptr(width),
		uintptr(height),
		0,
		0)
	return ret1 != 0
}

func STROBJ_bEnum(pstro *STROBJ, pc *ULONG, ppgpos *PGLYPHPOS) bool {
	ret1 := syscall3(sTROBJ_bEnum, 3,
		uintptr(unsafe.Pointer(pstro)),
		uintptr(unsafe.Pointer(pc)),
		uintptr(unsafe.Pointer(ppgpos)))
	return ret1 != 0
}

func STROBJ_bEnumPositionsOnly(pstro *STROBJ, pc *ULONG, ppgpos *PGLYPHPOS) bool {
	ret1 := syscall3(sTROBJ_bEnumPositionsOnly, 3,
		uintptr(unsafe.Pointer(pstro)),
		uintptr(unsafe.Pointer(pc)),
		uintptr(unsafe.Pointer(ppgpos)))
	return ret1 != 0
}

func STROBJ_bGetAdvanceWidths(pso *STROBJ, iFirst ULONG, c ULONG, pptqD *POINTQF) bool {
	ret1 := syscall6(sTROBJ_bGetAdvanceWidths, 4,
		uintptr(unsafe.Pointer(pso)),
		uintptr(iFirst),
		uintptr(c),
		uintptr(unsafe.Pointer(pptqD)),
		0,
		0)
	return ret1 != 0
}

func STROBJ_dwGetCodePage(pstro *STROBJ) uint32 {
	ret1 := syscall3(sTROBJ_dwGetCodePage, 1,
		uintptr(unsafe.Pointer(pstro)),
		0,
		0)
	return uint32(ret1)
}

func STROBJ_vEnumStart(pstro *STROBJ) {
	syscall3(sTROBJ_vEnumStart, 1,
		uintptr(unsafe.Pointer(pstro)),
		0,
		0)
}

func SaveDC(hdc HDC) int32 {
	ret1 := syscall3(saveDC, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func ScaleViewportExtEx(hdc HDC, xn int32, dx int32, yn int32, yd int32, lpsz *SIZE) bool {
	ret1 := syscall6(scaleViewportExtEx, 6,
		uintptr(hdc),
		uintptr(xn),
		uintptr(dx),
		uintptr(yn),
		uintptr(yd),
		uintptr(unsafe.Pointer(lpsz)))
	return ret1 != 0
}

func ScaleWindowExtEx(hdc HDC, xn int32, xd int32, yn int32, yd int32, lpsz *SIZE) bool {
	ret1 := syscall6(scaleWindowExtEx, 6,
		uintptr(hdc),
		uintptr(xn),
		uintptr(xd),
		uintptr(yn),
		uintptr(yd),
		uintptr(unsafe.Pointer(lpsz)))
	return ret1 != 0
}

func SelectClipPath(hdc HDC, mode int32) bool {
	ret1 := syscall3(selectClipPath, 2,
		uintptr(hdc),
		uintptr(mode),
		0)
	return ret1 != 0
}

func SelectClipRgn(hdc HDC, hrgn HRGN) int32 {
	ret1 := syscall3(selectClipRgn, 2,
		uintptr(hdc),
		uintptr(hrgn),
		0)
	return int32(ret1)
}

func SelectObject(hdc HDC, h HGDIOBJ) HGDIOBJ {
	ret1 := syscall3(selectObject, 2,
		uintptr(hdc),
		uintptr(h),
		0)
	return HGDIOBJ(ret1)
}

func SelectPalette(hdc HDC, hPal HPALETTE, bForceBkgd bool) HPALETTE {
	ret1 := syscall3(selectPalette, 3,
		uintptr(hdc),
		uintptr(hPal),
		getUintptrFromBool(bForceBkgd))
	return HPALETTE(ret1)
}

func SetAbortProc(hdc HDC, lpProc ABORTPROC) int32 {
	lpProcCallback := syscall.NewCallback(func(unnamed0RawArg HDC, unnamed1RawArg int32) uintptr {
		ret := lpProc(unnamed0RawArg, unnamed1RawArg)
		return uintptr(ret)
	})
	ret1 := syscall3(setAbortProc, 2,
		uintptr(hdc),
		lpProcCallback,
		0)
	return int32(ret1)
}

func SetArcDirection(hdc HDC, dir int32) int32 {
	ret1 := syscall3(setArcDirection, 2,
		uintptr(hdc),
		uintptr(dir),
		0)
	return int32(ret1)
}

func SetBitmapBits(hbm HBITMAP, cb uint32, pvBits /*const*/ uintptr) LONG {
	ret1 := syscall3(setBitmapBits, 3,
		uintptr(hbm),
		uintptr(cb),
		pvBits)
	return LONG(ret1)
}

func SetBitmapDimensionEx(hbm HBITMAP, w int32, h int32, lpsz *SIZE) bool {
	ret1 := syscall6(setBitmapDimensionEx, 4,
		uintptr(hbm),
		uintptr(w),
		uintptr(h),
		uintptr(unsafe.Pointer(lpsz)),
		0,
		0)
	return ret1 != 0
}

func SetBkColor(hdc HDC, color COLORREF) COLORREF {
	ret1 := syscall3(setBkColor, 2,
		uintptr(hdc),
		uintptr(color),
		0)
	return COLORREF(ret1)
}

func SetBkMode(hdc HDC, mode int32) int32 {
	ret1 := syscall3(setBkMode, 2,
		uintptr(hdc),
		uintptr(mode),
		0)
	return int32(ret1)
}

func SetBoundsRect(hdc HDC, lprect /*const*/ *RECT, flags uint32) uint32 {
	ret1 := syscall3(setBoundsRect, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lprect)),
		uintptr(flags))
	return uint32(ret1)
}

func SetBrushOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	ret1 := syscall6(setBrushOrgEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)
	return ret1 != 0
}

func SetColorAdjustment(hdc HDC, lpca /*const*/ *COLORADJUSTMENT) bool {
	ret1 := syscall3(setColorAdjustment, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpca)),
		0)
	return ret1 != 0
}

func SetColorSpace(hdc HDC, hcs HCOLORSPACE) HCOLORSPACE {
	ret1 := syscall3(setColorSpace, 2,
		uintptr(hdc),
		uintptr(hcs),
		0)
	return HCOLORSPACE(ret1)
}

func SetDCBrushColor(hdc HDC, color COLORREF) COLORREF {
	ret1 := syscall3(setDCBrushColor, 2,
		uintptr(hdc),
		uintptr(color),
		0)
	return COLORREF(ret1)
}

func SetDCPenColor(hdc HDC, color COLORREF) COLORREF {
	ret1 := syscall3(setDCPenColor, 2,
		uintptr(hdc),
		uintptr(color),
		0)
	return COLORREF(ret1)
}

func SetDIBColorTable(hdc HDC, iStart uint32, cEntries uint32, prgbq /*const*/ *RGBQUAD) uint32 {
	ret1 := syscall6(setDIBColorTable, 4,
		uintptr(hdc),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(prgbq)),
		0,
		0)
	return uint32(ret1)
}

func SetDIBits(hdc HDC, hbm HBITMAP, start uint32, cLines uint32, lpBits /*const*/ uintptr, lpbmi /*const*/ *BITMAPINFO, colorUse uint32) int32 {
	ret1 := syscall9(setDIBits, 7,
		uintptr(hdc),
		uintptr(hbm),
		uintptr(start),
		uintptr(cLines),
		lpBits,
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(colorUse),
		0,
		0)
	return int32(ret1)
}

func SetDIBitsToDevice(hdc HDC, xDest int32, yDest int32, w uint32, h uint32, xSrc int32, ySrc int32, startScan uint32, cLines uint32, lpvBits /*const*/ uintptr, lpbmi /*const*/ *BITMAPINFO, colorUse uint32) int32 {
	ret1 := syscall12(setDIBitsToDevice, 12,
		uintptr(hdc),
		uintptr(xDest),
		uintptr(yDest),
		uintptr(w),
		uintptr(h),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(startScan),
		uintptr(cLines),
		lpvBits,
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(colorUse))
	return int32(ret1)
}

func SetDeviceGammaRamp(hdc HDC, lpRamp LPVOID) bool {
	ret1 := syscall3(setDeviceGammaRamp, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpRamp)),
		0)
	return ret1 != 0
}

func SetEnhMetaFileBits(nSize uint32, pb /*const*/ *byte) HENHMETAFILE {
	ret1 := syscall3(setEnhMetaFileBits, 2,
		uintptr(nSize),
		uintptr(unsafe.Pointer(pb)),
		0)
	return HENHMETAFILE(ret1)
}

func SetGraphicsMode(hdc HDC, iMode int32) int32 {
	ret1 := syscall3(setGraphicsMode, 2,
		uintptr(hdc),
		uintptr(iMode),
		0)
	return int32(ret1)
}

func SetICMMode(hdc HDC, mode int32) int32 {
	ret1 := syscall3(setICMMode, 2,
		uintptr(hdc),
		uintptr(mode),
		0)
	return int32(ret1)
}

func SetICMProfile(hdc HDC, lpFileName LPWSTR) bool {
	ret1 := syscall3(setICMProfile, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpFileName)),
		0)
	return ret1 != 0
}

func SetLayout(hdc HDC, l uint32) uint32 {
	ret1 := syscall3(setLayout, 2,
		uintptr(hdc),
		uintptr(l),
		0)
	return uint32(ret1)
}

func SetMapMode(hdc HDC, iMode int32) int32 {
	ret1 := syscall3(setMapMode, 2,
		uintptr(hdc),
		uintptr(iMode),
		0)
	return int32(ret1)
}

func SetMapperFlags(hdc HDC, flags uint32) uint32 {
	ret1 := syscall3(setMapperFlags, 2,
		uintptr(hdc),
		uintptr(flags),
		0)
	return uint32(ret1)
}

func SetMetaFileBitsEx(cbBuffer uint32, lpData /*const*/ *byte) HMETAFILE {
	ret1 := syscall3(setMetaFileBitsEx, 2,
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(lpData)),
		0)
	return HMETAFILE(ret1)
}

func SetMetaRgn(hdc HDC) int32 {
	ret1 := syscall3(setMetaRgn, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func SetMiterLimit(hdc HDC, limit float32, old *float32) bool {
	ret1 := syscall3(setMiterLimit, 3,
		uintptr(hdc),
		uintptr(limit),
		uintptr(unsafe.Pointer(old)))
	return ret1 != 0
}

func SetPaletteEntries(hpal HPALETTE, iStart uint32, cEntries uint32, pPalEntries /*const*/ *PALETTEENTRY) uint32 {
	ret1 := syscall6(setPaletteEntries, 4,
		uintptr(hpal),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pPalEntries)),
		0,
		0)
	return uint32(ret1)
}

func SetPixel(hdc HDC, x int32, y int32, color COLORREF) COLORREF {
	ret1 := syscall6(setPixel, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(color),
		0,
		0)
	return COLORREF(ret1)
}

func SetPixelFormat(hdc HDC, format int32, ppfd /*const*/ *PIXELFORMATDESCRIPTOR) bool {
	ret1 := syscall3(setPixelFormat, 3,
		uintptr(hdc),
		uintptr(format),
		uintptr(unsafe.Pointer(ppfd)))
	return ret1 != 0
}

func SetPixelV(hdc HDC, x int32, y int32, color COLORREF) bool {
	ret1 := syscall6(setPixelV, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(color),
		0,
		0)
	return ret1 != 0
}

func SetPolyFillMode(hdc HDC, mode int32) int32 {
	ret1 := syscall3(setPolyFillMode, 2,
		uintptr(hdc),
		uintptr(mode),
		0)
	return int32(ret1)
}

func SetROP2(hdc HDC, rop2 int32) int32 {
	ret1 := syscall3(setROP2, 2,
		uintptr(hdc),
		uintptr(rop2),
		0)
	return int32(ret1)
}

func SetRectRgn(hrgn HRGN, left int32, top int32, right int32, bottom int32) bool {
	ret1 := syscall6(setRectRgn, 5,
		uintptr(hrgn),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom),
		0)
	return ret1 != 0
}

func SetStretchBltMode(hdc HDC, mode int32) int32 {
	ret1 := syscall3(setStretchBltMode, 2,
		uintptr(hdc),
		uintptr(mode),
		0)
	return int32(ret1)
}

func SetSystemPaletteUse(hdc HDC, use uint32) uint32 {
	ret1 := syscall3(setSystemPaletteUse, 2,
		uintptr(hdc),
		uintptr(use),
		0)
	return uint32(ret1)
}

func SetTextAlign(hdc HDC, align uint32) uint32 {
	ret1 := syscall3(setTextAlign, 2,
		uintptr(hdc),
		uintptr(align),
		0)
	return uint32(ret1)
}

func SetTextCharacterExtra(hdc HDC, extra int32) int32 {
	ret1 := syscall3(setTextCharacterExtra, 2,
		uintptr(hdc),
		uintptr(extra),
		0)
	return int32(ret1)
}

func SetTextColor(hdc HDC, color COLORREF) COLORREF {
	ret1 := syscall3(setTextColor, 2,
		uintptr(hdc),
		uintptr(color),
		0)
	return COLORREF(ret1)
}

func SetTextJustification(hdc HDC, extra int32, count int32) bool {
	ret1 := syscall3(setTextJustification, 3,
		uintptr(hdc),
		uintptr(extra),
		uintptr(count))
	return ret1 != 0
}

func SetViewportExtEx(hdc HDC, x int32, y int32, lpsz *SIZE) bool {
	ret1 := syscall6(setViewportExtEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpsz)),
		0,
		0)
	return ret1 != 0
}

func SetViewportOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	ret1 := syscall6(setViewportOrgEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)
	return ret1 != 0
}

func SetWinMetaFileBits(nSize uint32, lpMeta16Data /*const*/ *byte, hdcRef HDC, lpMFP /*const*/ *METAFILEPICT) HENHMETAFILE {
	ret1 := syscall6(setWinMetaFileBits, 4,
		uintptr(nSize),
		uintptr(unsafe.Pointer(lpMeta16Data)),
		uintptr(hdcRef),
		uintptr(unsafe.Pointer(lpMFP)),
		0,
		0)
	return HENHMETAFILE(ret1)
}

func SetWindowExtEx(hdc HDC, x int32, y int32, lpsz *SIZE) bool {
	ret1 := syscall6(setWindowExtEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpsz)),
		0,
		0)
	return ret1 != 0
}

func SetWindowOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	ret1 := syscall6(setWindowOrgEx, 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)
	return ret1 != 0
}

func SetWorldTransform(hdc HDC, lpxf /*const*/ *XFORM) bool {
	ret1 := syscall3(setWorldTransform, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpxf)),
		0)
	return ret1 != 0
}

func StartDoc(hdc HDC, lpdi /*const*/ *DOCINFO) int32 {
	ret1 := syscall3(startDoc, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpdi)),
		0)
	return int32(ret1)
}

func StartPage(hdc HDC) int32 {
	ret1 := syscall3(startPage, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func StretchBlt(hdcDest HDC, xDest int32, yDest int32, wDest int32, hDest int32, hdcSrc HDC, xSrc int32, ySrc int32, wSrc int32, hSrc int32, rop uint32) bool {
	ret1 := syscall12(stretchBlt, 11,
		uintptr(hdcDest),
		uintptr(xDest),
		uintptr(yDest),
		uintptr(wDest),
		uintptr(hDest),
		uintptr(hdcSrc),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(wSrc),
		uintptr(hSrc),
		uintptr(rop),
		0)
	return ret1 != 0
}

func StretchDIBits(hdc HDC, xDest int32, yDest int32, destWidth int32, destHeight int32, xSrc int32, ySrc int32, srcWidth int32, srcHeight int32, lpBits /*const*/ uintptr, lpbmi /*const*/ *BITMAPINFO, iUsage uint32, rop uint32) int32 {
	ret1 := syscall15(stretchDIBits, 13,
		uintptr(hdc),
		uintptr(xDest),
		uintptr(yDest),
		uintptr(destWidth),
		uintptr(destHeight),
		uintptr(xSrc),
		uintptr(ySrc),
		uintptr(srcWidth),
		uintptr(srcHeight),
		lpBits,
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(iUsage),
		uintptr(rop),
		0,
		0)
	return int32(ret1)
}

func StrokeAndFillPath(hdc HDC) bool {
	ret1 := syscall3(strokeAndFillPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func StrokePath(hdc HDC) bool {
	ret1 := syscall3(strokePath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func SwapBuffers(unnamed0 HDC) bool {
	ret1 := syscall3(swapBuffers, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func TextOut(hdc HDC, x int32, y int32, lpString string, c int32) bool {
	lpStringStr := unicode16FromString(lpString)
	ret1 := syscall6(textOut, 5,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(&lpStringStr[0])),
		uintptr(c),
		0)
	return ret1 != 0
}

func TranslateCharsetInfo(lpSrc *uint32, lpCs *CHARSETINFO, dwFlags uint32) bool {
	ret1 := syscall3(translateCharsetInfo, 3,
		uintptr(unsafe.Pointer(lpSrc)),
		uintptr(unsafe.Pointer(lpCs)),
		uintptr(dwFlags))
	return ret1 != 0
}

func UnrealizeObject(h HGDIOBJ) bool {
	ret1 := syscall3(unrealizeObject, 1,
		uintptr(h),
		0,
		0)
	return ret1 != 0
}

func UpdateColors(hdc HDC) bool {
	ret1 := syscall3(updateColors, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func UpdateICMRegKey(reserved uint32, lpszCMID LPWSTR, lpszFileName LPWSTR, command uint32) bool {
	ret1 := syscall6(updateICMRegKey, 4,
		uintptr(reserved),
		uintptr(unsafe.Pointer(lpszCMID)),
		uintptr(unsafe.Pointer(lpszFileName)),
		uintptr(command),
		0,
		0)
	return ret1 != 0
}

func WidenPath(hdc HDC) bool {
	ret1 := syscall3(widenPath, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func XFORMOBJ_bApplyXform(pxo *XFORMOBJ, iMode ULONG, cPoints ULONG, pvIn uintptr, pvOut uintptr) bool {
	ret1 := syscall6(xFORMOBJ_bApplyXform, 5,
		uintptr(unsafe.Pointer(pxo)),
		uintptr(iMode),
		uintptr(cPoints),
		pvIn,
		pvOut,
		0)
	return ret1 != 0
}

func XFORMOBJ_iGetXform(pxo *XFORMOBJ, pxform *XFORML) ULONG {
	ret1 := syscall3(xFORMOBJ_iGetXform, 2,
		uintptr(unsafe.Pointer(pxo)),
		uintptr(unsafe.Pointer(pxform)),
		0)
	return ULONG(ret1)
}

func XLATEOBJ_cGetPalette(pxlo *XLATEOBJ, iPal ULONG, cPal ULONG, pPal *ULONG) ULONG {
	ret1 := syscall6(xLATEOBJ_cGetPalette, 4,
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(iPal),
		uintptr(cPal),
		uintptr(unsafe.Pointer(pPal)),
		0,
		0)
	return ULONG(ret1)
}

func XLATEOBJ_hGetColorTransform(pxlo *XLATEOBJ) HANDLE {
	ret1 := syscall3(xLATEOBJ_hGetColorTransform, 1,
		uintptr(unsafe.Pointer(pxlo)),
		0,
		0)
	return HANDLE(ret1)
}

func XLATEOBJ_iXlate(pxlo *XLATEOBJ, iColor ULONG) ULONG {
	ret1 := syscall3(xLATEOBJ_iXlate, 2,
		uintptr(unsafe.Pointer(pxlo)),
		uintptr(iColor),
		0)
	return ULONG(ret1)
}

func EnableEUDC(fEnableEUDC bool) bool {
	ret1 := syscall3(enableEUDC, 1,
		getUintptrFromBool(fEnableEUDC),
		0,
		0)
	return ret1 != 0
}

func FontIsLinked(hdc HDC) bool {
	ret1 := syscall3(fontIsLinked, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func GdiDescribePixelFormat(hdc HDC, format int32, size uint32, descr *PIXELFORMATDESCRIPTOR) int32 {
	ret1 := syscall6(gdiDescribePixelFormat, 4,
		uintptr(hdc),
		uintptr(format),
		uintptr(size),
		uintptr(unsafe.Pointer(descr)),
		0,
		0)
	return int32(ret1)
}

func GdiDrawStream(hdc HDC, in ULONG, pvin uintptr) bool {
	ret1 := syscall3(gdiDrawStream, 3,
		uintptr(hdc),
		uintptr(in),
		pvin)
	return ret1 != 0
}

func GdiGetCharDimensions(hdc HDC, lptm LPTEXTMETRIC, height *LONG) LONG {
	ret1 := syscall3(gdiGetCharDimensions, 3,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lptm)),
		uintptr(unsafe.Pointer(height)))
	return LONG(ret1)
}

func GdiGetCodePage(hdc HDC) uint32 {
	ret1 := syscall3(gdiGetCodePage, 1,
		uintptr(hdc),
		0,
		0)
	return uint32(ret1)
}

func GdiGetSpoolMessage(ptr1 LPVOID, data2 uint32, ptr3 LPVOID, data4 uint32) uint32 {
	ret1 := syscall6(gdiGetSpoolMessage, 4,
		uintptr(unsafe.Pointer(ptr1)),
		uintptr(data2),
		uintptr(unsafe.Pointer(ptr3)),
		uintptr(data4),
		0,
		0)
	return uint32(ret1)
}

func GdiInitSpool() uint32 {
	ret1 := syscall3(gdiInitSpool, 0,
		0,
		0,
		0)
	return uint32(ret1)
}

func GdiInitializeLanguagePack(arg uint32) uint32 {
	ret1 := syscall3(gdiInitializeLanguagePack, 1,
		uintptr(arg),
		0,
		0)
	return uint32(ret1)
}

func GdiIsMetaFileDC(hdc HDC) bool {
	ret1 := syscall3(gdiIsMetaFileDC, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func GdiIsMetaPrintDC(hdc HDC) bool {
	ret1 := syscall3(gdiIsMetaPrintDC, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func GdiIsPlayMetafileDC(hdc HDC) bool {
	ret1 := syscall3(gdiIsPlayMetafileDC, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func GdiRealizationInfo(hdc HDC, info uintptr) bool {
	ret1 := syscall3(gdiRealizationInfo, 2,
		uintptr(hdc),
		info,
		0)
	return ret1 != 0
}

func GdiSetPixelFormat(hdc HDC, fmt int32, pfd /*const*/ *PIXELFORMATDESCRIPTOR) bool {
	ret1 := syscall3(gdiSetPixelFormat, 3,
		uintptr(hdc),
		uintptr(fmt),
		uintptr(unsafe.Pointer(pfd)))
	return ret1 != 0
}

func GdiSwapBuffers(hdc HDC) bool {
	ret1 := syscall3(gdiSwapBuffers, 1,
		uintptr(hdc),
		0,
		0)
	return ret1 != 0
}

func GetFontResourceInfoW(str string, size *uint32, buffer uintptr, aType uint32) bool {
	strStr := unicode16FromString(str)
	ret1 := syscall6(getFontResourceInfoW, 4,
		uintptr(unsafe.Pointer(&strStr[0])),
		uintptr(unsafe.Pointer(size)),
		buffer,
		uintptr(aType),
		0,
		0)
	return ret1 != 0
}

func GetRelAbs(hdc HDC, dwIgnore uint32) int32 {
	ret1 := syscall3(getRelAbs, 2,
		uintptr(hdc),
		uintptr(dwIgnore),
		0)
	return int32(ret1)
}

func GetTransform(hdc HDC, which uint32, xform *XFORM) bool {
	ret1 := syscall3(getTransform, 3,
		uintptr(hdc),
		uintptr(which),
		uintptr(unsafe.Pointer(xform)))
	return ret1 != 0
}

func MirrorRgn(hwnd HWND, hrgn HRGN) bool {
	ret1 := syscall3(mirrorRgn, 2,
		uintptr(hwnd),
		uintptr(hrgn),
		0)
	return ret1 != 0
}

func NamedEscape(hdc HDC, pDriver string, nEscape int32, cbInput int32, lpszInData /*const*/ LPCSTR, cbOutput int32, lpszOutData LPSTR) int32 {
	pDriverStr := unicode16FromString(pDriver)
	ret1 := syscall9(namedEscape, 7,
		uintptr(hdc),
		uintptr(unsafe.Pointer(&pDriverStr[0])),
		uintptr(nEscape),
		uintptr(cbInput),
		uintptr(unsafe.Pointer(lpszInData)),
		uintptr(cbOutput),
		uintptr(unsafe.Pointer(lpszOutData)),
		0,
		0)
	return int32(ret1)
}

func SetMagicColors(hdc HDC, u1 ULONG, u2 ULONG) bool {
	ret1 := syscall3(setMagicColors, 3,
		uintptr(hdc),
		uintptr(u1),
		uintptr(u2))
	return ret1 != 0
}

func SetRelAbs(hdc HDC, mode int32) int32 {
	ret1 := syscall3(setRelAbs, 2,
		uintptr(hdc),
		uintptr(mode),
		0)
	return int32(ret1)
}

func SetVirtualResolution(hdc HDC, horz_res uint32, vert_res uint32, horz_size uint32, vert_size uint32) bool {
	ret1 := syscall6(setVirtualResolution, 5,
		uintptr(hdc),
		uintptr(horz_res),
		uintptr(vert_res),
		uintptr(horz_size),
		uintptr(vert_size),
		0)
	return ret1 != 0
}