# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

### 1. Описание функциональности монолитного приложения

- Система поддерживает управление температурным режимом в жилом помещении клиента
- Система поддерживает проверку текущей температуры в жилом помещении клиента
- Система поддерживает подключение нового датчика специалистом компании

### 2. Анализ архитектуры монолитного приложения

**Язык программирования:** Golang <br>
**База данных:** PostgreSQL <br>
**Архитектура:** Монолит(Все компоненты системы находятся в одном приложении) <br>
**Взаимодействие:** Синхронное <br>
**Масштабируемость:** Ограниченна, монолит можно тольк вертикально масштабировать до существенных потерях в таймингах
**Развертывание:** Требует полной остановки инстанса всего приложения
**Отказоустойчивость:** Отсутствует, при ошибке будет недоступно все приложение

### 3. Определение доменов и границы контекстов

**Домен:** Управление отопительной системой <br>
- **Поддомен:** Управление датчиками <br>
    - **Контекст:** Подключение/отключение датчика <br>
    - **Контекст:** Информирование о состоянии датчика <br>
- **Поддомен:** Управление температурой <br>
    - **Контекст:** Изменение температуры <br>
    - **Контекст:** Информирование о текущей температуре <br>

### **4. Проблемы монолитного решения**

- *Дорогая стоимость обновлений* - релиз новой фичи будет долго развертываться(из-за перезапуска всего приложения) и намного ответственней придется подходить к сбору информации о ее безопасности для других частей приложения
- *Плохая отказоустойчивость* - при неудачном релизе весь функционал будет страдать или будет полностью нерабочим
- *Тесное взаимодействие команд* - команды не должны быть посвящены в нюансы других чтобы улучшать свое направление
- *Сложно мастабировать отдельные направления* - если сервис отопления больше всего востребован, то сложно будет масштабировать только его, не тратя ресурсы на другие части, которые будут в холостую разходовать денежные средства инфраструктуры

### 5. Визуализация контекста системы — диаграмма С4

[Контекст C4](https://www.plantuml.com/plantuml/uml/TP0_JyD03CNtV0gJAKWqMJenLOegTG2f0Z57piLDJhe_qTcXyEtnfON2mEXtl-VtV3iMJB9u1-sssplo58IBEnkCMmQ2IMJcnwP9UAv7AzFIBqp9n23ArIRwPdOOid_3CZZRD-rsSypIs_DJuv65aafTgmwqWL8zAriKHqeSGvMNFK3v-gcS1xEmH4-fL4bV610cGh5X15w9yhwue1LxRNs8IcLx7c2TuFAUDvof1niq_uI6u1GJv3J0C82F8JehWFFyixBRVqXrZLFtuM-mAlde51s3JFJTIzt52xdh7RTI_-EP_itT4XCPik_4CA3WMdDSHnEzrudLlACmwFTz0G00)


# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

[C4-container](https://www.plantuml.com/plantuml/uml/jLRVRzj637xtNy4Q0ve2S_omxML0W6PEkdP86YzEsyU0bcXBhQHJZrIoRDt_FXxrezO4sqNYvkKkl8_yFlB8irwoe9SgpqRFqYBAgfWW4Idvr-dKu_tnEfMaMbPCFdA5K277aSkdPOP5y9cmn5cwdCv-dio3wTtbnJH75l9gkfsf1wO5-UCon1yzYscb_yBP-VitiwktYsyWagh4sF_qbJ3LqybA4HCkVIh4BLzCT-6CEvPNLszlJZ-SBwukpmzAP3Qy9tOvJNuPjMNh4KwZp5MnfTsTTsbFF-8TRZih7SlQ4t_EjY4ynIbtXqFgt2tJZ9gIZ-RasHL78TqVORnuOA4Seef5rVknclI40O5pz09iv-CNwbX3RtzpLH6ZVpWoG6qCWGp-sgatB1ed-6Sq0dWEAszrZO6haZnWMQgvwuMZMjzfMMgOIzi76Q-oIiGK_H0JHnvEbAJrXtkzCy2r5cQDGe25aqIFcBnUEd7mgc_fP5o7rnHCIvv6thMmdWoi957UKrsSJ_z6ILq12qEQe3_qai9war3B0ofdvNpEW3qlr5ploAUhb1X4yQ7QiFPOI4zQk7SLjgaY8GrHh9NdTRtR9o5HianfF7skY8L11JMNbxYSe0cuXpPBrubev8kmFfNIWkrBDDmRQ-3Vm_fKHWksX_0kZSbftFTXVIfXX3bvt4UPkSWQWaDbc_q-PXRdYQs5IkyYOYPMpeeRetSkPuWyeL0CouUklmPQcewFKN29tE_siwNr-zboqF5dDTpqpUktP_7dnMRmhozIDvmduuN9xiGyPoZjeS7C1AYO8ViCeuG60ipQat5DSTehjJV3W9o9EKGGzuaA7b0sas0vjyFrf6mjCy2bgwGTCs1bTdv8RZDXr3PYXoEsm5jfRJ2CrkRiSBGMU2kjpPhHshGTZjO2RwLjvysOsu4x77dNtCE1R8QjLGRsSx-fxlRyBsaxS66WZIwS1VG6TMklWP1ZKQqmaieRjyMs1qay5fodOk38vGgCUuDYK2jJNuWrJIV3Ed8WmgpVJu8rK2V1kdCWmQpVJu8rKoV1EdKWmQo7a_3lIF_wNo7WARD6vT4rPSs7sUOtqJiczGnd6zRcQ-KIFn73wVmDH9Xb3E92btzaLtonJzoXOzE3Z_ArjdcKfztseppjaXxbsRu6tyWxwdug_tybTQsxtyjhmdZw8TmhVmaFoOi-gFrPw-5kPlC6rzM_sUm8sPJRbXsmfhQsx80rXRHb1soZQkrs2reBPctTgzOmmMx0T1Etmz5MVP3dNdGltRteNjQxqI_LMENPVm00)

**Диаграмма компонентов (Components)**

[C4-components](http://www.plantuml.com/plantuml/uml/j9NDJXn7483l-nHb5Cb4odeFIIwMB6MDiOm4Xh247AsYftQtxPxkILSFQFFpC7cZF5AgksVQimnB0W8kZBf-wgkVhVg90_hGrcRoGbjbsefW7KB3hsSpZpUlLZgiswkMoIjd0zdmIhbwrXYqqMRAeJBwQdRmm_GiFlroSZohaGDvUVfq81QeBVbNpJFvhnjdnIRxdxoeQ2dXuDtXvT71wSV5_mWKsWOh__tNW5gaqwLeJBdnEX3tyIgwZZ8kKTwVVhoO_tguE3qvVD90wGrlY5rDqnydfMp3W3DbN5jjbxMaFVkCrxXBrZWEAq_yczcbuQdIN8JZqBMxqeQwaa_EoBEp-p7Txs1liU50DQYMWz3xFNaI2GC2r-W3S9BlVIk6MVNJMzVQ2lrcFodan-WegNzmOWsBpWZ-c4m0NiBIfxfMm6r37h1fvBdCsdxccpUDk3b9tn7ZlMb3I4JVLCJAmni9qjl3ZVGCS2M5MM4WGEl2MaHCNffE7ArobzxiPVUIGcAfjVAkKuCfpDkm7i9yJGz5KhABQixhtp5ePs6HpLE287zbhFTPh-oHZUMG3pYooR5tnkJI5iNNi10wkuAJomCu9mu3nPBncxsvCUw6GLA3s2y83enRWRRobLvHAU4DajmsXvH5SuWDFizV4VmPQ78S-KRRimXTKicJCo8BxWjPbkWNwUDfa3QkZJGrMbnH0Np9mCGS4yzHHp2TkEjdR5kd7-5Ad-y6Ay82SqP-wNoTAxHihKerTqjel5F47Atk1yf3LsZYr0qX_ldx8HHe3E3MtAAjm0_duVvMfRroZ1lojol3i0ZEusf7QTuIsc6Jm4Gd85WICP175KGj_uPZ_FCOQBiqm3d02AfRXprMtcPeu1q6V5gginnfyVCnL1Zm2fcacNdqbTH6V6iqd1SBI19ckCdA9ZZ1ByGmFpi2XSOm11VnFWlrd-a8YVtOl6_Je-tpV3tCl4VU7iAMQQ1JkEvJweBt7TsbDdJn8DsjmTpMk-rcUroAybyJkEiQV22CCt3tGLXdORa9dV9zru0iE_-E0gLLq1c8wI8984kUPJ4hPwtudacXkiyapaIvNYFGplS8itj_1iYizzza5sIefk0t8xOYUOxRfaXVzsGG-aWx6Bin614-zeYWCSF3uPOTIBail0CapVfj37xuWhnOKxSYSNXbox8icxAinn6GSdMDjffNs8GXK_S01rc-kqOxiFfVB5dbv7bMKVoVQzR5PBYYS4EKBnmsZT6gNC71nCNLkAEGbsYqfERkA6OHFHWzL_GwsrEUkCOnQrdTEcoap_rlDphWXfHUQjN1ygxP2mB8GQinQH4zabGvOmI6-t9nxuzJ8LqR_gtU37hJ1e8e_5UuzTY_K2bKlVGKgUqP0Wr88KR4L3sq_xtUnpvOtdei3rt2SgozC_FP7R-tTeRjlZkRFXtzVcXfuMN-fefL89cAZ2-TTNs6EAXhoTKuGWrsROp1SAsXGc85QduDzTVaW-KC5g5nC8-9_jlxwlXr_Rng-xSAB-7KSQE0XUEt2LXU-oQ1aFLj4KXMwmQ1V3VT7X2Ydiq1YUvhQq3iwZS6vAnyMq3C0fk24gbKL1p935NFfM1uC-0DKs9f0pjKi-TYCMocZZ2ffTRl21G9SN3QaRRXVmdFPR_XFuTJHyF7ml66Z-Mr6puXwpDyajLg-EIxgU4JeXx39z6T6JwnQpTySbPk-CGiOFXCLSJm_SaCLS-bODZmlM52h8PlXshsN2o6ZTGH9hNKy1s18Y4EJZlIDlplwyMowJVyN-7KqV2nSBpXOtdjXa_8-WoVPBKQFlbkQlY4gCVmINHVXa_iwWsVd9KRFZ4B63vJ5J5yVp93rNCf63PyRvWGg-7RePez5ujX8_M4IIqnV4UWI8Y3quWS_gJV76J2-AeVpnvjdyi7d7yL1k9ETjC5OKIkY5N1fjwImQTLirdcEh5hciqPmMDtu3e0ePG0y1vm8eCteFni5MU-xBiGq_8MN0t1k0_i7mwZi7tUnPiDTTR-l5HDfSbSQOGXGbcJG589bQsrvudq_kTv1vPIDg0RlAtddmEdSumfJUVPmEWqBW-tu3HMbWwfdgPya__cq1A7VG8RQyjnjm008gt-C7jQBJC073t3mCnTIh330Ick4vHQaaDHk6OEHFtNrUBA9hDvisydCNbqZxmu_ioAqlY0Mwum4fjbyR4NPV4QMQui4fjbySg9iZ2pScMHs2pBHm6LALd5YgVDZQSLkqiAA1usaaqmGdzcpPz_jyUnF2nAkl0Xh4vAy_XXON2i_SzEAlSxr4Pkp8AskzdhssRTOZ7V9z8BmzAEJK7GD1637QaGE7ez13HGH61tAWAENWy1pHGHsAcAWAFb488e1mK__AGX_JN1juFnVS2OFzQaewEfN3ukPukFFq9UKJouLWZwqpvjksHT4fb1fkOdEQ9Eeyz3WkuXG3XOOYCYKz-q32LrJ4eI6HCi-cpOwU399qRqAGdb3tfKtKlfqZH-ChGloUJB8uNvKh8qZPyCNKqoUKMYTbzAbgRnar6raLYsDlpd-0x-j7nTx3Sl_m00)

**Диаграмма кода (Code)**

[C4-code](http://www.plantuml.com/plantuml/uml/tLTDZzis4BthL_1MBzerUXOC8rjl2YoG88lO1NBbKhCo6ubKoT4YxiT_xqWUHQ8yAXoWXs2D12lnlPbvyq5HUXDH1sphAekVhMjqqBKobOxn3Z7OfnRXmNYtjt-2-ZdBVb0xNpVUWSDStHMwGGZTKgrTeVOGNgo1B6kq-QnBKA5Tnvgy7tqRONq4ZTQLQrFvjb2RpQ--AY1ij-glB5Fdc6hNOO_1lrY2Tjuzsv9WHRyVbPbY7w6q4SDfTPCJ_8_eOMInZrkbpv9ZhjxhPhF7G89-Icqugst6OKFlu-lwgc5UhjXyS9gBVcxOJGauU2YWAzD3CNYuaIs7lAwevfebf5LLtSywMd_M1iGOpEeLmXSmBG9tSsNEVtDrtylbPTBALnyXjXMEvXYqYvM-ma4lX6z7-qQ726UdgpufkcwyYp2KIGhD9L93zpzPxMlR1xZ8y5Pz6T5TPViP_ZHV4jaqB0q4Pa-NijTLlaDxqal5GyAIqXtI1O79PMFcAohQjspDcdA-OcUoE2xHT5-Y7pRja6N1oScvilJHmNFmzSp5_xB3ldUfF7RVfz3k_zKZv0keX0dgqqobJdVi_9mP4MtGUdUvtX__S7b0VdZw7GnkHoBQ6YBgkidLWwEBVlipV3xixZduIYWUmNUzWWNmyQiC8aX2f6xtdHQgC3-l4mIzyLMk7ldgS6fWMerJvRLGFjI1LFuMgak86bVCoyUaJlN5aszrpI0dX6bkiJK68aLxyhu2xKQafcLTWg1TwXnZ0Vve2PGcHfhj4UPih7S7MnDytmRTtSrdSOyQsqZuv2wfYetxSqfKbpGf7ettvxmM1kSm94TucXuJTjCq5foydUP2d5asp9gcGAGa9Q672JvjeNULTI1qaAnm3-24FZrfH0YdT_02LGSZb116EDhIwOeo0rVYKI102BwRqNSsukOix6stj9truaElRQaLlqnK42SDmvnz8aLa7NNiWwSjbml8EgVrCxuGvh60Q89jvATS9CN1uYbNlRSzttxBqstdYwNnkHz3Bp26W0HB2lXie7-Pz3bqUxjT-CZ9bGcW4MAMoGGobRy2SbLhfqko5E7Dt-krz6gTArit5TJWyByiXtDOjii4noHKUD78LHj9OkgEsRmuvsR2grMI3Li8oOXcIPXhCqaFwBmxTmn4ahTaC8q-wQhCJay7eWV0DhYOoOLQ5YFIPIaoTr582X0rVHgIM-zGMw8cwqHCTls4r-sKB7i3hcZhwby0)

# Задание 3. Разработка ER-диаграммы

[ER-диаграмма](http://www.plantuml.com/plantuml/uml/nLV1Rjj64BtpAmQ-D44XVu0G654aMHAiMeKakyY9s90ZYWYvI-mkxGgnWBRN7ee2xR7_a4kG0aN-WVwZ3aTKo3NA1ghY1hxipjlblFTsThnwRgpGDayJ3sMUmibSgNIvoV04X05D4z_I3Dvu0DF9zUHovCzdiszezc8ug4ONasL_FBmiOxtbk0gEUzFfTxsND9gTdq_wGngCUlFUg1eCglNJ_kpofRVrABkDxGPEhWngpftJ09uzw-Ls3GlKDt60PsVCusk8movSNKq6X7zxSNP6iKw70Kn5d7HXONKi8uvamfXRfKD_BSowYN02AL9qWh7nHM3Z6ug-K2f18JaSQ1GMGr_OBYpZ5CcqD6Capq92NA2fPwpIdPTh6X3VgGg4ZPNy9qqLaLhLZgJlGEUEO4gIk0f466eqfeuTGv_BuMCf74oVHVyR_OV-rnNNXNqbSdoLAAMxC94M8zJ7dWbjf8G390NLuUmYViWXebaj-LJbrPmRW2inI5GULXzkMD5oUY6kH9vOtsAQeHOsruGFLFugGSRJMBPZuliMx1Ypfd6qjWshad9U6yNm5xFfbIxJIpIcheqc7gX4wSzq5KQ4Ddo8Q5hRK8BcVxT1v5RvYGfU-oW57NNeTYr6Gap4ngUwoOdXvxf1rt68gc76JJcltM3ueHlE8l-nhT4OA1tIpBNc-8Q-h6_OE3PMwGqxO2dihtU10_k7stpQ3MYvsvLK3Div3oNgQDDNqkHfLhPlfemSzHCLFHBXJDsYzdE3uO5ykQbS6H4v5dFdgJq-eFNmpa_IlHN6LzAjaxO6upIMGrIDjCgpXfoedE_5-0_5CEh_Tqc7bpa_kxQgqr5lTa-MBeWW8FreF9xVtJ5MtbN20YMjYANnlAzWrm_0AfYNeCOLQfIq5vxqvxF5ehCOpg_fAGWTa0hEByhDnchQRirJpoqgs9wUxdCaAc0HiVHgrMrB0BowCxTzfkv4hKj8pn8JJD7g3Kb9-DLZrd5cNAPxGzm6q8MrCCr_ndkaPNDz5QkjLI2ehsobtTstYxbPSoEYEX_fuDFWLtpvnFTOwxRJkRjhlEzvAmyzhrcOLIKqSpSFtVEaikMHPw1Mu6hpOD_GuVv7-lkv-AjuMxox_ud6lq3neNZB-2dXnP_5Ut2Vbho_-AFu7LPAOnn9U8qRU7B_0ttZVV61llOEYb-BtvvwA4CeMRXSQiR7Cm4uYWlV8m-gjtmh1yO_ScZ-xZau-tCQvMdoDm00)

# Задание 4. Создание и документирование API

### 1. Тип API

Использую REST API, так как требуется немедленный ответ. К примеру изменение информации о датчике должно сразу отобразиться у пользователя.

### 2. Документация API

[Swagger](https://editor.swagger.io)

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


