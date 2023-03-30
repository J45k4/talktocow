import { useState, useEffect } from 'react';
import { serverUrl } from './config';
import { getHeaders } from './headers';

type QueryParams = Record<string, string | number>;

export interface FetchOptions<RequestBody> {
  path: string;
  query?: QueryParams;
  body?: RequestBody;
}

export interface FetchResult<Data> {
  data?: Data;
  error?: Error;
  loading: boolean;
  refetch: () => void;
}

export const useFetch = <Data = any, RequestBody = any>(
  options: FetchOptions<RequestBody>
): FetchResult<Data> => {
  const [data, setData] = useState<Data>();
  const [error, setError] = useState<Error>();
  const [loading, setLoading] = useState(true);

  const fetchData = () => {
    setLoading(true);

    let url = new URL(options.path, serverUrl);

    if (options.query) {
      url.search = new URLSearchParams(options.query as any).toString();
    }

    fetch(url.toString(), {
      method: options.body ? 'POST' : 'GET',
      headers: getHeaders(),
      body: options.body ? JSON.stringify(options.body) : undefined
    })
      .then(response => response.json())
      .then(jsonData => setData(jsonData))
      .catch(error => setError(error))
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    fetchData();
  }, [serverUrl, options.path, options.query, options.body]);

  return { data, error, loading, refetch: fetchData };
};
