# MyGram

MyGram adalah layanan backend yang dibangun menggunakan Golang, menyediakan API RESTful untuk fungsi-fungsi yang mirip dengan Instagram. Aplikasi ini memungkinkan pengguna untuk mendaftar, login, mengunggah foto, berkomentar, dan berinteraksi dengan media sosial dalam sebuah ekosistem virtual.

## Fitur

- **Pendaftaran Pengguna**: Memungkinkan pengguna baru untuk membuat akun.
- **Login**: Pengguna dapat login ke dalam sistem.
- **Model Foto**: Pengguna dapat mengunggah dan melihat foto.
- **Komentar**: Pengguna dapat berkomentar pada foto.
- **Sosial Media**: Interaksi sosial antar pengguna.

## URL
https://mygram-production-69f8.up.railway.app/

## Endpoint API 
- **GET /users/get** - Menampilkan seluruh user yang terdaftar.
- **POST /users/register** - Daftar akun user.
- **POST /users/login** - Login ke sistem.
- **PUT /users/edit/{id}** - Merubah data pada user.
- **DELETE /users/delete/{id}** - Menghapus user.

- **POST /photos/post** - Membuat data foto.
- **POST /photos/get** - Menampilkan foto.
- **PUT /photos/edit/{photoId}** - Merubah data pada foto.
- **DELETE /photos/delete/{photoId}** - Menghapus foto.

- **POST /comments/post** - Membuat data komen.
- **POST /comments/get** - Menampilkan komen.
- **PUT /comments/edit/{photoId}** - Merubah data pada komen.
- **DELETE /comments/delete/{photoId}** - Menghapus komen.

- **POST /socialmedias/post** - Membuat data social media.
- **POST /socialmedias/get** - Menampilkan social media.
- **PUT /socialmedias/edit/{photoId}** - Merubah data pada social media.
- **DELETE /socialmedias/delete/{photoId}** - Menghapus social media.
