const TelegramHint = () => {
  return (
    <div className="my-4 max-w-2xl mx-auto bg-white shadow-lg rounded-2xl p-8 space-y-6">
      <h2 className="text-2xl font-semibold text-gray-800">
        Подключение Telegram-бота
      </h2>

      <ol className="space-y-4 list-decimal list-inside text-gray-700">
        <li>
          Создайте бота через{" "}
          <a
            href="https://t.me/BotFather"
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-600 hover:underline"
          >
            @BotFather
          </a>.
        </li>

        <li>
          После создания откройте настройки бота и{" "}
          <span className="font-medium">скопируйте Bot Token</span>.
        </li>

        <li>
          Создайте <span className="font-medium">Telegram-канал</span> для
          получения уведомлений.
        </li>

        <li>
          Добавьте созданного бота в канал и{" "}
          <span className="font-medium">назначьте его администратором</span>.
        </li>

        <li>
          Отправьте любое сообщение в созданный канал.
        </li>

        <li>
          Перешлите это сообщение боту{" "}
          <a
            href="https://t.me/FIND_MY_ID_BOT"
            target="_blank"
            rel="noopener noreferrer"
            className="text-blue-600 hover:underline"
          >
            @FIND_MY_ID_BOT
          </a>.
        </li>

        <li>
          В ответ вы получите <span className="font-medium">Chat ID</span> -
          используйте его для настройки интеграции.
        </li>
      </ol>

      <div className="bg-gray-100 rounded-lg p-4 text-sm text-gray-600">
        <span className="font-semibold">Важно:</span> Chat ID для каналов обычно
        начинается с <code className="bg-gray-200 px-1 py-0.5 rounded">-100</code>.
      </div>
    </div>
  );
}

export default TelegramHint;