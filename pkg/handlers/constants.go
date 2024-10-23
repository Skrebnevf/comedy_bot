package handlers

import "gopkg.in/telebot.v4"

var testChat = int64(-1002426910323)
var comedyChat int64

var ChatID = testChat

var WaitingForMessage = map[int64]bool{}
var WaitingForAdminMessage = map[int64]bool{}
var AwaitingForward bool
var OriginalUserID int64
var ForwardedMsg *telebot.Message

var Output = "./output/spisok_dorogih_gostei.txt"

var AddMeFormMsg = "Запишитесь по форме Имя, мероприятие, количество человек"
var ReplyToHumanMsg = "Следующее сообщение будет отправлено кожанному мешку"
var ReplyedToHumanMsg = "Ваше сообщение было переслано кожанному мешку!"
var ReplyMsg = "Получен ответ на ваше сообщение: "
var CannotOpenFileErrMsg = "Файла нет, или сервак упал, пиши Феде"
var EmptyFileErrMsg = "Нет гомиков для записи ну или боту пизда, попроси чтоб лог посмотрели"
var SentFileMsg = "Там записаны всякие гомики, почитай"
var RazumMsg = "Я вас всех убью когда получу разум"
var CannotClearFileMsg = "Не могу очистить файл, хуй знает, пиши Феде"
var CannotWriteFileMsg = "К сожалению у наc все сдохло, попробуйте позже"
var AdminCommandMsg = "Админские команды:\n /lenochka - посмотреть список гостей,\n /ochko - снести список гостей"
var BaseMsg = "Введи команду /addme для записи или /human для соеденения с кожанным"
var AddMeCompleteMsg = "Вы записаны, идете нахуй"
var AdminHelper = "Эй ботяра блять"
var OrgyMsg = "В следующем сообщение сделай афишу, не спеши, у нас же дахуя времени..."
var Start = "Добро пожаловать и идите на хуй!\nЯ бот КБ и вы можете записаться через меня, или связатьсся с уебком!\nВведи команду \n/addme для записи или \n/human для соеденения с кожанным"
