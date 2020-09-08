package main

import (
	"github.com/jamesmoriarty/gomem"
	"fmt"
	"log"
	"time"
	"ReadWriteMemory" // If the import is not working, put the /ReadWriteMemory in your GOPATH
)
const ( // https://github.com/frk1/hazedumper for offsets
	dwLocalPlayer       = 0xd3fc5c
	dwEntityList        = 0x4d5442c
	iGlowIndex          = 0xa438
	m_iTeamNum          = 0xF4
	m_iCrosshairId	    = 0xb3e4
	dwForceAttack	    = 0x3185984
)

func main() {
	process, err := ReadWriteMemory.ProcessByName("csgo")
	if err != nil {
		log.Panicf("csgo.exe not found. Error: %s", err.Error())
	}
	client := process.Modules["client.dll"].ModBaseAddr
	if err != nil {
		fmt.Println(err)
	}
	for {
		if !gomem.IsKeyDown(0x10) {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		
		player, _ := process.ReadIntPtr(client + dwLocalPlayer)
		entity_id, _ := process.ReadIntPtr(player + m_iCrosshairId)
		entity, _ := process.ReadIntPtr(client + dwEntityList + (entity_id - 1) * 0x10)
		entity_team, _ := process.ReadIntPtr(entity + m_iTeamNum)
		player_team, _ := process.ReadIntPtr(player + m_iTeamNum)
		
		if gomem.IsKeyDown(0x10){ // 0x10 = Shift
			
			if entity_id > 0 && entity_id <= 64 && player_team != entity_team { // Checks if the crosshair is on an entity and checks if that entity is an enemy player
					process.WriteInt(client + dwForceAttack, 6)
					time.Sleep(time.Millisecond * 120)
		time.Sleep(time.Millisecond * 10)
		}
	}
	}
}
