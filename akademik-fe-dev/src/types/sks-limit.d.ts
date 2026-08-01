// {
//     "id": "09d9c8cf-e2ac-4db7-8cf9-7420ce2fdcb8",
//     "ips_min": 2,
//     "ips_max": 2,
//     "sks_limit": 2,
//     "created_at": 1742630103154,
//     "updated_at": 1742635600489,
//     "deleted_at": null
// }

interface ISksLimit {
  id: string;
  ips_min: number;
  ips_max: number;
  sks_limit: number;
  created_at: number;
  updated_at: number | null;
  deleted_at: number | null;
}
interface IQueryParamsSksLimits  extends QueryParam {
  sort?: string | null;
}

type ISksLimitForm = {
  ips_min: string;
  ips_max: string;
  sks_limit: string;
};
