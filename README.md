# Menu API

## การติดตั้งและการใช้งาน

### ข้อกำหนดระบบ

โปรเจกต์นี้ต้องการ:

- Go 1.19 ขึ้นไป
- MySQL หรือ MariaDB
- Git (ถ้าต้องการ Clone โปรเจกต์จาก Git repository)

### ขั้นตอนการติดตั้ง

1. Clone โปรเจกต์จาก Git repository

   ````bash
   git clone https://github.com/username/menu-api.git
   cd menu-api```

   ````

2. ติดตั้ง dependencies ใช้คำสั่งต่อไปนี้เพื่อติดตั้ง dependencies

`go mod tidy`

3. ตั้งค่าฐานข้อมูล MySQL

`dsn := "user:password@tcp(127.0.0.1:3307)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"`

4. รันโปรเจกต์

`go run main.go`

5. ทดสอบ api

`curl http://localhost:8080/menus`
