/* tslint:disable */
/* eslint-disable */
/**
 * Redocly Museum API
 * An imaginary, but delightful Museum API for interacting with museum services and information. Built with love by Redocly.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: team@redocly.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


// May contain unused imports in some cases
// @ts-ignore
import type { TicketType } from './ticket-type';

/**
 * Request payload used for purchasing museum tickets.
 * @export
 * @interface BuyMuseumTicketsRequest
 */
export interface BuyMuseumTicketsRequest {
    /**
     * 
     * @type {TicketType}
     * @memberof BuyMuseumTicketsRequest
     */
    'ticketType': TicketType;
    /**
     * Identifier for a special event.
     * @type {string}
     * @memberof BuyMuseumTicketsRequest
     */
    'eventId'?: string;
    /**
     * 
     * @type {string}
     * @memberof BuyMuseumTicketsRequest
     */
    'ticketDate': string;
    /**
     * Email address for ticket purchaser.
     * @type {string}
     * @memberof BuyMuseumTicketsRequest
     */
    'email': string;
    /**
     * Phone number for the ticket purchaser (optional).
     * @type {string}
     * @memberof BuyMuseumTicketsRequest
     */
    'phone'?: string;
}


