import { useParams } from "react-router-dom";

export const Profile = () => {
  const { uuid } = useParams<{ uuid: string }>();
};
