"use client";

import { useParams } from "next/navigation"
import TelegramConnect from "@/components/TelegramConnect";

const ShopTelegramPage = () => {
  const params = useParams();
  const shopId = Number(params.id);

  return (
    <TelegramConnect shopId={shopId}/>
  );
};

export default ShopTelegramPage;
