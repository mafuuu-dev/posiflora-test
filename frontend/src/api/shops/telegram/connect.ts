import HTTP from "@core/http";

type Request = {
  bot_token: string;
  chat_id: string;
  is_enabled: boolean;
};

const IntegrationConnect = async (shopId: number, request: Request) => {
  return HTTP.getInstance().axios().post(`/shops/${shopId}/telegram/connect`, request);
};

export default IntegrationConnect;

