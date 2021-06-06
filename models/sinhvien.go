package models

import (
	"errors"
	"strings"
)

var (
	dssv map[string]*SinhVien
)

type SinhVien struct {
	MaSV     string   `json:"ma_sv"`
	TenSV    string   `json:"ten_sv"`
	GiơiTinh string   `json:"gioi_tinh"`
	QueQuan  string   `json:"que_quan"`
	DSLop    []string `json:"ds_lop"`
}

type SV struct {
	MaSV     string `json:"ma_sv"`
	TenSV    string `json:"ten_sv"`
	GiơiTinh string `json:"gioi_tinh"`
	QueQuan  string `json:"que_quan"`
}

func AddST(st SinhVien) string {
	// kiểm tra danh sach lop tôn tai hay ko
	/*st.DSLop*/
	for _, item := range st.DSLop {
		if dslop[item] == nil {
			return "lop chua ton tai"
		}
		// thêm sv vào lớp
		dslop[item].DSSV = append(dslop[item].DSSV, st.MaSV)
	}
	dssv[st.MaSV] = &st
	return "them thành công"
}
func init() {
	dssv = make(map[string]*SinhVien)
	u := SinhVien{"123", "Nguyễn Tuấn Toàn", "Nam", "MC", []string{}}
	u1 := SinhVien{"1234", "Nguyễn Tuấn Việt", "Nam", "QN", []string{}}
	dssv["123"] = &u
	dssv["1234"] = &u1
}

func GetST(uid string) (st *SinhVien, err error) {
	if st, ok := dssv[uid]; ok {
		return st, nil
	}
	return nil, errors.New("SinhVien not exists")
}

func GetAllST() []*SinhVien {
	values := []*SinhVien{}
	for _, value := range dssv {
		values = append(values, value)
	}
	return values
}

/*func GetListST() []*SinhVien {
	keys := make([]*SinhVien, 0, len(dssv))
	for _,k := range dssv {
		keys = append(keys, k)
	}
	return keys
}*/

func FindStudent(input string) []*SinhVien {
	keys := make([]*SinhVien, 0, len(dssv))
	for _, k := range dssv {
		println(strings.Contains(k.TenSV, input))
		if strings.Contains(k.MaSV, input) == true || strings.Contains(k.TenSV, input) == true {
			keys = append(keys, k)
		}
	}
	return keys
}

func DeleteST(maSV string) string {

	sv := &SinhVien{}
	sv = dssv[maSV]
	// kiem tra sinh vien tồn tại chưa
	if sv == nil {
		return "sinh viên chưa tồn tại"
	} else {
		//xóa sv trong lớp nếu có
		for _, item := range dssv[maSV].DSLop {
			// lớp cần xóa sinh viên
			lp_token := dslop[item]
			lp_token.DSSV = DeleteItemInArray(maSV, lp_token.DSSV)
		}
		// xóa sv trong dssv
		for i, _ := range dssv {
			if sv == dssv[maSV] {
				delete(dssv, i)
			}
		}
	}
	return "xóa sinh viên thành công"
}

func UpdateST(uid string, us *SinhVien) (a *SinhVien, err error) {
	if u, ok := dssv[uid]; ok {
		if us.GiơiTinh != "" {
			u.GiơiTinh = us.GiơiTinh
		}
		if us.QueQuan != "" {
			u.QueQuan = us.QueQuan
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}
