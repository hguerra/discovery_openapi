import { Configuration, EventsApi, OperationsApi, TicketsApi } from './internal';

const configuration = new Configuration({
  basePath: 'http://localhost:8080/v1',
});

export const eventsApi = new EventsApi(configuration);
export const operationsApi = new OperationsApi(configuration);
export const ticketsApi = new TicketsApi(configuration);
