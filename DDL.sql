CREATE TABLE `konsumen` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nik` varchar(20) NOT NULL,
  `full_name` varchar(30) NOT NULL,
  `legal_name` varchar(30) NOT NULL,
  `tempat_lahir` varchar(30) NOT NULL,
  `tanggal_lahir` varchar(30) NOT NULL,
  `gaji` float NOT NULL,
  `foto_ktp` varchar(100) ,
  `foto_selfie` varchar(100) ,
  PRIMARY KEY (`id`)
);

CREATE TABLE `pinjaman_limit` (
  `id` int NOT NULL AUTO_INCREMENT,
  `konsumen_id` int NOT NULL,
  `tenor1` float NOT NULL,
  `tenor2` float NOT NULL,
  `tenor3` float NOT NULL,
  `tenor4` float NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `pinjaman_trx` (
  `id` int NOT NULL AUTO_INCREMENT,
  `no_kontak` varchar(100) NOT NULL,
  `otr` float NOT NULL,
  `nama_asset` varchar(100) NOT NULL,
  `admin_fee` float NOT NULL,
  `jumlah_bunga` float NOT NULL,
  `jumlah_cicilan` float NOT NULL,
  PRIMARY KEY (`id`)
);