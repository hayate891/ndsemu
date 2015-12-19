package debugger

import (
	"encoding/hex"
	"fmt"
	"strings"

	ui "github.com/gizak/termui"
)

func (dbg *Debugger) initUi() {
	dbg.uiCode = ui.NewList()
	dbg.uiCode.BorderLabel = "Code"
	dbg.uiCode.BorderFg = ui.ColorGreen

	dbg.uiRegs = ui.NewList()
	dbg.uiRegs.BorderLabel = "Regs"
	dbg.uiRegs.BorderFg = ui.ColorGreen

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(9, 0, dbg.uiCode),
			ui.NewCol(3, 0, dbg.uiRegs),
		),
	)

	ui.Body.Align()
}

func (dbg *Debugger) refreshCode() {
	curpc := dbg.curcpu.GetPc()
	_, data := dbg.curcpu.Disasm(curpc)

	nlines := ui.TermHeight() - 4
	lines := make([]string, nlines)
	dbg.linepc = make([]uint32, nlines)

	pc := curpc - uint32((nlines/2)*len(data))
	for i := 0; i < nlines; i++ {
		var text string
		var buf []byte
		if pc>>24 != curpc>>24 {
			// avoid disassembling cross-block
			text = "unknown"
			buf = make([]byte, len(data))
		} else {
			text, buf = dbg.curcpu.Disasm(pc)
		}
		datahex := hex.EncodeToString(buf)
		dbg.linepc[i] = pc
		lines[i] = fmt.Sprintf("   %08x  %-16s%s", pc, datahex, text)
		if pc == curpc {
			lines[i] = fmt.Sprintf("[%s%s](bg-green)", lines[i],
				strings.Repeat(" ", dbg.uiCode.Width-5-len(lines[i])))
			dbg.pcline = i
			if dbg.focusline == -1 {
				dbg.focusline = i
			}
		} else if i == dbg.focusline {
			lines[i] = fmt.Sprintf("[%s%s](bg-red)", lines[i],
				strings.Repeat(" ", dbg.uiCode.Width-5-len(lines[i])))
		}
		pc += uint32(len(buf))
	}
	dbg.uiCode.Items = lines
	dbg.uiCode.Height = len(lines) + 2
}

func (dbg *Debugger) refreshRegs() {
	names := dbg.curcpu.GetRegNames()
	values := dbg.curcpu.GetRegs()

	lines := make([]string, len(names))
	for idx := range names {
		lines[idx] = fmt.Sprintf("%4s: %08x", names[idx], values[idx])
		if len(dbg.uiRegs.Items) > idx && !strings.Contains(dbg.uiRegs.Items[idx], lines[idx]) {
			lines[idx] = fmt.Sprintf("[%s](fg-bold)", lines[idx])
		}
	}
	dbg.uiRegs.Items = lines
	dbg.uiRegs.Height = len(lines) + 2
}

func (dbg *Debugger) refreshUi() {
	dbg.refreshCode()
	dbg.refreshRegs()

	ui.Body.Align()
	ui.Render(ui.Body)
}