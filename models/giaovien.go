package models

import (
	"errors"
	"strings"
)

var (
	dsgv map[string]*GiaoVien
)

type GiaoVien struct {
	MaGV     string   `json:"ma_gv"`
	TenGV    string   `json:"ten_gv"`
	GioiTinh string   `json:"tai_khoan"	`
	QueQuan  string   `json:"mat_khau"`
	DSLop    []string `json:"ds_lop"`
}

func AddGV(gv GiaoVien) string {
	// kiểm tra đã tồn tại lớp chưa
	for _, item := range gv.DSLop {
		if dslop[item] == nil {
			return "lop chua ton tai"
		}

		dslop[item].DSGV = append(dslop[item].DSGV, gv.MaGV)
	}

	dsgv[gv.MaGV] = &gv
	return "them thành công"
}
func init() {
	dsgv = make(map[string]*GiaoVien)
	u := GiaoVien{"123", "toan", "toan", "111", []string{}}
	u1 := GiaoVien{"1234", "dd", "dd", "121", []string{}}
	dsgv["123"] = &u
	dsgv["1234"] = &u1
}

func GetGV(uid string) (gv *GiaoVien, err error) {
	if gv, ok := dsgv[uid]; ok {
		return gv, nil
	}
	return nil, errors.New("teacher not exists")

}
func FindGV(input string) []*GiaoVien {
	keys := make([]*GiaoVien, 0, len(dsgv))
	for _, k := range dsgv {
		println(strings.Contains)
		if strings.Contains(k.MaGV, input) == true || strings.Contains(k.TenGV, input) == true {
			keys = append(keys, k)
		}
	}
	return keys
}
func GetAllGV() []*GiaoVien {
	values := []*GiaoVien{}
	for _, value := range dsgv {
		values = append(values, value)
	}
	return values
}

func DeleteGV(maGV string) string {
	gv := &GiaoVien{}
	gv = dsgv[maGV]
	// kiểm tra giáo viên đã tồn tại chưa
	if gv == nil {
		return "giao vien chua ton tai"
	} else {
		//xóa giao viên trong lớp nếu có
		for _, item := range dsgv[maGV].DSLop {
			//lop cần xóa giao viên
			gv_token := dslop[item]
			gv_token.DSGV = DeleteItemInArray(maGV, gv_token.DSGV)
		}
		//xóa giáo viên trong dsgv
		for i, _ := range dsgv {
			if gv == dsgv[maGV] {
				delete(dsgv, i)
			}
		}
	}
	return "xóa giao vien thành công"
}
func UpdateGV(uid string, us *GiaoVien) (a *GiaoVien, err error) {
	if u, ok := dsgv[uid]; ok {

		if us.TenGV != "" {
			u.TenGV = us.TenGV
		}
		if us.GioiTinh != "" {
			u.GioiTinh = us.GioiTinh
		}
		if us.QueQuan != "" {
			u.QueQuan = us.QueQuan
		}

		return u, nil
	}
	return nil, errors.New("User Not Exist")
}
