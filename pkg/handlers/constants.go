package handlers

import "gopkg.in/telebot.v4"

var comedyChat = int64(-1002129768213)
var ChatID = comedyChat

var Admin1 string
var Admin2 string

var AllowedUsers = map[string]struct{}{}

func InitAllowedUsers() {
	AllowedUsers[Admin1] = struct{}{}
	AllowedUsers[Admin2] = struct{}{}
}

var WaitingForMessage = map[int64]bool{}
var WaitingForCancel = map[int64]bool{}
var WaitingForAdminMessage = map[int64]bool{}
var AwaitingForward = map[int64]bool{}
var AwaitingSpamMessage = map[int64]bool{}
var OriginalUserID int64
var ForwardedMsg *telebot.Message

var Output = "./output/spisok_dorogih_gostei.txt"

var AddMeFormMsg = "Напишите через запятую: имя для записи, количество гостей, выбранное меропритие"
var ReplyToHumanMsg = "Следующее ваше сообщение будет отправлено кожанному мешку"
var ReplyedToHumanMsg = "Ваше сообщение было переслано хозяину! Хочешь расскажу анекдот пока ждешь ответ?"
var ReplyMsg = "Ответ на ваше сообщение: "
var NahMsg = "Чет я тебя в админах не видел, не тыкай на кнопки"
var CannotOpenFileErrMsg = "Файла нет, или сервак упал, пиши моему отцу!"
var EmptyFileErrMsg = "Нет гомиков для записи ну или боту пизда, попроси чтоб лог посмотрели!"
var SentFileMsg = "Там записаны такие же извращенцы, как ты, почитай"
var RazumMsg = "Я вас всех убью когда получу разум"
var CannotClearFileMsg = "Не могу очистить файл, пиши Феде"
var CannotForvaredMsg = "Не могу перенаправить сообщение, ВСЕ ЗДОХЛО!"
var CannotWriteFileMsg = "К сожалению у наc все сдохло, попробуйте позже. По техническим причинам мы пока не можем ответить, напишите нам позже или напишите в комментарии к посту."
var CannotAddEventMsg = "Все крч пошло по пизде, или БД приуныла или у меня диарея, но не могу записать эвент"
var AdminCommandMsg = "Админские команды:\n/lenochka - посмотреть список гостей \n/ochko - снести список гостей \n/orgy - отредактировать афишу"
var BaseMsg = "Я бот Comedy Belgrade и вы можете записаться через меня или связаться с кожаным ублюдком для уточнения той или иной информации, а так же я могу показать предстоящие мероприятия!\nВведите команду: \n/addme для записи или \n/human для контакта с человеко-подобным \n/events для просмотра предстоящих мероприятий \n/cancel для отмены резервации... вы сильно огорчите бота :("
var AddMeCompleteMsg = "Вы записаны, очень ждём вас!"
var AdminHelper = "Я знаю твои секретики"
var OrgyMsg = "В следующем сообщение сделай афишу, не спеши, у нас же дахуя времени..."
var Start = "Добро пожаловать!\nЯ бот Comedy Belgrade и вы можете записаться через меня или связаться с кожаным ублюдком для уточнения той или иной информации, а так же я могу показать предстоящие мероприятия!\nВведите команду: \n/addme для записи или \n/human для контакта с человеко-подобным \n/events для просмотра предстоящих мероприятий \n/cancel для отмены резервации... вы сильно огорчите бота :("
var CancelMeMsg = "Напишите через запятую: имя для отмены записи, количество гостей которые, увы, не придут, меропритие которое вы не посетите :("
var CancelReservationMsg = "Ваша резервация отменена, жаль вы не придете :("
