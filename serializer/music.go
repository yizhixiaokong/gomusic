package serializer

import "gomusic/model"

//  获取音乐列表序列化器
type GetMusicList struct {
	Total int     `json:"total"`
	List  []Music `json:"list"`
}

// 失败列表
type FailMusicList struct {
	Success  int     `json:"success"`
	FailList []Music `json:"failList"`
}

type Music struct {
	ID               uint   `json:"id"`               //id
	MusicName        string `json:"musicName"`        //歌名
	OriginalSinger   string `json:"originalSinger"`   //原唱
	Language         string `json:"language"`         //语言
	MusicType        string `json:"musicType"`        //分类
	ProficiencyLevel string `json:"proficiencyLevel"` //熟练度
	DesignateBy      string `json:"designateBy"`      //指定者
}

// BuildMusic 序列化歌曲
func BuildMusic(item model.Music) (res Music) {
	return Music{
		ID:               item.ID,
		MusicName:        item.MusicName,
		OriginalSinger:   item.OriginalSinger,
		Language:         item.Language,
		MusicType:        item.MusicType,
		ProficiencyLevel: item.ProficiencyLevel,
		DesignateBy:      item.DesignateBy,
	}
}

// BuildMusicList 序列化歌曲列表
func BuildMusicList(items []model.Music, total int) (res GetMusicList) {
	var musics []Music
	for _, item := range items {
		music := BuildMusic(item)
		musics = append(musics, music)
	}

	res.Total = total
	res.List = musics
	return res
}
