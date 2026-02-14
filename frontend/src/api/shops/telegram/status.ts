import HTTP from "@core/http";

const IntegrationStatus = async (shopId: number) => {
  return HTTP.getInstance().axios().get(`/shops/${shopId}/telegram/status`);
};

export default IntegrationStatus;

