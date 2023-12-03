# 3. Use CQRS Pattern

Date: 2023/12/02

## Status

Accepted

## Context

Saya menggunakan CQRS untuk memecah logic yang rumit menjadi lebih sederhana.

## Decision

CQRS, singkatan dari Command Query Responsibility Segregation, adalah sebuah pola desain arsitektur perangkat lunak yang memisahkan operasi membaca (query) dan operasi menulis (command) menjadi dua tanggung jawab terpisah. Pemisahan ini bertujuan untuk meningkatkan kinerja, skalabilitas, dan fleksibilitas dalam pengembangan aplikasi.

Berikut adalah penjelasan singkat tentang konsep CQRS:

1. **Command (Menulis)**:
Operasi yang mengubah status atau data aplikasi.
Contoh: Menambahkan, mengubah, atau menghapus entitas.

1. **Query (Membaca)**:
Operasi yang hanya membaca data tanpa mengubah status atau data aplikasi.
Contoh: Menampilkan daftar entitas atau detail entitas.

1. **Responsibility Segregation**:
Memisahkan tanggung jawab antara operasi membaca dan menulis.
Masing-masing tanggung jawab memiliki model data dan logika bisnisnya sendiri.

1. **Keuntungan CQRS**:
  - Kinerja: Memungkinkan pengoptimalan terhadap operasi membaca dan menulis secara terpisah.
  - Skalabilitas: Meningkatkan kemampuan untuk menangani beban tinggi dengan skalabilitas horizontal.
  - Fleksibilitas: Memungkinkan penggunaan penyimpanan data yang berbeda untuk operasi membaca dan menulis.

1. **Event Sourcing (Opsional)**:
Seringkali digunakan bersama dengan CQRS.
Menggunakan peristiwa (event) sebagai sumber kebenaran untuk semua perubahan data dalam aplikasi.

1. **Kompleksitas Tambahan**:
Penerapan CQRS seringkali memperkenalkan kompleksitas tambahan dalam pengembangan dan pemeliharaan aplikasi.

## Consequences

Untuk mengimplementasikan CQRS membutuhkan setup aplikasi yang lebih banyak, karena harus membuat file berdasarakan command/query.
