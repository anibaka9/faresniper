# **Technical Specification: Personal Flight and Transit Route Scanner (Faresniper)**

**Project Goal:** Development of an automated system to find anomalously cheap flight tickets (regular flights, low-cost carriers, and charters) based on specified parameters, incorporating ground transportation logic and complex routing.

**Tech Stack:**

* **Backend/Workers:** Go (for fast, concurrent API polling and scraping).
* **Database:** SQLite (for storing price history and calculating statistical metrics).
* **Interface:** Web admin panel (Go + HTML templates, chi router). Telegram Bot in later phases.

### ---

## Phase 1: MVP (Basic Route Tracker)

**Concept:** Poll Travelpayouts `prices_for_dates` for a fixed list of routes, accumulate price history, surface price data in the admin panel, and flag anomalously cheap tickets.

**Data Sources:**
* **Travelpayouts IATA databases** — one-time load of reference data: countries, cities, airports, airlines.
* **Travelpayouts `prices_for_dates` API** — periodic polling for price snapshots per route. Returns airline, airport codes, price, transfers, departure time, booking link.

---

### User Stories

**Route watchlist**

> As a user, I want to add a flight route (origin + destination IATA) to a watchlist, so that the system starts collecting price data for it.

Acceptance criteria:
- I can open the admin panel and see the list of watched routes
- I can add a new route by entering two IATA codes
- I can deactivate a route (stop polling without deleting history)

---

**Price data visibility — before anomaly detection is ready**

> As a user, I want to see current prices for a watched route as a price calendar, so that I can immediately understand which departure dates are cheap or expensive.

Acceptance criteria:
- After the first polling cycle I can open a route page and see a calendar view for the current and next month
- Each day cell shows the cheapest available price for that departure date
- Days with no data are visually distinct (grey / empty)
- I can see which airline offers the cheapest price for each day

> As a user, I want to see a table of recent price snapshots for a route, so that I can inspect raw data and verify the system is working.

Acceptance criteria:
- Table shows: departure date, price, airline, transfers, observed_at, booking link
- Sortable by price and by departure date
- Shows when the last polling cycle ran

> As a user, I want to see how the price for a specific departure date has changed over multiple polling cycles, so that I can understand whether prices are rising or falling.

Acceptance criteria:
- After at least 2 polling cycles: I can see a price trend per departure date (e.g., a small chart or a list of observations over time)
- Available once enough polling history accumulates (not on day 1)

---

**Anomaly detection**

> As a user, I want the system to automatically flag tickets that are significantly cheaper than the historical norm for that route and month, so that I don't have to manually compare prices.

Acceptance criteria:
- Anomaly detection activates only after at least 50 price snapshots are accumulated for a route+month combination
- A ticket is flagged as anomaly if its price is more than 25% below the median of all observed prices for that route+month
- Flagged tickets appear in a dedicated Anomalies feed in the admin panel
- Each anomaly shows: route, departure date, price, airline, % below median, booking link
- Anomalies feed is sorted by largest discount first

---

**Polling infrastructure**

> As a user, I want price data to be collected automatically on a schedule without manual intervention, so that the history builds up while I'm not at my computer.

Acceptance criteria:
- A cron job (OS crontab) runs the polling script every 6–12 hours
- Each run: fetches `prices_for_dates` for current month + next month for all active routes
- Request params: `limit=1000`, `sorting=price`, `one_way=true`, `currency=usd`
- Results are saved as price snapshots with `source="travelpayouts"` and `observed_at` timestamp
- If the API returns an error, the run logs it and continues with the next route (no crash)

---

### Implementation Steps

1. **Reference data (done)**
   * ✅ DB schema: countries, cities, airports, airlines, price_snapshots
   * ✅ Load IATA JSON files (countries, cities, airports, airlines) into DB via `cmd/travelpayouts/iata_save`

2. **Route watchlist**
   * 🔲 Add `watched_routes` table: origin IATA, destination IATA, active flag
   * 🔲 Admin panel UI: list, add, deactivate routes
   * MVP: start with one hardcoded route (BSZ → ALA) to validate the pipeline

3. **Price polling**
   * 🔲 Wire `prices_for_dates` for watchlist routes (2 requests per route: current + next month)
   * 🔲 Save each returned ticket as price snapshot with `source`, `observed_at`
   * 🔲 OS crontab entry to run polling every 6–12 hours
   * ✅ Basic price saving structure (`price_snapshots` table, `prices_load`/`prices_save`) — needs wiring to watchlist

4. **Admin panel — price visibility**
   * ✅ Web server (`cmd/web`), HTML templates, CRUD for reference entities
   * 🔲 Price calendar view: month grid, cheapest price per day, per route
   * 🔲 Snapshots table: filterable, sortable, with booking link
   * 🔲 Price trend view: history of observed prices for a specific departure date

5. **Anomaly detection**
   * 🔲 Calculate median price per route+month from accumulated snapshots
   * 🔲 Minimum 50 snapshots threshold before activating detection
   * 🔲 Flag tickets at price < median × 0.75
   * 🔲 Anomalies feed in admin panel: sorted by discount, with booking link

---

### Architecture: Polling

**MVP: OS crontab**
- Build the polling binary: `go build ./cmd/travelpayouts/prices_load`
- Add to crontab: runs every 6 hours, logs output to a file
- DB file lives in the project directory on the local machine
- Zero new infrastructure

**Next step (when always-on is needed): VPS + long-running process**
- Single Go binary with internal `time.Ticker` for polling loop
- Deploy to a cheap VPS (~$5/month)
- Docker optional — add only when portability becomes a real need

---

## Phase 1.5: Route Intelligence (aeroroutes.com Monitoring)

**Concept:** Monitor aviation news to get early warning about new low-cost carrier routes.

* Cron job scrapes aeroroutes.com news feed for watchlist airports.
* New uncovered low-cost route → notification in admin panel.
* Human-in-the-loop: Aleksei reviews and decides whether to build a new scraper.

---

## Phase 2: Direct Airline Scraping

**Concept:** Add real-time price sources for watched routes. Same anomaly detection pipeline, better data quality.

* Scrapers for specific low-cost carrier websites (AirAsia, VietJet, etc.)
* Snapshots saved with `source="airline_site"`, `is_confirmed=true`
* Confirmed prices improve baseline accuracy over time (real-time, includes taxes)

---

## Phase 3: Telegram Bot

* Anomaly alerts sent as Telegram messages with booking link.
* Bot commands for watchlist management.
* Admin panel remains as companion tool.

---

## Phase 4: Advanced Search (Clusters and Ground Penalty)

**Concept:** Find cheaper entry/exit points via neighboring airports, factoring in ground transport cost and time.

* Cluster config: static file grouping airports into zones (e.g., ["FRU", "ALA"], ["BKK", "DMK", "UTP"]).
* Ground penalty: cost + time constants per connection (e.g., Surat Thani → Samui: +$15, +3h).
* Scoring: Ticket Price + Ground Penalty. Flag if cheaper than direct and within time limit.

---

## Phase 5: Complex Routing (Open-Jaw and Transit Corridors)

**Concept:** Assemble itineraries arriving in one city and departing from another via ground transport.

* Transit corridors: CAN ↔ SZX ↔ HKG, KIX ↔ NGO ↔ TYO, etc.
* Stitching: cheap outbound to any corridor point + cheap inbound from another point N days later + fixed ground cost.
* Discard if ground cost or transit time exceeds limits.

---

## Phase 6: Charter Sweeping

**Concept:** Intercept price dumps from tour operators 24–72 hours before departure.

* Hard price threshold trigger (e.g., any ticket under $50) — no median baseline.
* Fixed return date constraints (7/10/14 days).
