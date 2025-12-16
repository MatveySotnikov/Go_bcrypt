# Практическая работа 9
## Go и Bcrypt
### Выполнил Сотников М.Е. ЭФМО-01-25

Необходимые пакеты:
```ps
go get github.com/go-chi/chi/v5
go get gorm.io/gorm gorm.io/driver/postgres
go get golang.org/x/crypto/bcrypt
```

<img width="890" height="641" alt="Registration" src="https://github.com/user-attachments/assets/abe4de1b-d07c-4b0c-b1ef-0c96e209e427" />   

Попытка регистрации с таким же email (409):    
<img width="914" height="596" alt="SameMail" src="https://github.com/user-attachments/assets/67c2b9e3-c187-4e86-ae4d-654ad60e8597" />    

Успешный вход:    

<img width="842" height="688" alt="image" src="https://github.com/user-attachments/assets/9e2e6618-5562-4183-8b0e-9653d37e9f6e" />

Неверный логин или пароль:    
<img width="834" height="539" alt="err_cred" src="https://github.com/user-attachments/assets/acf35463-cdc6-4199-868a-886c7993eadc" />

ВАЖНО: Никогда нельзя говорить пользователю: «email найден, но пароль неверный». Сообщаем только обобщённо: «Неверный логин или пароль». Это защищает от перебора email-адресов.

Хэширование в коде:   

```go
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), h.BcryptCost)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "hash_failed")
		return
	}
```
* Хэш — это результат работы функции, которую нельзя "развернуть" обратно. Мы храним результат, и при входе сравниваем хэш введенного пароля с тем, что в базе.
* Без соли одинаковые пароли (например, "123456") имеют одинаковые хэши (риск использования Радужных Таблиц для взлома). Соль делает каждый хэш уникальным, даже если пароли одинаковые.
* На современном железе можно перебирать миллионы SHA-хэшей в секунду. bcrypt специально спроектирован как медленный алгоритм.
* Параметр cost определяет количество итераций (циклов) хэширования (в данной работе = 12), повышение cost увеличивает время вычисления хэша, что лучше защищает от перебора, но нагружает сервер.







