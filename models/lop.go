package models

import (
	"log"
)

var (
	dslop map[string]*Lop
)

func init() {
	dslop = make(map[string]*Lop)
	u := Lop{"123", "VatLy1", []string{}, []string{}}
	u1 := Lop{"1234", "VatLy2", []string{}, []string{}}
	dslop["123"] = &u
	dslop["1234"] = &u1

}

type CrudST struct {
	MaLop string `json:"ma_lop"`
	MaSV  string `json:"ma_sv"`
}
type CrudGV struct {
	MaLop string `json:"ma_lop"`
	MaGV  string `json:"ma_gv"`
}

type Lop struct {
	MaLop  string   `json:"ma_lop"`
	TenLop string   `json:"ten_lop"`
	DSSV   []string `json:"dssv"`
	DSGV   []string `json:"dsgv"`
}

func GetClass(uid string) []*Lop {
	values := []*Lop{}
	for _, value := range dslop {
		if value == dslop[uid] {
			values = append(values, value)
		}
	}
	return values
}
func GetAllClass() []*Lop {
	values := []*Lop{}
	for _, value := range dslop {
		values = append(values, value)
	}
	return values
}

func AddClass(lp Lop) string {

	for _, item := range lp.DSSV {
		if dssv[item] == nil {
			return "Sinh vien chua ton tai"
		}
		// thêm lop vào sv
		dssv[item].DSLop = append(dssv[item].DSLop, lp.MaLop)
	}

	for _, item := range lp.DSGV {
		if dsgv[item] == nil {
			return "Giao vien chua ton tai"
		}
		// thêm lớp vào gv
		dsgv[item].DSLop = append(dsgv[item].DSLop, lp.MaLop)
	}

	dslop[lp.MaLop] = &lp
	return lp.MaLop
}

func UpdateClass(lopId string, lp *Lop) string {
	// check cac danh sach
	// chek ds sv <ca dssv_them va dssv_xoa sv neu co>
	dssv_xoa := make([]bool, 0)
	dssv_them := make([]bool, 0)
	dssv_xoa = CheckListAdd_Delete(dslop[lopId].DSSV, lp.DSSV)
	dssv_them = CheckListAdd_Delete(lp.DSSV, dslop[lopId].DSSV)
	//lấy ra mã sinh viên cần thêm
	dssvLop := lp.DSSV
	for index, value := range dssv_xoa {
		if value == false {
			DeleteSTInClass(dslop[lopId].MaLop, dslop[lopId].DSSV[index])
		}
	}
	for key, value := range dssv_them {
		if value == false {
			for _, v := range dssvLop {
				if dssv[v] != nil {
					re := AddSTToClass(dslop[lopId].MaLop, dssvLop[key])
					if re == "sinh viên đã có trong lớp" {
						return "sinh vien da ton tai"
					}
				} else {
					return "sinh vien chua ton tai"
				}
			}
		}
	}
	log.Println("dssvlop: ", dssvLop)
	// check ds gv
	dsgv_xoa := make([]bool, 0)
	dsgv_them := make([]bool, 0)
	dsgv_xoa = CheckListAdd_Delete(dslop[lopId].DSGV, lp.DSGV)
	dsgv_them = CheckListAdd_Delete(lp.DSGV, dslop[lopId].DSGV)
	// lấy ra mã giáo viên cần thêm
	dsgvLop := lp.DSGV
	for index, value := range dsgv_xoa {
		if value == false {
			DeleteSTInClass(dslop[lopId].MaLop, dslop[lopId].DSGV[index])
		}
	}
	for key, value := range dsgv_them {
		if value == false {
			for _, v := range dsgvLop {
				if dsgv[v] != nil {
					AddSTToClass(dslop[lopId].MaLop, dsgvLop[key])
				} else {
					return "giao vien chua ton tai"
				}
			}
		}
	}
	// cho thay the
	dslop[lopId] = lp
	return "them thanh cong"
}

func AddSTToClass(maLop string, maSV string) string {

	lop_byid := &Lop{}
	lop_byid = dslop[maLop]
	// kiem tra lop ton tai chua => lay op ra
	if lop_byid == nil {
		return "Lớp chưa tồn tại"
	} else {
		log.Println(dssv)
		sv := &SinhVien{}
		sv = dssv[maSV]
		// add sinh vien vao lop
		if sv == nil {
			// sv da ton tai chua
			return "sinh viên chưa tồn tại"
		} else {
			// kiem tra sinh vien da co trong lop chua
			for _, item := range lop_byid.DSSV {
				if item == maSV {
					return "sinh viên đã có trong lớp"
				}
			}
			lop_byid.DSSV = append(lop_byid.DSSV, sv.MaSV)
			sv.DSLop = append(sv.DSLop, lop_byid.MaLop)
			return "thêm thành công"
		}
	}

}

func AddGVToClass(maLop string, maGV string) string {
	// kiem tra ma lop ton tai chua
	lp := &Lop{}
	lp = dslop[maLop]
	if lp == nil {
		return "lop chưa tồn tại"
	} else {
		// kiem tra giao vien da ton tai chua
		gv := &GiaoVien{}
		gv = dsgv[maGV]
		if gv == nil {
			return "giáo viên chưa tồn tại"
		} else {
			//kiem tra giao vien da co trong lop chua
			log.Println("dsgv: ", lp.DSGV)
			for _, item := range lp.DSGV {
				if item == maGV {
					return "giáo viên đã có trong lớp"
				}
			}
			lp.DSGV = append(lp.DSGV, gv.MaGV)
			gv.DSLop = append(gv.DSLop, lp.MaLop)
			return "thêm thành công"
		}
	}
}

func DeleteSTInClass(maLop string, maSV string) string {

	lp := &Lop{}
	lp = dslop[maLop]
	// kiem tra lop da ton tai chua
	if lp == nil {
		return "Lớp chưa tồn tại"

	} else {
		sv := &SinhVien{}
		sv = dssv[maSV]
		//Arr := []string{}
		for _, item := range lp.DSSV {
			//kiem tra sinh vien co trong lop chua
			if item == maSV {
				lp.DSSV = DeleteItemInArray(maSV, lp.DSSV)
				sv.DSLop = DeleteItemInArray(maSV, sv.DSLop)
				return "xóa thành công"
			}
		}
		return "Sinh viên chưa có trong lớp "
	}
}

func DeleteTeacherInClass(maLop string, maGV string) string {
	lp := &Lop{}
	lp = dslop[maLop]
	// kiem tra lop da ton tai chua
	if lp == nil {
		return "Lớp chưa tồn tại"
	} else {
		gv := &GiaoVien{}
		gv = dsgv[maGV]
		for _, item := range lp.DSGV {
			//kiem tra giáo viên có trong lớp chưa
			if item == maGV {
				// xóa giáo viên trong dsgv
				lp.DSGV = DeleteItemInArray(maGV, lp.DSGV)
				//xóa gv trong dslop
				gv.DSLop = DeleteItemInArray(maGV, gv.DSLop)

				return "xóa thành công"
			}
		}
		return "Giáo viên chưa có trong lớp"
	}
}

func DeleteClass(maLop string) string {
	lp := &Lop{}
	lp = dslop[maLop]
	// kiểm tra lớp tồn tại chưa
	if lp == nil {
		return "Lớp chưa tồn tại"
	} else {
		//xoa lơp trong sinh vien
		for _, item := range dslop[maLop].DSSV {
			// sinh viên cần xóa
			sv_token := dssv[item]
			//sv_token.DSLop
			sv_token.DSLop = DeleteItemInArray(maLop, sv_token.DSLop)
		}
		//xóa lớp trong giáo viên
		for _, item := range dslop[maLop].DSGV {
			// sinh viên cần xóa
			gv_token := dsgv[item]
			//sv_token.DSLop
			gv_token.DSLop = DeleteItemInArray(maLop, gv_token.DSLop)
		}
		//xoa lop trong danh sach lop
		for i, _ := range dslop {
			if lp == dslop[maLop] {
				delete(dslop, i)
			}
		}
	}

	return "xóa lớp thành công"
}

func DeleteItemInArray(item string, Arr []string) []string {

	A := []string{}
	for _, k := range Arr {
		if k != item {
			A = append(A, k)
		}
	}
	return A
}
func CheckListAdd_Delete(Arr []string, ArrInput []string) []bool {
	ds := make([]bool, 0)
	for i1, v1 := range Arr {
		ds = append(ds, false)
		for _, v2 := range ArrInput {
			if v1 == v2 {
				ds[i1] = true
				break
			}
		}
	}
	return ds
}
