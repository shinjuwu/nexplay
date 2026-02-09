## Go 目錄結構

### `/internal`

**私有應用程式**和**函式庫**的程式碼，是你不希望其他人在其應用程式或函式庫中匯入的程式碼。請注意：這個目錄結構是由 Go 編譯器本身所要求的。有關更多細節，請參閱 Go 1.4 的 [`release notes`](https://golang.org/doc/go1.4#internalpackages)。注意：這個目錄並不侷限於放在專案最上層的 `internal` 目錄。事實上，你在專案目錄下的任何子目錄都可以包含 `internal` 目錄。

你可以選擇性的加入一些額外的目錄結構到你的內部套件(`internal package`)中，用來區分你想「共用」與「非共用」的內部程式碼(internal code)。這不是必要的（尤其是對小型專案來說），但有視覺上的線索來表達套件的共用意圖來說，肯定會更好(nice to have)。你的應用程式程式碼可以放在 `/internal/app` 目錄下 (例如：`/internal/app/myapp`)，而這些應用程式共享的程式碼就可以放在 `/internal/pkg` 目錄下 (例如：`/internal/pkg/myprivlib`)。

### `/pkg`

函式庫的程式碼當然可以讓外部應用程式來使用 (例如：`/pkg/mypubliclib`)，其他專案會匯入這些函式庫，並且期待它們能正常運作，所以要把程式放在這個目錄下請多想個幾遍！:-) 注意：使用 `internal` 目錄可以確保私有套件不會被匯入到其他專案使用，因為它是由 Go 的編譯器強制執行的，所以是比較好的解決方案。使用 `/pkg` 目錄仍然是一種很好的方式，它代表其他專案可以安全地使用這個目錄下的程式碼。由 Travis Jeffery 撰寫的 [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) 文章提供了關於 `pkg` 和 `internal` 目錄很好的概述，以及使用它們的時機點。

當專案的根目錄包含許多不是用 Go 所寫的元件與目錄時，將 Go 程式碼放在一個集中的目錄下也是種不錯的方法，這使得運行各種 Go 工具變得更加容易（正如以下這些演講中提到的那樣：來自 GopherCon EU 2018 的 [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg)、[GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) 和 [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)）。

如果你想查看哪些知名的 Go 專案使用本專案的目錄結構，請查看 [`/pkg`](pkg/README.md) 目錄。這是一組常見的目錄結構，但並不是所有人都接受它，有些 Go 社群的人也不推薦使用。

如果你的應用程式專案真的很小，或是套用這些資料夾不會對你有太大幫助（除非你真的很想用XD），不使用本專案推薦的目錄結構是完全沒問題的。當你的專案變的越來越大，根目錄將會會變得越來越複雜（尤其是當你有許多不是 Go 所寫的元件時)，你可以考慮參考這個專案所建議的目錄結構來組織你的程式碼。
