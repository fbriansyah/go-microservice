# 2. Use a DDD

Date: 2023/12/02

## Status

Accepted

## Context

Menggunakan DDD untuk membangun microservice agar aplikasi mudah untuk di-*scaling*.

## Decision

Saya menggunakan domain driven development dengan structur folder sebagai berikut:

```
.
|-- cmd
|   `-- api
|-- docs
|   `-- ADL
|-- internal
|   |-- application
|   |   |-- commands
|   |   |-- domain
|   |   |   `-- user
|   |   `-- queries
|   |-- echo
|   `-- postgres
`-- util
```
Di dalam folder internal merupakan inti dari aplikasi. Seluruh domain berada di dalam folder application/domain. 
Untuk detail kenapa menggunakan DDD, berikut ini penjelasan singkat dari DDD.

### DDD
Domain-Driven Development (DDD) adalah suatu pendekatan pengembangan perangkat lunak yang berfokus pada pemahaman yang mendalam terhadap domain bisnis yang menjadi fokus proyek. Beberapa keuntungan dari penerapan Domain-Driven Development termasuk:

1. **Pemahaman Domain yang Lebih Baik**:
DDD membantu tim pengembang dan pemangku kepentingan untuk mendapatkan pemahaman yang lebih baik tentang bisnis dan domain yang menjadi fokus proyek. Dengan begitu, solusi yang dihasilkan dapat lebih sesuai dengan kebutuhan sebenarnya.

1. **Komunikasi yang Meningkat**:
DDD menyediakan bahasa yang kaya dan ekspresif (Ubiquitous Language) yang digunakan oleh semua pihak terkait, termasuk pengembang, analis bisnis, dan pemilik produk. Hal ini meningkatkan komunikasi dan meminimalkan kesalahpahaman antara berbagai pemangku kepentingan.

1. **Model yang Konsisten**:
Dengan fokus pada domain, DDD mendorong pembentukan model yang konsisten dan memadai untuk merepresentasikan realitas bisnis. Model ini dapat digunakan sebagai panduan untuk pengambilan keputusan dan pengembangan perangkat lunak.

1. **Desain yang Lebih Baik**:
DDD mendorong pemisahan perangkat lunak ke dalam lapisan-lapisan (layers) yang terkait dengan domain bisnis. Hal ini dapat menghasilkan desain yang lebih bersih, terstruktur, dan mudah dimengerti.

1. **Fleksibilitas dan Adaptabilitas**:
DDD memungkinkan perubahan dalam domain bisnis dapat diakomodasi dengan lebih baik. Model-domain yang baik dirancang memungkinkan adaptasi lebih mudah terhadap perubahan dalam bisnis atau persyaratan aplikasi.

1. **Meningkatkan Kualitas dan Produktivitas**:
Dengan pemahaman domain yang lebih baik, pengembang dapat membuat solusi perangkat lunak yang lebih relevan dan akurat. Hal ini dapat meningkatkan kualitas perangkat lunak dan produktivitas pengembangan.

1. **Fokus pada Nilai Bisnis**:
DDD membantu tim untuk fokus pada fitur dan aspek yang memberikan nilai bisnis. Ini dapat membantu menghindari pemborosan sumber daya pada fitur atau fungsi yang kurang penting.

1. **Pemisahan Kompleksitas**:
DDD membantu memisahkan kompleksitas teknis dari kompleksitas bisnis. Ini memudahkan pengembang untuk fokus pada solusi teknis tanpa kehilangan pandangan terhadap masalah bisnis yang sebenarnya.
Namun, penting untuk diingat bahwa penerapan DDD bukanlah solusi ajaib dan dapat memerlukan pemahaman dan komitmen yang baik dari seluruh tim pengembangan.

Namun, penting untuk diingat bahwa penerapan DDD bukanlah solusi ajaib dan dapat memerlukan pemahaman dan komitmen yang baik dari seluruh tim pengembangan.

### Note:
- Folder internal detail [/internal](https://go.dev/doc/go1.4#internalpackages).

## Consequences

- Proses development yang cukup complex.
- Butuh pemahan yang baik tentang DDD sebelum bisa menerapkannya.
