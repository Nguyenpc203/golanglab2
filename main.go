package main

import (
	"fmt"
	"strings"
)

// Struct sinh viên
type Student struct {
	ID     string
	Name   string
	Age    int
	Gender string // Nam hoặc Nữ
}

// Khai báo map để lưu trữ sinh viên
var students = make(map[string]Student)

// Thêm sinh viên vào map
func addStudent(id string, name string, age int, gender string) {
	gender = strings.ToLower(gender)
	if gender != "nam" && gender != "nữ" {
		fmt.Println("Giới tính không hợp lệ. Chỉ chấp nhận 'Nam' hoặc 'Nữ'")
		return
	}
	students[id] = Student{
		ID:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
	}
	fmt.Println("Thêm sinh viên thành công!")
}

// Tìm kiếm sinh viên theo ID
func findStudent(id string) {
	student, exists := students[id]
	if exists {
		fmt.Printf("Thông tin sinh viên: ID: %s, Tên: %s, Tuổi: %d, Giới tính: %s\n", student.ID, student.Name, student.Age, student.Gender)
	} else {
		fmt.Println("Không tìm thấy sinh viên.")
	}
}

// Sửa thông tin sinh viên
func updateStudent(id string, name string, age int, gender string) {
	_, exists := students[id]
	if exists {
		students[id] = Student{
			ID:     id,
			Name:   name,
			Age:    age,
			Gender: strings.ToLower(gender),
		}
		fmt.Println("Cập nhật sinh viên thành công!")
	} else {
		fmt.Println("Không tìm thấy sinh viên để cập nhật.")
	}
}

// Xóa sinh viên theo ID
func deleteStudent(id string) {
	_, exists := students[id]
	if exists {
		delete(students, id)
		fmt.Println("Xóa sinh viên thành công!")
	} else {
		fmt.Println("Không tìm thấy sinh viên để xóa.")
	}
}

// Xuất danh sách sinh viên theo giới tính
func listStudentsByGender(gender string) {
	gender = strings.ToLower(gender)
	fmt.Printf("Danh sách sinh viên giới tính %s:\n", gender)
	for _, student := range students {
		if student.Gender == gender {
			fmt.Printf("ID: %s, Tên: %s, Tuổi: %d\n", student.ID, student.Name, student.Age)
		}
	}
}

func main() {
	// Thêm một số sinh viên mẫu
	addStudent("SV001", "Nguyễn Văn A", 20, "Nam")
	addStudent("SV002", "Trần Thị B", 21, "Nữ")
	addStudent("SV003", "Phạm Văn C", 22, "Nam")
	addStudent("SV004", "Lê Thị D", 19, "Nữ")

	// Tìm kiếm sinh viên
	findStudent("SV001")

	// Cập nhật sinh viên
	updateStudent("SV001", "Nguyễn Văn A", 21, "Nam")

	// Xóa sinh viên
	deleteStudent("SV003")

	// Xuất danh sách sinh viên là Nam
	listStudentsByGender("Nam")

	// Xuất danh sách sinh viên là Nữ
	listStudentsByGender("Nữ")
}
