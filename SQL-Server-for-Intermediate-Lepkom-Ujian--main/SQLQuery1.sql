/*create & use database*/
create database lentera_food_50421119;
use lentera_food_50421119;

/*create table*/
CREATE TABLE Persediaan_Barang(
ID_Barang VARCHAR (10) PRIMARY KEY,
Nama_Barang VARCHAR (50),
Jumlah VARCHAR (50),
Tanggal_Masuk DATE,
Tanggal_Keluar DATE);

CREATE TABLE Data_Pegawai(
NIP VARCHAR (20) PRIMARY KEY,
NAMA VARCHAR (30),
TEMPAT_LAHIR VARCHAR (10),
TGL_LAHIR DATE,
ALAMAT VARCHAR (20),
DIVISI VARCHAR (10),
GAJI VARCHAR (10));

SELECT * FROM Persediaan_Barang;
SELECT * FROM Data_Pegawai;

/*INSERT DATA*/
INSERT INTO Persediaan_Barang VALUES
('A0201', 'Tepung Jagung', '1 Ton', '2016-12-10', '2016-12-15'),
('A0202', 'Minyak Goreng', '1000 Liter', '2016-12-11', '2016-12-20'),
('F0203', 'Gula', '4 Ton', '2016-12-05', '2016-12-10'),
('S0201', 'Gandum', '3 Ton', '2016-12-08', '2016-12-10'),
('D0202', 'Tepung Beras', '3 Ton', '2016-12-10', '2016-12-16');

INSERT INTO Data_Pegawai VALUES
('53210985', 'Dio Susilo', 'Bogor', '1992-08-16', 'Jl. Kenari', 'Accounting', '2000'), 
('18111525', 'Hardi', 'Jakarta', '1993-07-11', 'Jl. Lereng', 'HRD', '1500'),
('52409380', 'Fadel Andreza', 'Depok', '1992-08-12', 'Jl. Melati', 'Staff IT', '2000'),
('18110000', 'Firmansyah', 'Bogor', '1993-11-10', 'Jl. Dahlia', 'Finance', '1000'),
('19283332', 'Dimas Hadi', 'Bekasi', '1994-09-11', 'Jl. Delima', 'HRD', '1000');

/*ERROR MASSAGE*/
SP_ADDMESSAGE 50099, 16, 
'Mohon maaf, Tanggal yang anda input lebih kecil dari tanggal masuk silahkan anda perbaiki', 
'US_ENGLISH', 'TRUE';

SELECT*FROM SYS.MESSAGES;

/*PESAN ERROR TERTENTU 1*/
BEGIN TRY RAISERROR (50099, 16, 1) WITH LOG 
END TRY 
BEGIN CATCH 
SELECT ERROR_MESSAGE(), ERROR_NUMBER()
END CATCH;

/*PESAN ERROR TERTENTU 2*/
CREATE VIEW v_data_peg
as select * from Data_Pegawai
where DIVISI = 'hrd';

ALTER VIEW v_data_peg as select nip,nama as ename,
tempat_lahir, tgl_lahir,alamat as address, DIVISI as division
from Data_Pegawai;

select * from v_data_peg;

/*SQL EXTENTION*/
if (select count(*) from Data_Pegawai
where Gaji < 1500 ) > 2
print 'Gaji perusahaan masih dibawah standar'
else begin
print 'Berikut adalah karyawan yang memiliki gaji diatas standar'
select Nama, Gaji from Data_Pegawai
where Gaji > 1500
end

select * from Data_Pegawai where Gaji > 1500;