# **Technical Specification: Personal Flight and Transit Route Scanner (Faresniper)**

**Project Goal:** Development of an automated system to find anomalously cheap flight tickets (regular flights, low-cost carriers, and charters) based on specified parameters, incorporating ground transportation logic and complex routing.  
**Proposed Tech Stack:**

* **Backend/Workers:** Go (for fast, concurrent API polling and scraping).  
* **Database:** PostgreSQL (for storing price history and calculating statistical metrics) or SQLite (for the MVP).  
* **Interface:** Telegram Bot (for configuration and alerts). Potential future dashboard in TypeScript.

### ---

**Phase 1: MVP (Basic Route Tracker)**

**Concept:** Collect prices for a fixed list of destinations, build a historical baseline, and send notifications when the price drops below the calculated norm.

* **Data Sources:** Aggregator APIs (e.g., Kiwi Tequila) or direct requests to low-cost carrier websites (AirAsia, VietJet, etc.).  
* **Workflow:**  
  1. Define a list of pairs (e.g., BKK \-\> KIX, BKK \-\> NRT).  
  2. A cron job polls the sources every N hours for specific date ranges (e.g., departures in the next 3-6 months).  
  3. Retrieved prices are saved to the database.  
  4. The system calculates the median ticket price for each route per month.  
* **Alert Trigger:** If the algorithm finds a ticket priced X% below the calculated median for that route and month, the bot sends a Telegram message with a booking link.

### ---

**Phase 2: Advanced Search (Clusters and Ground Penalty)**

**Concept:** Evaluate neighboring airports to find a cheaper entry/exit point, factoring in the time and money spent on ground transportation.

* **Cluster Configuration:** Create a static configuration file (JSON/YAML) grouping airports into logical zones.  
  * *Example:* "Central Asia Departure" \["FRU", "ALA"\], "Thailand Departure" \["BKK", "DMK", "UTP"\].  
* **Ground Penalty:** Add constants for the cost and time of transit between neighboring nodes to the database (e.g., Surat Thani \-\> Samui ferry: \+$15, \+3 hours).  
* **Scoring Logic:**  
  1. The system searches for tickets not only to the target city but to all cities within its cluster.  
  2. The actual cost is calculated: Ticket Price \+ Ground Penalty.  
  3. If the final cost is better than a direct flight (and fits within the allowed transit time limit), an alert is triggered.

### ---

**Phase 3: Complex Routing (Open-Jaw and Transit Corridors)**

**Concept:** Automatically assemble itineraries arriving in one city and departing from another, connected by developed ground infrastructure (e.g., high-speed rail networks).

* **Transit Graphs:** Define railway lines and bus routes connecting airports within macro-regions in the database.  
  * *Example 1 (South China):* CAN (Guangzhou) \<-\> SZX (Shenzhen) \<-\> HKG (Hong Kong). Fixed train cost is set (\~$20).  
  * *Example 2 (Japan):* KIX (Osaka) \<-\> NGO (Nagoya) \<-\> TYO (Tokyo).  
* **Stitching Logic (Async Pipeline):**  
  1. Search for a cheap "Outbound" ticket to any point in the transit corridor (e.g., flight to CAN).  
  2. Search for a cheap "Inbound" ticket from any *other* point in the same corridor N days later (e.g., return from HKG).  
  3. Summation: Outbound Ticket \+ Inbound Ticket \+ Fixed Train Cost Between Cities.  
* **Validation:** Check that the total Ground Cost and transit time do not exceed predefined limits; otherwise, the route is discarded as unfeasible.

### ---

**Phase 4: Charter Sweeping**

**Concept:** A dedicated module to intercept price dumps from tour operators 24–72 hours before departure.

* **Sources:** APIs or parsing of charter ticket exchanges and tour aggregators ("Tickets Only" sections).  
* **Logic Specifics:**  
  * Ignore the median price baseline (charters are inherently cheaper than regular flights).  
  * Trigger based on a hard price threshold (e.g., "any ticket under $50").  
  * Account for charter constraints: notify for One-Way (return flights) or Round-Trip tickets with strictly fixed return dates (e.g., exactly 7/10/14 days).
