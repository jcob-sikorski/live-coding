/**
 * VendingMachine follows the Singleton pattern to ensure a single source of truth 
 * for inventory and cash reserves. It utilizes a thread-safe collection (e.g., ConcurrentHashMap)
 * to manage product quantities and prices, ensuring consistency across concurrent transactions.
 */

/**
 * Strategy Pattern: Decouples the change-calculation logic from the machine logic.
 * This allows the machine to swap between algorithms (e.g., Greedy approach for 
 * standard change vs. an Exact Change Only algorithm when low on specific coins).
 */

/**
 * State Pattern: Manages the machine's lifecycle (Idle, MoneyInserted, Dispensing, OutOfOrder).
 * Each state encapsulates its own behavior, preventing complex if-else branching 
 * and making the machine's transitions predictable and extensible.
 */

/**
 * Facade Pattern: Provides a simplified, unified interface for different actors.
 * It abstracts the complex internal state transitions and hardware-level logic 
 * into high-level methods for the 'Buyer' (select, pay) and the 'Operator' (restock, withdraw cash).
 */

/**
 * Factory Method: While the machine dispenses existing items, a Factory is used 
 * internally to instantiate specific State objects or to generate Transaction logs, 
 * keeping the core VendingMachine class clean and focused on orchestration.
 */