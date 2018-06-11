package neteaseMusicApi

import "strconv"

//Search
func (search *Search)SetKeyWord(keyword string){
	search.S = keyword
}

func (search *Search)SetOffset(offset int){
	search.Offset = offset
}

func (search *Search)SetLimit(limit int){
	search.Limit = limit
}

func (search Search)GetJson() string{
	return doPost(SearchUrl,search.GetParams())
}

func (search Search) GetParams() string{
	return AesEncrypt(AesEncrypt(search.EncStr(),nonce),randStr)
}

func NewSearch() Search{
	return Search{
		S : "",
		Type:1,
		Offset:0,
		Total:true,
		Limit:30,
		Csrf_token:"",
	}
}

//Comments
func (comments *Comments)SetId(id int){
	comments.Rid="R_SO_4_" + strconv.Itoa(id)
}

func (comments *Comments)SetOffset(offset int){
	comments.Offset = offset
}

func (comments *Comments)SetLimit(limit int){
	comments.Limit = limit
}

func (comments Comments)GetJson() string{
	return doPost(CommentUrl+comments.Rid,comments.GetParams())
}

func (comments Comments)GetParams() string{
	return AesEncrypt(AesEncrypt(comments.EncStr(),nonce),randStr)
}

func NewComments() Comments{
	return Comments{
		Rid : "",
		Offset:0,
		Total:true,
		Limit:0,
		Csrf_token:"",
	}
}

//Detail

func (detail *Detail)SetId(id int){
	detail.Id = id
	detail.C[0].Id=id
	//detail.Ids[0] = id
}

func (detail Detail)GetParams() string{
	return AesEncrypt(AesEncrypt(detail.EncStr(),nonce),randStr)
}

func (detail Detail)GetJson() string{
	return doPost(DetailUrl,detail.GetParams())
}

func NewDetail()Detail{
	return Detail{
		Id:0,
		C:[]C{
			C{
				Id:0,
			},
		},
		//Ids:0,
	}
}

//SongMp3
func (songMp3 *SongMP3)SetIds(ids []int){
	songMp3.Ids = ids
}

func (songMp3 *SongMP3)SetBr(br int){
	songMp3.Br = br
}

func (songMp3 *SongMP3)GetJson() string{
	return doPost(SongMp3Url,songMp3.GetParams())
}

func (songMp3 *SongMP3) GetParams() string{
	return AesEncrypt(AesEncrypt(songMp3.EncStr(),nonce),randStr)
}

func NewSongMP3() SongMP3{
	return SongMP3{
		Ids : []int{},
		Br : 320000,
		Csrf_token : "",
	}
}

//Lyric
func (lyric *Lyric)SetId(id int){
	lyric.Id = id
}

func (lyric *Lyric)GetJson() string{
	return doPost(LyricUrl,lyric.GetParams())
}

func (lyric Lyric) GetParams() string{
	return AesEncrypt(AesEncrypt(lyric.EncStr(),nonce),randStr)
}

func NewLyric() Lyric{
	return Lyric{
		Id : 0,
		Os : "linux",
		Lv : -1,
		Kv : -1,
		Tv : -1,
	}
}


