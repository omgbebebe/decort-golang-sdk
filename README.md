# Decort SDK

Decort SDK - это библиотека, написанная на языке GO, позволяющая взаимодействовать с API облачной платформы **DECORT**. Библиотека содеражит в себе структуры и методы, необходимые для отправки запросов. Decort SDK имеет встроенный http-клиент и поддерживает разные способы авторизации на платформе. Библиотека так же содержит в себе модели ответов от платформы.

## Версии

    - Версия 1.0.x Decort-SDK соответствует 3.8.4 версии платформы
    - Версия 1.1.x Decort-SDK соответствует 3.8.5 версии платформы
    - Версия 1.2.x Decort-SDK соответствует 3.8.5 версии платформы
    - Версия 1.3.x Decort-SDK соответствует 3.8.5 версии платформы
    - Версия 1.4.x Decort-SDK соответствует 3.8.6 версии платформы
    - Версия 1.5.x Decort-SDK соответствует 3.8.7 версии платформы
    - Версия 1.6.x Decort-SDK соответствует 3.8.8 версии платформы

## Оглавление

- [Установка](#установка)
- [Список API](#список-api)
  - [Cloudapi](#cloudapi)
  - [Cloudbroker](#cloudbroker)
- [Работа с библиотекой](#работа-с-библиотекой)
  - [Настройка конфигурации клиента](#настройка-конфигурации-клиента)
    - [Пример конфигурации клиента](#пример-конфигурации-клиента)
    - [Парсинг конфигурации из файла](#парсинг-конфигурации-из-файла)
    - [Пример JSON конфигурации](#пример-json-конфигурации)
    - [Пример YAML конфигурации](#пример-yaml-конфигурации)
  - [Создание клиента](#создание-клиента)
  - [Создание структуры запроса](#cоздание-структуры-запроса)
  - [Выполнение запроса](#выполнение-запроса)
  - [Фильтрация](#фильтрация)
  - [Сортировка](#сортировка)
  - [Сериализация](#сериализация)
- [Работа с legacy клиентом](#работа-с-legacy-клиентом)
  - [Настройка конфигурации legacy клиента](#настройка-конфигурации-legacy-клиента)
    - [Пример конфигурации legacy клиента](#пример-конфигурации-legacy-клиента)
    - [Парсинг legacy конфигурации из файла](#парсинг-legacy-конфигурации-из-файла)
    - [Пример legacy JSON конфигурации](#пример-legacy-json-конфигурации)
    - [Пример legacy YAML конфигурации](#пример-legacy-yaml-конфигурации)
  - [Создание legacy клиента](#создание-legacy-клиента)
  - [Создание структуры запроса](#cоздание-структуры-запроса)
  - [Выполнение запроса](#выполнение-запроса)

## Установка

Выполните команду в терминале:

```bash
go get -u repository.basistech.ru/BASIS/decort-golang-sdk
```

## Список API

Библиотека поддерживает две группы API платформы:

- `cloudapi` - пользовательская группа, которая позволяет воспользоваться всем стардартным функционалом платформы;
- `cloudbroker` - административная группа, которая позволяет воспользоваться всем стандартным функционалом платформы и расширенными возможностями, включающими в себя управление пользователями, ресурсами, платформами размещения ресурсов и т.д.

### Cloudapi

`Cloudapi` позволяет выполнять запросы к группе пользовательских конечных точек
Данная группа ручек позволяет выполнять следующие операции в платформе:

- `Account` - управление аккаунтами - внутренними учетными записями платформы, которые являются владельцами вычислительных ресурсов;
- `BService` - управление группами виртуальных машин (computes);
- `Compute` - управление виртуальными машинами (индивидуально);
- `Disks` - управление виртуальными дисками;
- `ExtNet` - управление виртуальными сетями, отвечающими за внешний доступ;
- `FLIPgroup` - управление группами "плавающими" ip - адресами;
- `Image` - управление образами операционных систем;
- `K8CI` - получение информации о конвейере для создания кластера;
- `K8S` - управление кластерами kubernetes;
- `KVMPPC` - создание виртуальной машины Power PC (IBM);
- `KVMx86` - создание виртуальной машины x86;
- `LB` - управление балансировщиками нагрузки;
- `Locations` - получение информации о grid площадки;
- `RG` - управление ресурсными группами аккаунта;
- `Sizes` - получение информации о потребляемых ресурсах виртуальными машинами и дисками;
- `Stack` - получение информации о вычислительных узлах;
- `Tasks` - получение информации о ходе выполнения асинхронных задач (например, создание кластера);
- `VINS` - управление виртуальными изолированными сетями.

### Cloudbroker

`Cloudbroker` позволяет выполнять запросы к группе пользовательских конечных точек
Данная группа ручек позволяет выполнять следующие операции в платформе:

- `Account` - управление аккаунтами - внутренними учетными записями платформы, которые являются владельцами вычислительных ресурсов;
- `APIAccess` - управление доступом к API и его объектам;
- `Backup` - управление резервным копированием;
- `Compute` - управление виртуальными машинами (индивидуально);
- `Disks` - управление виртуальными дисками;
- `ExtNet` - управление виртуальными сетями, отвечающими за внешний доступ;
- `FLIPGroup` - управление группами с «плавающими» ip адресами;
- `Grid` - управление площадками;
- `Group` - управление группами пользователей;
- `Image` - управление образами операционных систем;
- `K8CI` - управление конвейром для создания кластера;
- `K8S` - управление кластерами kubernetes;
- `KVMPPC` - создание виртуальной машины Power PC (IBM);
- `KVMx86` - создание виртуальной машины x86;
- `LB` - управление балансировщиками нагрузки;
- `PCIDevice` - управление устройствами;
- `RG` - управление ресурсными группами аккаунта;
- `SEP` - управление storage endpoint (sep);
- `Stack` - получение информации о вычислительных узлах;
- `Tasks` - получение информации о ходе выполнения асинхронных задач (например, создание кластера);
- `User` - управление пользователями (индивидуально);
- `VGPU` - управление виртуальными графическими процессорами;
- `VINS` - управление виртуальными изолированными сетями.

## Работа с библиотекой

Алгоритм работы с библиотекой выглядит следующим образом:

1. Выполнение одного из действий:
 - настройка конфигурации клиента;
 - парсинг конфигурации из файла.
2. Создание клиента.
3. Создание структуры запроса.
4. Выполнение запроса.

### Настройка конфигурации клиента

Сначала, необходимо создать переменную конфигурации клиента. Конфигурация состоит как из обязательных, так и необязательных полей.
| Поле | Тип | Обязательный | Описание |
| --- | --- | --- | --- |
| AppID | string | Да | app_id ключа для выполнения запросов |
| AppSecret | string | Да | app_secret ключ для выполнения запроса |
| SSOURL | string | Да | URL адрес сервиса аутентификации и авторизации |
| DecortURL | string | Да | URL адрес платформы, с которой будет осуществляться взаимодействие |
| Retries | uint | Нет | Кол-во неудачных попыток выполнения запроса, по умолчанию - 5 |
| Timeout | config.Duration | Нет | Таймаут HTTP клиента, по умолчанию - без ограничений |
| SSLSkipVerify | bool | Нет | Пропуск проверки подлинности сертификата |
| Token | string | Нет | JWT токен |

#### Пример конфигурации клиента

```go
import (
    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

func main(){
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APP_ID>",
        AppSecret: "<APP_SECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    cfg.SetTimeout(5 * time.Minute)
}
```

#### Парсинг конфигурации из файла

Также возможно создать переменную конфигурации из JSON или YAML файла, используя функцию `ParseConfigJSON` (или `ParseConfigYAML`) из пакета config. 
<br>
*См. пример файлов конфигурации ниже и в директории `samples/`.*

```go
import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

func main() {
	// Парсинг конфигурации из JSON-файла
	cfg := config.ParseConfigJSON("<PATH>")
}
```

#### Пример JSON конфигурации

```json
{
    "appId": "<APP_ID>",
    "appSecret": "<APP_SECRET>",
    "ssoUrl": "https://sso.digitalenergy.online",
    "decortUrl": "https://mr4.digitalenergy.online",
    "retries": 5,
    "timeout": "5m",
    "sslSkipVerify": false
}
```

#### Пример YAML конфигурации

```yaml
appId: <APP_ID>
appSecret: <APP_SECRET>
ssoUrl: https://sso.digitalenergy.online
decortUrl: https://mr4.digitalenergy.online
retries: 5
timeout: 5m
sslSkipVerify: false
```

### Создание клиента

Создание клиента происходит с помощью функции-строителя `New` из основного пакета `decort-sdk`, для избежания проблем с именами, пакету можно присвоить алиас `decort`. Функция принимает конфигурацию, возвращает структуру `DecortClient`, с помощью которой можно взаимодействовать с платформой.

### Пример

```go
package main

import (
    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
    decort "repository.basistech.ru/BASIS/decort-golang-sdk"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    cfg.SetTimeout(5 * time.Minute)

    // Создание клиента
    client := decort.New(cfg)
}
```

### Создание структуры запроса

Структуры запросов определены в пакетах:

- `pkg/cloudapi` - для `cloudapi`
- `pkg/cloudbroker` - для `cloudbroker`

В каждом пакете находятся пакеты групп API:

- **cloudapi**:
  - `pkg/cloudapi/account` - для `Account`
  - `pkg/cloudapi/bservice` - для `Basic Service`
  - `pkg/cloudapi/compute` - для `Compute`
  - `pkg/cloudapi/disks` - для `Disks`
  - `pkg/cloudapi/extnet` - для `ExtNet`
  - `pkg/cloudapi/flipgroup` - для `FLIPGroup`
  - `pkg/cloudapi/image` - для `Image`
  - `pkg/cloudapi/k8ci` - для `K8CI`
  - `pkg/cloudapi/k8s` - для `K8S`
  - `pkg/cloudapi/kvmppc` - для `KVMPPC`
  - `pkg/cloudapi/kvmx86` - для `KVMX86`
  - `pkg/cloudapi/lb` - для `LB`
  - `pkg/cloudapi/locations` - для `Locations`
  - `pkg/cloudapi/rg` - для `RG`
  - `pkg/cloudapi/sizes` - для `Sizes`
  - `pkg/cloudapi/stack` - для `Stack`
  - `pkg/cloudapi/tasks` - для `Tasks`
  - `pkg/cloudapi/vins` - для `VINS`
- **cloudbroker**:
  - `pkg/cloudbroker/account` - для `Account`
  - `pkg/cloudbroker/apiaccess` - для `APIAccess`
  - `pkg/cloudbroker/backup` - для `Backup`
  - `pkg/cloudbroker/compute` - для `Compute`
  - `pkg/cloudbroker/disks` - для `Disks`
  - `pkg/cloudbroker/extnet` - для `ExtNet`
  - `pkg/cloudbroker/flipgroup` - для `FLIPGroup`
  - `pkg/cloudbroker/grid` - для `Grid`
  - `pkg/cloudbroker/group` - для `Group`
  - `pkg/cloudbroker/image` - для `Image`
  - `pkg/cloudbroker/k8ci` - для `K8CI`
  - `pkg/cloudbroker/k8s` - для `K8S`
  - `pkg/cloudbroker/kvmppc` - для `KVMPPC`
  - `pkg/cloudbroker/kvmx86` - для `KVMX86`
  - `pkg/cloudbroker/lb` - для `LB`
  - `pkg/cloudbroker/pcidevice` - для `PCIDevice`
  - `pkg/cloudbroker/rg` - для `RG`
  - `pkg/cloudbroker/sep` - для `SEP`
  - `pkg/cloudbroker/stack` - для `Stack`
  - `pkg/cloudbroker/tasks` - для `Tasks`
  - `pkg/cloudbroker/user` - для `User`
  - `pkg/cloudbroker/vgpu` - для `VGPU`
  - `pkg/cloudbroker/vins` - для `VINS`

Все поля структуры имеют описание, в которых содержится:

- Их назначение;
- Обязательный или нет - поле required в комментариях;
- Доп. информация (допустимые значения, значения по умолчанию).

#### Пример комментариев структуры

```go
type CreateRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"` 

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name" json:"name"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu" json:"cpu"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram" json:"ram"`

	// ID of the OS image to base this VM on;
	// Could be boot disk image or CD-ROM image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId"`

	// Size of the boot disk in GB
	// Required: false
	BootDisk uint64 `url:"bootDisk,omitempty" json:"bootDisk,omitempty"`

	// ID of SEP to create boot disk on.
	// Uses images SEP ID if not set
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool to use if SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	//	- NONE
	// Required: false
	NetType string `url:"netType,omitempty" json:"netType,omitempty"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	// when network type is not "NONE"
	// Required: false
	NetID uint64 `url:"netId,omitempty" json:"netId,omitempty"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`

	// Input data for cloud-init facility
	// Required: false
	Userdata string `url:"userdata,omitempty" json:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Start VM upon success
	// Required: false
	Start bool `url:"start,omitempty" json:"start,omitempty"`

	// Stack ID
	// Required: false
	StackID uint64 `url:"stackId,omitempty" json:"stackId,omitempty"`

	// System name
	// Required: false
	IS string `url:"IS,omitempty" json:"IS,omitempty"`

	// Compute purpose
	// Required: false
	IPAType string `url:"ipaType,omitempty" json:"ipaType,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}
```

#### Пример создания запроса для развертывания виртуальной машины:

```go
package main

import (
    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
    decort "repository.basistech.ru/BASIS/decort-golang-sdk"
    "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/kvmx86"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    client := decort.New(cfg)

    // Создание структуры запроса
    // CreateRequest - реквест на создание виртуальной машины
    req := kvmx86.CreateRequest{
        RGID:    123,
        Name:    "compute",
        CPU:     4,
        RAM:     4096,
        ImageID: 321,
    }
}
```

### Выполнение запроса

Чтобы выполнить запрос, необходимо:

1. Вызвать у клиента метод, отвечаеющий за определение группы API для взаимодействия, это может быть `.CloudAPI()`, либо `.CloudBroker()`. Данные методы возвращаеют соответствующие структуры, с помощью которых можно совершать запросы.
2. Вызвать у возвращенной структуры метод, определяющий группу ручек для взаимодействия.

   Доступные методы для `.CloudAPI()`:

   - `.Account()` - для работы с `Account`
   - `.BService()` - для работы с `BService`
   - `.Compute()` - для работы с `Compute`
   - `.Disks()` - для работы с `Disks`
   - `.ExtNet()` - для работы с `ExtNet`
   - `.FLIPgroup()` - для работы с `FLIPGroup`
   - `.Image()` - для работы с `Image`
   - `.K8CI()` - для работы с `K8CI`
   - `.K8S()` - для работы с `K8S`
   - `.KVMPPC()` - для работы с `KVMPPC`
   - `.KVMx86()` - для работы с `KVMX86`
   - `.LB()` - для работы с `LB`
   - `.Locations()` - для работы с `Locations`
   - `.RG()` - для работы с `RG`
   - `.Sizes()` - для работы с `Sizes`
   - `.Stack()` - для работы с `Stack`
   - `.Tasks()` - для работы с `Tasks`
   - `.VINS()` - для работы с `VINS`

   Доступные методы для `.CloudBroker()`:

   - `.Account()` - для работы с `Account`
   - `.APIAccess()` - для работы с `APIAccess`
   - `.Backup()` - для работы с `Backup`
   - `.Compute()` - для работы с `Compute`
   - `.Disks()` - для работы с `Disks`
   - `.ExtNet()` - для работы с `ExtNet`
   - `.FLIPGroup()` - для работы с `FLIPGroup`
   - `.Grid()` - для работы с `Grid`
   - `.Group()` - для работы с `Group`
   - `.Image()` - для работы с `Image`
   - `.K8CI()` - для работы с `K8CI`
   - `.K8S()` - для работы с `K8S`
   - `.KVMPPC()` - для работы с `KVMPPC`
   - `.KVMx86()` - для работы с `KVMX86`
   - `.LB()` - для работы с `LB`
   - `.PCIDevice()` - для работы с `PCIDevice`
   - `.RG()` - для работы с `RG`
   - `.SEP()` - для работы с `SEP`
   - `.Stack()` - для работы с `Stack`
   - `.Tasks()` - для работы с `Tasks`
   - `.User()` - для работы с `User`
   - `.VGPU()` - для работы с `VGPU`
   - `.VINS()` - для работы с `VINS`

3. Вызвать метод, отвечающий за выполнение запроса и передать в него:

- контекст;
- структуру запроса.
  У каждой группы ручек API имеются свои доступные методы, которые определяются платформой.

4. Обработать результат и ошибки.

Т.к. все вызовы методов идут последовательно, можно их объеденить в конвейер:
Общий вид конвейера будет выглядеть так:

```go
  client.<API>.<группа>.<метод>
```

#### Пример выполнения запроса

```go
package main

import (
    "log"
    "fmt"

    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
    decort "repository.basistech.ru/BASIS/decort-golang-sdk"
    "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/kvmx86"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    client := decort.New(cfg)

    // Создание структуры запроса
    // CreateRequest - реквест на создание виртуальной машины
    req := kvmx86.CreateRequest{
        RGID:    123,
        Name:    "compute",
        CPU:     4,
        RAM:     4096,
        ImageID: 321,
    }

    //Выполнение запроса с помощью конвейера
    res, err := client.CloudAPI().KVMX86().Create(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}
```

Для запросов Get и List реализованы запросы GetRaw и ListRaw, которые возвращают ответ не в виде соответствующей структуры, а в виде массива байт (JSON).
Выполнение таких запросов происходит аналогично.

#### Пример выполнения GetRaw и ListRaw запросов

```go
package main

import (
    "log"
    "fmt"

    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
    decort "repository.basistech.ru/BASIS/decort-golang-sdk"
    "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/account"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    client := decort.New(cfg)

    // 1. Создание структуры запроса GetRequest на создание аккаунта и выполнение GetRaw запроса с помощью конвейера
    req1 := account.GetRequest{
        AccountID:    123,
    }
    res1, err := client.CloudAPI().Account().GetRaw(context.Background(), req1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(res1))

    // 2. Создание структуры запроса ListRequest на получение аккаунтов и выполнение ListRaw запроса с помощью конвейера
    req2 := account.ListRequest{}
    res2, err := client.CloudAPI().Account().ListRaw(context.Background(), req2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(res2))
}
```

### Фильтрация

Для каждого `ListRequest` в SDK есть группа функций для фильтрации ответа платформы. Для того чтобы произвести фильтрацию по заданным полям, достаточно описать анонимную функцию (предикат) в `.FilterFunc()`, например:

```go
// Создание запроса account/list
req := account.ListRequest{}

resp, err := client.CloudAPI().Account().List(context.Background(), req)
if err != nil {
	log.Fatal(err)
}

// Тип filtered - ListAccount, тот же, что и у ответа платформы.
filtered := resp.
	FilterFunc(func(ia account.ItemAccount) bool {
	// ItemAccount открывает доступ к полям модели ответа платформы.
		for _, itemAcl := range ia.ACL {
			if itemAcl.UgroupID == "<UserGroupID>" { // Фильтр по ACL/UgroupID
				return true
			}
		}
		return false
	})
```

Для удобства пользователей, особенно важные фильтры вынесены в отдельные функции, что облегчает фильтрацию сразу по нескольким полям:

```go
// Создание запроса account/list
req := account.ListRequest{}

resp, err := client.CloudAPI().Account().List(context.Background(), req)
if err != nil {
	log.Fatal(err)
}

// Несколько фильтров объединены в конвейер
filtered := resp.
	FilterByName("<NAME>").
	FilterByStatus("<STATUS>")
	// ....
```

Также у `compute`, `disks`, и `lb` присутствуют специфические функции фильтрации, отправляющие дополнительные запросы. В качестве параметров принимают:
- context.Context - контекст для доп. запроса
- id (или другое поле, по которому производится фильтрация)
- interfaces.Caller - DECORT-клиент для запроса

Так как эти функции возвращают не только результирующий слайс, но и возможную ошибку - конвейер придется прервать для обработки ошибки.


#### Использование на примере `compute.FilterByK8SID`:

```go
func main() {
    // Чтение конфигурации из файла
    cfg := config.ParseConfigJSON("<PATH>")

    // Создание клиента
    client := decort.New(cfg)

    // Создание структуры запроса compute/list
    req := compute.ListRequest{}

    // Запрос
    resp, err := client.CloudAPI().Compute().List(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    // Фильтрация по id кластера.
    // Первый аргумент - контекст
    // Второй - ID кластера
    // Третий - DECORT-клиент
    filtered, err := resp.FilterByK8SID(context.Background(), <k8sID>, client)
    if err != nil {
        log.Fatal(err) // Возможная ошибка запроса
    }

    // Доп. фильтрация и сортировка результата - worker ноды кластера
    workers := filtered.FilterByK8SWorkers().SortByCreatedTime(false)

    // Доп. фильтрация и сортировка результата - master ноды кластера
    masters := filtered.FilterByK8SMasters().SortByCreatedTime(true)

    // ....
}
```

### Сортировка

Функции сортировки так же могут быть объединены в конвейер:

```go
// Создание запроса compute/list
req := compute.ListRequest{}

resp, err := client.CloudAPI().Compute().List(context.Background(), req)
if err != nil {
	log.Fatal(err)
}

// Функции сортировки имеют параметр inverse (bool):
// При значении false -> сортировка по возрастанию
// При true -> сортировка по убыванию
sorted := resp.
	SortByCPU(false).
	SortByRAM(false).
	SortByCreatedTime(true)
	// ....
```

### Сериализация

Результат преобразований легко сериализовать в JSON при помощи функции `.Serialize()`. Она принимает в себя 2 необязательных аргумента: префикс (prefix) и отступ (Indent).

В случае если функция вызывается без аргументов, то маршализация пройдет успешно, но без отступов и префиксов.

```go
// Сериализация данных с префиксом "" и отступом "\t"
serialized, err := filtered.Serialize("", "\t")
if err != nil {
	log.Fatal(err)
}

// Запись сериализованных данных в файл
err = serialized.WriteToFile("<PATH>")
if err != nil {
	log.Fatal(err)
}
```

#### Комплексный пример

```go
package main

import (
	"context"
	"log"

	decort "repository.basistech.ru/BASIS/decort-golang-sdk"
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/compute"
)

func main() {
	cfg := config.Config{
		AppID:         "<APP_ID>",
		AppSecret:     "<APP_SECRET>",
		SSOURL:        "<SSO_URL>",
		DecortURL:     "<DECORT_URL>",
		Retries:       5,
	}

	cfg.SetTimeout(5 * time.Minute)
	
	// Создание клиента
	client := decort.New(cfg)

	// Создание запроса compute/list
	req := compute.ListRequest{}

	resp, err := client.CloudBroker().Compute().List(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	// Цепь преобразований.
	filtered, _ := resp.
        FilterFunc(func(ic compute.ItemCompute) bool {
            return ic.GUID == 123 // Фильтр по GUID
        }).
        FilterByName("<NAME>"). // Фильтр по имени
        SortByCPU(false). // Сортировка по кол-ву ядер CPU по возрастанию
        SortByCreatedTime(true). // Сортировка по времени создания по убыванию
        Serialize("", "\t") // Сериализация с префиксом "" и отступом "\t"
        // Serialize - терминальная функция, возвращает Serialized (обертку над []byte)

	// Запись данных в файл
    err = filtered.WriteToFile("<PATH>")
	if err != nil {
		log.Fatal(err)
	}
}

```

## Работа с legacy клиентом

Работа с legacy клиентом применяется для пользователей, которые не используют для авторизации decs3o.

### Настройка конфигурации legacy клиента

Сначала, необходимо создать переменную конфигурации клиента. Конфигурация состоит как из обязательных, так и необязательных полей.

| Поле          | Тип    | Обязательный | Описание                                                           |
| ------------- | ------ | ------------ | ------------------------------------------------------------------ |
| Username      | string | Да           | username legacy пользователя                                       |
| Password      | string | Да           | пароль legacy пользователя                                         |
| DecortURL     | string | Да           | URL адрес платформы, с которой будет осуществляться взаимодействие |
| Retries       | uint   | Нет          | Кол-во неудачных попыток выполнения запроса, по умолчанию - 5      |
| Timeout       | config.Duration | Нет | Таймаут HTTP клиента, по умолчанию - без ограничений               |
| SSLSkipVerify | bool   | Нет          | Пропуск проверки подлинности сертификата     |
| Token         | string | Нет          | JWT токен                                                          |

#### Пример конфигурации legacy клиента

```go
import (
    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

func main(){
    // Настройка конфигурации
    legacyCfg := config.LegacyConfig{
        Username:  "<USERNAME>",
        Password:  "<PASSWORD>",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    legacyCfg.SetTimeout(5 * time.Minute)
}
```

#### Парсинг legacy конфигурации из файла

Также возможно создать переменную конфигурации из JSON или YAML файла, используя функцию `ParseLegacyConfigJSON` (или `ParseLegacyConfigYAML`) из пакета config.
<br>
*См. пример файлов конфигурации ниже и в директории `samples/`.*

```go
import (
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

func main() {
	// Парсинг конфигурации из YAML-файла
	legacyCfg := config.ParseLegacyConfigYAML("<PATH>")
}
```

#### Пример legacy JSON конфигурации

```json
{
    "username": "<USERNAME>",
    "password": "<PASSWORD>",
    "decortUrl": "https://mr4.digitalenergy.online",
    "retries": 5,
    "timeout": "5m",
    "sslSkipVerify": true
}
```

#### Пример legacy YAML конфигурации
```yaml
username: <USERNAME>
password: <PASSWORD>
decortUrl: https://mr4.digitalenergy.online
retries: 5
timeout: 5m
sslSkipVerify: true
```
### Создание legacy клиента

Создание клиента происходит с помощью функции-строителя `NewLegacy` из основного пакета `decort-sdk`, для избежания проблем с именами, пакету можно присвоить алиас `decort`. Функция принимает конфигурацию, возвращает структуру `DecortClient`, с помощью которой можно взаимодействовать с платформой.

#### Пример создания legacy клиента

```go
package main

import (
    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
    decort "repository.basistech.ru/BASIS/decort-golang-sdk"
)

func main() {
    // Настройка конфигурации
    legacyCfg := config.LegacyConfig{
        Username:  "<USERNAME>",
        Password:  "<PASSWORD>",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    legacyCfg.SetTimeout(5 * time.Minute)

    // Создание клиента
    legacyClient := decort.NewLegacy(cfg)
}
```

#### Пример выполнения запроса

```go
package main

import (
	"fmt"

    "repository.basistech.ru/BASIS/decort-golang-sdk/config"
    decort "repository.basistech.ru/BASIS/decort-golang-sdk"
)

func main() {
    // Настройка конфигурации
    legacyCfg := config.LegacyConfig{
        Username:  "<USERNAME>",
        Password:  "<PASSWORD>",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    legacyClient := decort.NewLegacy(cfg)
    
    // Создание структуры запроса
    // CreateRequest - реквест на создание виртуальной машины
    req := kvmx86.CreateRequest{
        RGID:    123,
        Name:    "compute",
        CPU:     4,
        RAM:     4096,
        ImageID: 321,
    }

    // Выполнение запроса
    res, err := client.CloudAPI().KVMX86().Create(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}
