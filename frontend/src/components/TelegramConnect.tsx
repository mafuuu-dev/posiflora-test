import {useEffect, useState} from "react";
import IntegrationStatus from "@api/shops/telegram/status";
import IntegrationConnect from "@api/shops/telegram/connect";

type Props = {
  shopId: number;
};

const TelegramConnect = ({shopId}: Props) => {
  const [isLoading, setIsLoading] = useState<boolean>(true);

  const [chatId, setChatId] = useState<string>("");
  const [botToken, setBotToken] = useState<string>("");
  const [isEnabled, setIsEnabled] = useState<boolean>(false);

  const [integrationChatID, setIntegrationChatID] = useState<string>("");
  const [integrationIsEnabled, setIntegrationEnabled] = useState<boolean>(false);

  const [statsSentCount, setStatsSentCount] = useState<number>(0);
  const [statsFailedCount, setStatsFailedCount] = useState<number>(0);
  const [statsLastSentAt, setStatsLastSentAt] = useState<string>("");

  const integrationStatus = async () => {
    setIsLoading(true);

    IntegrationStatus(shopId)
      .then(response => {
        setIntegrationChatID(response.data.data.integration.chat_id);
        setIntegrationEnabled(response.data.data.integration.is_enabled);

        setStatsSentCount(response.data.data.stats.sent_count);
        setStatsFailedCount(response.data.data.stats.failed_count);
        setStatsLastSentAt(response.data.data.stats.last_sent_at);
      })
      .finally(() => setIsLoading(false));
  };

  const integrationChange = async () => {
    setIsLoading(true);

    IntegrationConnect(shopId, {
      bot_token: botToken,
      chat_id: chatId,
      is_enabled: isEnabled
    }).then(() => {
      setChatId("");
      setBotToken("");
      setIsEnabled(false);
    }).finally(() => {
      void integrationStatus();
    });
  };

  useEffect(() => {
    void integrationStatus();
  }, []);

  return !isLoading ? (
    <div className={`flex flex-1 min-h-screen min-w-screen justify-center items-center`}>
      <div className={`w-75 bg-white rounded-2xl shadow-lg p-6 space-y-6`}>

        <div className={`space-y-1`}>
          <h1 className={`text-xl font-semibold text-gray-800 mb-4`}>
            Shop #{shopId}
          </h1>

          <hr className={`my-4 text-gray-300`}/>

          <div className={`flex flex-row justify-between items-center gap-2 text-sm mb-4`}>
            <span className={`text-gray-500`}>Status:</span>
            <span className={`font-medium ${integrationIsEnabled ? "text-green-600" : "text-red-600"}`}>
              {integrationIsEnabled ? "Enabled" : "Disabled"}
            </span>
          </div>

          <div className={`flex flex-row justify-between items-center gap-2 text-sm`}>
            <span className={`text-gray-500`}>Chat ID:</span>
            <span>{integrationChatID}</span>
          </div>

          <div className={`flex flex-row justify-between items-center gap-2 text-sm`}>
            <span className={`text-gray-500`}>Last sent at:</span>
            <span>{statsLastSentAt}</span>
          </div>

          <div className={`flex flex-row justify-between items-center gap-2 text-sm`}>
            <span className={`text-gray-500`}>Sent count:</span>
            <span>{statsSentCount}</span>
          </div>

          <div className={`flex flex-row justify-between items-center gap-2 text-sm`}>
            <span className={`text-gray-500`}>Failed count:</span>
            <span>{statsFailedCount}</span>
          </div>
        </div>

        <hr className={`my-4 text-gray-300`}/>

        <div className={`space-y-4`}>
          <div className={`flex flex-row gap-2`}>
            <label className={`text-sm text-gray-600`}>Is Enabled</label>

            <label className={`items-center gap-3 cursor-pointer select-none`}>
              <input
                type="checkbox"
                checked={isEnabled}
                onChange={(e) => setIsEnabled(e.target.checked)}
                className={`sr-only peer`}
              />

              <div
                className={`
                  w-5 h-5 rounded-md border border-gray-300 flex items-center justify-center transition
                  peer-checked:bg-green-600 peer-checked:border-green-600
                `}
              >
                <svg
                  className={`w-3.5 h-3.5 text-white opacity-0 peer-checked:opacity-100 transition`}
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="3"
                >
                  <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7"/>
                </svg>
              </div>
            </label>
          </div>

          <div className={`flex flex-col gap-1`}>
            <label className={`text-sm text-gray-600`}>Bot Token</label>
            <input
              type="text"
              className={`
                border border-gray-300 rounded-lg px-3 py-2 text-sm 
                focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition
              `}
              placeholder="Enter Telegram Bot Token"
              value={botToken}
              onChange={(e) => setBotToken(e.target.value)}
            />
          </div>

          <div className={`flex flex-col gap-1`}>
            <label className={`text-sm text-gray-600`}>Chat ID</label>
            <input
              type="text"
              className={`
                border border-gray-300 rounded-lg px-3 py-2 text-sm 
                focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition
              `}
              placeholder="Enter Chat ID"
              value={chatId}
              onChange={(e) => setChatId(e.target.value)}
            />
          </div>
        </div>

        <button
          type="button"
          onClick={integrationChange}
          className={`
            w-full py-2.5 rounded-lg text-white font-medium transition-colors duration-200 cursor-pointer
            bg-green-600 hover:bg-green-500 active:bg-green-700
          `}
        >
          Save
        </button>

      </div>
    </div>
  ) : (
    <div className={`flex flex-1 min-h-screen min-w-screen justify-center items-center`}>
      <div className={`w-75 bg-white rounded-2xl shadow-lg p-6 space-y-6`}>
        Loading...
      </div>
    </div>
  );
}

export default TelegramConnect;