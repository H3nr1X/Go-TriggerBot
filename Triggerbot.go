package main

import (
	"github.com/jamesmoriarty/gomem"
	"fmt"
	"log"
	"time"
	"ReadWriteMemory"


	
)

const (
	dwLocalPlayer       = 0xd3fc5c
	dwEntityList        = 0x4d5442c
	dwGlowObjectManager = 0x529c208
	iGlowIndex          = 0xa438
	m_iTeamNum          = 0xF4
	bDormant            = 0xed
	m_iCrosshairId		= 0xb3e4
	dwForceAttack		= 0x3185984
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
		if gomem.IsKeyDown(0x10){
			if entity_id > 0 && entity_id <= 64 && player_team != entity_team {
					process.WriteInt(client + dwForceAttack, 6)
					time.Sleep(time.Millisecond * 120)
		}
	}
	}
}