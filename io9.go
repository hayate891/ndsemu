package main

import "ndsemu/emu/hwio"

type NDSIOCommon struct {
	postflg uint8
}

type NDS9IOMap struct {
	TableLo hwio.Table
	TableHi hwio.Table

	GetPC   func() uint32
	Card    *Gamecard
	Ipc     *HwIpc
	Mc      *HwMemoryController
	Timers  *HwTimers
	Irq     *HwIrq
	Common  *NDSIOCommon
	Lcd     *HwLcd
	Div     *HwDivisor
	Dma     [4]*HwDmaChannel
	DmaFill *HwDmaFill
	E2d     [2]*HwEngine2d
}

func (m *NDS9IOMap) Reset() {
	m.TableLo.Name = "io9"
	m.TableLo.Reset()
	m.TableHi.Name = "io9-hi"
	m.TableHi.Reset()

	m.TableLo.MapBank(0x4000000, m.E2d[0], 0)
	m.TableLo.MapBank(0x4000200, m.Irq, 0)
	m.TableLo.MapBank(0x4000240, m.Mc, 0)
	m.TableLo.MapBank(0x4000280, m.Div, 0)
	m.TableLo.MapBank(0x40000B0, m.Dma[0], 0)
	m.TableLo.MapBank(0x40000BC, m.Dma[1], 0)
	m.TableLo.MapBank(0x40000C8, m.Dma[2], 0)
	m.TableLo.MapBank(0x40000D4, m.Dma[3], 0)
	m.TableLo.MapBank(0x40000E0, m.DmaFill, 0)
	m.TableLo.MapBank(0x4000180, m.Ipc, 0)
	m.TableLo.MapBank(0x4001000, m.E2d[1], 0)

	m.TableHi.MapBank(0x4100000, m.Ipc, 1)
}

func (m *NDS9IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0300:
		return m.Common.postflg
	default:
		return m.TableLo.Read8(addr)
	}
}

func (m *NDS9IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	default:
		m.TableLo.Write8(addr, val)
	}
}

func (m *NDS9IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0004:
		return m.Lcd.ReadDISPSTAT()
	case 0x0006:
		return m.Lcd.ReadVCOUNT()
	case 0x0100:
		return m.Timers.Timers[0].ReadCounter()
	case 0x0102:
		return m.Timers.Timers[0].ReadControl()
	case 0x0104:
		return m.Timers.Timers[1].ReadCounter()
	case 0x0106:
		return m.Timers.Timers[1].ReadControl()
	case 0x0108:
		return m.Timers.Timers[2].ReadCounter()
	case 0x010A:
		return m.Timers.Timers[2].ReadControl()
	case 0x010C:
		return m.Timers.Timers[3].ReadCounter()
	case 0x010E:
		return m.Timers.Timers[3].ReadControl()
	default:
		return m.TableLo.Read16(addr)
	}
}

func (m *NDS9IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	case 0x0004:
		m.Lcd.WriteDISPSTAT(val)
	case 0x0100:
		m.Timers.Timers[0].WriteReload(val)
	case 0x0102:
		m.Timers.Timers[0].WriteControl(val)
	case 0x0104:
		m.Timers.Timers[1].WriteReload(val)
	case 0x0106:
		m.Timers.Timers[1].WriteControl(val)
	case 0x0108:
		m.Timers.Timers[2].WriteReload(val)
	case 0x010A:
		m.Timers.Timers[2].WriteControl(val)
	case 0x010C:
		m.Timers.Timers[3].WriteReload(val)
	case 0x010E:
		m.Timers.Timers[3].WriteControl(val)
	default:
		m.TableLo.Write16(addr, val)
	}
}

func (m *NDS9IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFF {
	case 0x01A0:
		return uint32(m.Card.ReadAUXSPICNT()) | (uint32(m.Card.ReadAUXSPIDATA()) << 16)
	case 0x01A4:
		return m.Card.ReadROMCTL()
	default:
		return m.TableLo.Read32(addr)
	}
}

func (m *NDS9IOMap) Write32(addr uint32, val uint32) {
	switch addr & 0xFFFF {
	case 0x0100:
		m.Timers.Timers[0].WriteReload(uint16(val))
		m.Timers.Timers[0].WriteControl(uint16(val >> 16))
	case 0x0104:
		m.Timers.Timers[1].WriteReload(uint16(val))
		m.Timers.Timers[1].WriteControl(uint16(val >> 16))
	case 0x0108:
		m.Timers.Timers[2].WriteReload(uint16(val))
		m.Timers.Timers[2].WriteControl(uint16(val >> 16))
	case 0x010C:
		m.Timers.Timers[3].WriteReload(uint16(val))
		m.Timers.Timers[3].WriteControl(uint16(val >> 16))
	case 0x01A0:
		m.Card.WriteAUXSPICNT(uint16(val & 0xFFFF))
		m.Card.WriteAUXSPIDATA(uint16(val >> 16))
	case 0x01A4:
		m.Card.WriteROMCTL(val)
	default:
		m.TableLo.Write32(addr, val)
	}
}
