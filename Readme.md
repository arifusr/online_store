Flash sale

terjadinya kesalahan stok atau stok minus bisa terjadi dikarenakan tidak adanya handle untuk race condition ketika user berbarengan membeli atau checkout suatu produk

dengan menangani race condition bisa menjaga stok untuk tetap akurat
race condition bisa diatasi dengan lock database ataupun menggunakan metode optimistic locking yang mana sebelum mengupdate atau mengubah suatu data diharuskan untuk mengecek version dan dibandingkan dengan version yg ada jika tidak sesuai maka ada yang telah lebih dahulu merubahnya sehingga dibutuhkan perhitungan ulang