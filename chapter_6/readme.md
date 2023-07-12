
1. <b>WaitGroup</b> - механизм ожидания завершения группы задач
wg.Add() // кол-во задач которе надо выполнить
wg.Done() // наш такс который мы создали надо обязательно завершить
wg.Wait() // блокирует основной поток, чтобы наша задача которую мы создали выполнилась

2. <b>DataRace</b> - обращение к одним и тем же данным из разных программ/тредов/горутин, где одно из обращений - запись

3. <b>Mutex и RWMutex</b> - механизм получения исключительной блокировки
m.Lock()
m.Unlock()

<b>Channels</b>
type chan struct {
    mx sync.mutex
    buffer []T
    readers []Goroutines
    writers []Goroutines
}

I. Nil channel
    1. Создание
    2. len и cap
    3. Запись/получение значения: <-
    4. Закрытие канала: close

II. Unbuffered channel
    1. Создание
    2. len и cap
    3. Запись/получение значения: <-

  unbufferedChannel {
    len(buffer) = 0
    readers []Goroutines
    writers []Goroutines
  }

4. Направленность канала
5. Закрытие канала: close. Всегда закрывать отправителем

III. Buffered channel
    1. Создание
    2. len и cap
    3. Запись/получение значения: <-

bufferedChannel {
    len(buffer) > 0;   []
    readers []Goroutines
    writers []Goroutines
  }

    4. Закрытие канала: close. Всегда закрывать отправителем

IV: Проверка на закрытие
V: For..range


<b> Select, Graceful shutdown</b>
1. select - block, unblock, default
2. time.After
3. graceful shutdown