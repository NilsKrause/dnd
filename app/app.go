package app

import (
	"errors"

	"de.nilskrau.dndbot/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func NewApp() *App {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.Player{},
		&models.Character{},
		&models.Message{},
	)

	return &App{
		db,
	}
}

func (a *App) DeleteCharacter(char *models.Character) error {
	tx := a.db.Delete(char)
	return tx.Error
}

func (a *App) CreateOrGetPlayer(discordid string) (*models.Player, error) {
	p := models.Player{DiscordId: discordid}

	tx := a.db.Where("discord_id = ?", discordid).First(&p)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	if tx.RowsAffected != 0 && p.ID != 0 {
		return &p, nil
	}

	tx = a.db.Create(&p)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &p, nil
}

func (a *App) GetPlayerCharacters(player *models.Player, serverID string) ([]*models.Character, error) {
	chars := make([]*models.Character, 0)
	tx := a.db.Where("player_id = ? AND server_id = ?", player.ID, serverID).Find(&chars)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return chars, nil
}

func (a *App) LogMessage(player *models.Player, char *models.Character, message string) error {
	msg := &models.Message{
		PlayerID: player.ID,
		CharacterID: char.ID,
		Message: message,
	}
	tx := a.db.Create(msg)
	return tx.Error
}

func (a *App) CharacterSetActive(id uint, p *models.Player, active bool) (*models.Character, error) {
	char := models.Character{}
	tx := a.db.Where("id = ? AND player_id = ?", id, p.ID).First(&char)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("not found")
	}

	char.Default = active

	tx = a.db.Save(&char)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &char, nil
}

func (a *App) EditCharacter(id uint, p *models.Player, name string, handle string, pic string, def bool) (*models.Character, error) {
	char := models.Character{}

	tx := a.db.Where("id = ? AND player_id = ?", id, p.ID).First(&char)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, errors.New("not found")
	}

	char.Name = name
	char.Handle = handle
	char.Image = pic
	char.Default = def

	tx = a.db.Save(&char)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &char, nil
}

func (a *App) CreateCharacter(p *models.Player, server string, name string, handle string, pic string, def bool) (*models.Character, error) {

	c := models.Character{
		Name: name,
		Handle: handle,
		Image: pic,
		Default: def,
		PlayerID: p.ID,
		ServerID: server,
	}

	tx := a.db.Create(&c)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &c, nil
}
