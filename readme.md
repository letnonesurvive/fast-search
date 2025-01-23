# Постановка задачи

Отчет по проделанной работе можно найти в папке Report

Есть функиця, которая что-то там ищет по файлу. Но делает она это не очень быстро. Надо её оптимизировать.

Задание на работу с профайлером pprof.

Цель задания - научиться работать с pprof, находить горячие места в коде, уметь строить профиль потребления cpu и памяти, оптимизировать код с учетом этой информации. Написание самого быстрого решения не является целью задания.

Для генерации графа вам понадобится graphviz. Для пользователей windows не забудьте добавить его в PATH чтобы была доступна команда dot.

Рекомендую внимательно прочитать доп. материалы на русском - там ещё много примеров оптимизации и объяснений как работать с профайлером. Фактически там есть вся информация для выполнения этого задания.

Есть с десяток мест где можно оптимизировать.
Вам надо писать отчет, где вы заоптимайзили и что. Со скриншотами и объяснением что делали. Чтобы именно научиться в pprof находить проблемы, а не прикинуть мозгами и решить что вот тут медленно.

Для выполнения задания необходимо чтобы один из параметров ( ns/op, B/op, allocs/op ) был быстрее чем в *BenchmarkSolution* ( fast < solution ) и ещё один лучше *BenchmarkSolution* + 20% ( fast < solution * 1.2), например ( fast allocs/op < 10422*1.2=12506 ).

По памяти ( B/op ) и количеству аллокаций ( allocs/op ) можно ориентироваться ровно на результаты *BenchmarkSolution* ниже, по времени ( ns/op ) - нет, зависит от системы.

Параллелить (использовать горутины) или sync.Pool в это задании не нужно.

Результат в fast.go в функцию FastSearch (изначально там то же самое что в SlowSearch).

Пример результатов с которыми будет сравниваться:
```
$ go test -bench . -benchmem

goos: windows

goarch: amd64

BenchmarkSlow-8 10 142703250 ns/op 336887900 B/op 284175 allocs/op

BenchmarkSolution-8 500 2782432 ns/op 559910 B/op  10422 allocs/op

BenchmarkFast-16    766 1710115 ns/op 622096 B/op  13310 allocs/op
                    686 1673687 ns/op 618502 B/op  11246 allocs/op
                    720 1659494 ns/op 618540 B/op  11246 allocs/op

                    848 1369394 ns/op 564344 B/op  6076 allocs/op - DONE
                    883 1334051 ns/op 564387 B/op  6076 allocs/op
PASS

ok coursera/hw3 3.897s
```

Запуск:
* `go test -v` - чтобы проверить что ничего не сломалось
* `go test -bench . -benchmem` - для просмотра производительности
* `go tool pprof -http=:8083 /path/ho/bin /path/to/out` - веб-интерфейс для pprof, пользуйтесь им для поиска горячих мест. Не забывайте, что у вас 2 режиме - cpu и mem, там разные out-файлы.

Советы:
* Смотрите где мы аллоцируем память
* Смотрите где мы накапливаем весь результат, хотя нам все значения одновременно не нужны
* Смотрите где происходят преобразования типов, которые можно избежать
* Смотрите не только на графе, но и в pprof в текстовом виде (list FastSearch) - там прямо по исходнику можно увидеть где что
* Задание предполагает использование easyjson. На сервере эта библиотека есть, подключать можно. Но сгенерированный через easyjson код вам надо поместить в файл с вашей функцией
* Можно сделать без easyjson

Примечание:
* easyjson основан на рефлекции и не может работать с пакетом main. Для генерации кода вам необходимо вынести вашу структуру в отдельный пакет, сгенерить там код, потом забрать его в main
