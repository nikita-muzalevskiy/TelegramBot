from __future__ import print_function

import random
import telebot
from telebot import types
from telebot.async_telebot import AsyncTeleBot
import asyncio
import os.path
import logging
import grpc
import pb_pb2
import pb_pb2_grpc
import os
from dotenv import load_dotenv, dotenv_values


load_dotenv("variables.env")
bot = AsyncTeleBot(os.getenv("TOKEN"))

services = {
    "MANGA": pb_pb2_grpc.MangaStub(grpc.insecure_channel(os.getenv("HOST"))).Channel1
    # "TESTS": pb_pb2_grpc.TestsStub(grpc.insecure_channel("localhost:50050")),
}

async def CallbackToMessage(call):
    # user = call.from_user.id
    user_id, service, action, param = call.data.split("/")[0], call.data.split("/")[1], call.data.split("/")[2], call.data.split("/")[3]
    response = services[service](pb_pb2.CallbackRequest(user=user_id, action=action, param=param))

    keyboard = types.InlineKeyboardMarkup(row_width=2)
    for b in response.buttons:
        keyboard.add(types.InlineKeyboardButton(text=b.text, callback_data=b.data))
    await bot.send_message(user_id, response.text, reply_markup=keyboard)

@bot.message_handler(commands=['start'])
async def start(message):
    await bot.send_message(message.from_user.id, "👋 Привет! Я аниме-бот! Я могу тупить и лагать, да и вообще я еще не доделан... Но я постараюсь для тебя!")
@bot.message_handler(commands=['menu'])
async def menu(message):
    keyboard = types.InlineKeyboardMarkup(row_width=2)
    buttons = [types.InlineKeyboardButton(text='Манга', callback_data=f"{message.from_user.id}/MANGA/menu/"),
               # types.InlineKeyboardButton(text='Тесты', callback_data=f"{message.from_user.id}/TESTS/to-tests-section/")
               ]
    keyboard.add(*buttons)
    await bot.send_message(message.from_user.id, "Ты в меню! Выбери раздел:", reply_markup=keyboard)

# реакции на сообщения
@bot.message_handler(content_types=['text'])
async def get_text_messages(message):
    if message.text == 'Hi':
        print('Hi')

# реакции на инлйан-кнопки
@bot.callback_query_handler(func=lambda call: True) #вешаем обработчик событий на нажатие всех inline-кнопок
async def callback_inline(call: telebot.types.CallbackQuery):
    print(call.data)
    if call.data: #проверяем есть ли данные если да, далаем с ними что-то
        await CallbackToMessage(call)

asyncio.run(bot.polling())

if __name__ == "__main__":
    logging.basicConfig()
    # run()
