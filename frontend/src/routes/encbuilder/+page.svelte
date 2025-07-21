<script lang="ts">
    type Monster = {
        index: string;
        name: string;
        size: string;
        type: string;
        alignment: string;
        armor_class: string;
        hit_points: string;
        hit_dice: string;
        speed: string;
        actions: string;
        strength: string;
        dexterity: string;
        constitution: string;
        intelligence: string;
        wisdom: string;
        charisma: string;
        damage_vulnerabilities: string;
        damage_resistances: string;
        damage_immunities: string;
        condition_immunities: string;
        senses: string;
        languages: string;
        challenge_rating: string;
        xp: string;
        special_abilities: string;
        legendary_actions: string;
        image: string;
        url: string;
    };

    const selectedAttributes: { key: keyof Monster; label: string }[] = [
        { key: "size", label: "Size" },
        { key: "type", label: "Type" },
        { key: "alignment", label: "Alignment" },
        { key: "armor_class", label: "Armor Class" },
        { key: "hit_points", label: "Hit Points" },
        { key: "hit_dice", label: "Hit Dice" },
        { key: "speed", label: "Speed" },
        { key: "actions", label: "Actions" },
        { key: "strength", label: "Strength" },
        { key: "dexterity", label: "Dexterity" },
        { key: "constitution", label: "Constitution" },
        { key: "intelligence", label: "Intelligence" },
        { key: "wisdom", label: "Wisdom" },
        { key: "charisma", label: "Charisma" },
        { key: "damage_vulnerabilities", label: "Damage Vulnerabilities" },
        { key: "damage_resistances", label: "Damage Resistances" },
        { key: "damage_immunities", label: "Damage Immunities" },
        { key: "condition_immunities", label: "Condition Immunities" },
        { key: "senses", label: "Senses" },
        { key: "languages", label: "Languages" },
        { key: "special_abilities", label: "Special Abilities" },
        { key: "legendary_actions", label: "Legendary Actions" },
    ];

    import { onMount } from "svelte";
    let monsters: Monster[] = [];
    let searchResults: Monster[] = [];
    let monsterCounts: { [index: string]: number } = {};
    let monstersInTray: MonsterInstance[] = [];
    let searchQuery = "";
    let showSearch = false;
    let showAllMonsters = false;

    // Card removal code
    let showConfirmDialog = false;
    let monsterToRemove: number | null = null;

    interface MonsterInstance extends Monster {
        currentHP: number;
        maxHP: number;
    }

    //
    // Environmental effects microservice handler
    //

    let environmentEffect = "";
    let isGeneratingEffect = false;
    let effectError = "";

    async function generateEffect(location: string = "") {
        isGeneratingEffect = true;
        effectError = "";
        try {
            const response = await fetch(
                `http://localhost:8080/api/environment?location=${location}`
            );
            if (!response.ok) {
                throw new Error("Failed to generate environment effect");
            }
            const data = await response.json();
            if (data.success) {
                environmentEffect = `${data.data.effect}: ${data.data.description}`;
            } else {
                throw new Error(
                    data.error || "Failed to generate environment effect"
                );
            }
        } catch (err: any) {
            effectError = err?.message || "An unknown error occurred";
            console.error("Environment effect generation error:", err);
        } finally {
            isGeneratingEffect = false;
        }
    }

    //
    // Combat tactics microservice handler
    //

    // Add these new variables for the tactics generator
    let tactic = "";
    let isGeneratingTactic = false;
    let tacticError = "";

    async function generateTactic(
        aggression: string = "",
        isGroup: boolean = false
    ) {
        isGeneratingTactic = true;
        tacticError = "";
        try {
            const response = await fetch(
                `http://localhost:8080/api/tactics?aggression=${aggression}&group=${isGroup}`
            );
            if (!response.ok) {
                throw new Error("Failed to generate tactic");
            }
            const data = await response.json();
            if (data.success) {
                tactic = `${data.data.tactic}: ${data.data.description}`;
            } else {
                throw new Error(data.error || "Failed to generate tactic");
            }
        } catch (err: any) {
            tacticError = err?.message || "An unknown error occurred";
            console.error("Tactic generation error:", err);
        } finally {
            isGeneratingTactic = false;
        }
    }

    //
    // Plot twist microservice handling
    //

    let plotTwist = "";
    let isGeneratingPlot = false;
    let plotError = "";

    async function generatePlotTwist(category = "combat") {
        isGeneratingPlot = true;
        plotError = "";
        try {
            const response = await fetch(
                `http://localhost:8080/api/plot-twist?category=${category}`
            );
            if (!response.ok) {
                throw new Error("Failed to generate plot twist");
            }
            const data = await response.json();
            if (data.success) {
                plotTwist = data.data.text;
            } else {
                throw new Error(data.error || "Failed to generate plot twist");
            }
        } catch (err: any) {
            // Type assertion for error
            plotError = err?.message || "An unknown error occurred";
            console.error("Plot twist generation error:", err);
        } finally {
            isGeneratingPlot = false;
        }
    }

    //
    // Name generator microservice handling
    //

    let generatedName = "";
    let isGeneratingName = false;
    let nameError = "";

    async function generateName() {
        isGeneratingName = true;
        nameError = "";
        try {
            const response = await fetch(
                "http://localhost:8080/api/generate-name"
            );
            if (!response.ok) {
                throw new Error("Failed to generate name");
            }
            const data = await response.json();
            generatedName = data.name;
        } catch (err) {
            nameError = "Failed to generate name. Please try again.";
            console.error("Name generation error:", err);
        } finally {
            isGeneratingName = false;
        }
    }

    let isMinimized = false;

    function toggleMinimize() {
        isMinimized = !isMinimized;
    }

    // Add HP control functions
    function incrementHP(index: number) {
        monstersInTray = monstersInTray.map((monster, i) => {
            if (i === index) {
                return {
                    ...monster,
                    currentHP: Math.min(monster.currentHP + 1, monster.maxHP),
                };
            }
            return monster;
        });
    }

    function decrementHP(index: number) {
        monstersInTray = monstersInTray.map((monster, i) => {
            if (i === index) {
                return {
                    ...monster,
                    currentHP: Math.max(monster.currentHP - 1, 0),
                };
            }
            return monster;
        });
    }
    function updateHP(index: number, newHP: string | number) {
        const numericValue =
            typeof newHP === "string"
                ? parseInt(newHP.replace(/[^0-9]/g, "")) || 0
                : newHP;

        monstersInTray = monstersInTray.map((monster, i) => {
            if (i === index) {
                const validatedHP = Math.max(
                    0,
                    Math.min(numericValue, monster.maxHP)
                );
                return {
                    ...monster,
                    currentHP: validatedHP,
                };
            }
            return monster;
        });
    }

    function initiateRemove(index: number) {
        monsterToRemove = index;
        showConfirmDialog = true;
    }

    function confirmRemove() {
        if (monsterToRemove !== null) {
            monstersInTray = monstersInTray.filter(
                (_, i) => i !== monsterToRemove
            );
        }
        showConfirmDialog = false;
        monsterToRemove = null;
    }

    // Function to cancel removal
    function cancelRemove() {
        showConfirmDialog = false;
        monsterToRemove = null;
    }

    // Card expansion code
    let expandedCards: { [key: number]: boolean } = {};

    function toggleCard(index: number) {
        expandedCards[index] = !expandedCards[index];
        expandedCards = expandedCards; // Trigger reactivity
    }

    function shouldHideAttribute(attributeKey: string): boolean {
        const breakpointIndex = selectedAttributes.findIndex(
            (attr) => attr.key === "actions"
        );
        const attributeIndex = selectedAttributes.findIndex(
            (attr) => attr.key === attributeKey
        );
        return attributeIndex > breakpointIndex;
    }

    // Interfaces for monster actions
    interface DamageType {
        index: string;
        name: string;
        url: string;
    }

    interface Damage {
        damage_dice: string;
        damage_type: DamageType;
    }

    interface MonsterAction {
        name: string;
        desc: string;
        attack_bonus: number;
        damage: Damage[];
    }

    // Fetch monster data from the Go backend on component mount
    onMount(async () => {
        const response = await fetch("http://localhost:8080/api/monsters");
        monsters = await response.json();
    });

    // Show instructions on page load
    let showInstructions = true;
    let searchButtonsDisabled = true;

    // An action allowing the input node to be autofocused
    const autofocus = (node: HTMLElement) => {
        node.focus();
    };

    // Parse and format the monster actions
    function formatActions(actionsString: string): string {
        try {
            // Parse the JSON string into an array of actions
            const actions: MonsterAction[] = JSON.parse(actionsString);

            // Format each action
            return actions
                .map((action) => {
                    let formattedAction = `${action.name}\n`;

                    if (action.attack_bonus) {
                        formattedAction += `Attack Bonus: +${action.attack_bonus}\n`;
                    }

                    if (action.damage && action.damage.length > 0) {
                        action.damage.forEach((damage) => {
                            formattedAction += `Damage: ${damage.damage_dice} ${damage.damage_type.name.toLowerCase()}\n`;
                        });
                    }

                    formattedAction += `${action.desc}\n`;
                    return formattedAction;
                })
                .join("\n");
        } catch (e) {
            // If parsing fails, return the original string
            return actionsString;
        }
    }

    function openSearch() {
        showSearch = true;
        showAllMonsters = false;

        // Clear the search results and search query on opening search
        searchQuery = "";
        searchResults = [];
    }

    function openShowAll() {
        showAllMonsters = true;
        showSearch = false;

        // Initialize counts for all monsters
        monsters.forEach((monster) => {
            if (!(monster.index in monsterCounts)) {
                monsterCounts[monster.index] = 1;
            }
        });
    }

    function closeSearch() {
        showSearch = false;
        showAllMonsters = false;
        searchResults = [];
    }

    function incrementCount(index: string) {
        monsterCounts[index] = Math.min(monsterCounts[index] + 1, 10);
    }

    function decrementCount(index: string) {
        monsterCounts[index] = Math.max(monsterCounts[index] - 1, 1);
    }

    function searchMonsters() {
        // Only filter if we're not showing all monsters
        if (!showAllMonsters) {
            searchResults = monsters.filter((monster) =>
                monster.name.toLowerCase().includes(searchQuery.toLowerCase())
            );

            // Initialize counts for search results
            searchResults.forEach((monster) => {
                if (!(monster.index in monsterCounts)) {
                    monsterCounts[monster.index] = 1;
                }
            });
        }
    }

    function addMonster(index: string) {
        const count = monsterCounts[index];
        const monster = monsters.find((m) => m.index === index);
        if (monster) {
            const newMonsters = Array(count)
                .fill(null)
                .map(() => ({
                    ...monster,
                    currentHP: parseInt(monster.hit_points) || 0,
                    maxHP: parseInt(monster.hit_points) || 0,
                }));
            monstersInTray = [...monstersInTray, ...newMonsters];
        }
    }

    function closeInstructions() {
        showInstructions = false;
        searchButtonsDisabled = false;
    }

    function handleShowAll() {
        openShowAll();
    }
</script>

<div class="container">
    <!-- Menu Bar with Buttons -->
    <div class="menu-bar">
        <button
            class:active-search={showSearch}
            on:click={openSearch}
            disabled={searchButtonsDisabled}
        >
            Search
        </button>
        <button
            class:active-search={showAllMonsters}
            on:click={handleShowAll}
            disabled={searchButtonsDisabled}
        >
            Show All
        </button>
        <h2 class="frontpage-title">Battle Manager - Monster Tray</h2>
    </div>

    <!-- microservices interface code -->
    <div class="name-generator-wrapper">
        <div class="name-generator">
            <button class="name-generator-minimize" on:click={toggleMinimize}>
                {isMinimized ? "▼" : "▲"}
            </button>
            {#if !isMinimized}
                <!-- Name Generator Section -->
                <p
                    style="color: black; font-weight: bold; font-size: 0.875rem; margin: 0.5rem 0 0.25rem;"
                >
                    Name Generator
                </p>
                <div class="name-generator-controls">
                    <button
                        class="name-generator-btn"
                        on:click={generateName}
                        disabled={isGeneratingName}
                    >
                        {isGeneratingName
                            ? "Generating..."
                            : "Generate Random Name"}
                    </button>
                    {#if generatedName}
                        <span class="name-generator-result"
                            >{generatedName}</span
                        >
                    {/if}
                </div>
                {#if nameError}
                    <div class="name-generator-error">{nameError}</div>
                {/if}

                <!-- Plot Twist Generator Section -->
                <p
                    style="color: black; font-weight: bold; font-size: 0.875rem; margin: 0.5rem 0 0.25rem;"
                >
                    Plot Twists
                </p>
                <div
                    class="name-generator-controls"
                    style="flex-direction: column; gap: 0.5rem;"
                >
                    <div
                        style="display: flex; gap: 0.5rem; flex-wrap: wrap; justify-content: center;"
                    >
                        <button
                            class="name-generator-btn"
                            on:click={() => generatePlotTwist("combat")}
                            disabled={isGeneratingPlot}
                        >
                            {isGeneratingPlot ? "Generating..." : "Combat"}
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generatePlotTwist("treasure")}
                            disabled={isGeneratingPlot}
                        >
                            Treasure
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generatePlotTwist("social")}
                            disabled={isGeneratingPlot}
                        >
                            Social
                        </button>
                    </div>
                    {#if plotTwist}
                        <span
                            class="name-generator-result"
                            style="width: 100%; box-sizing: border-box;"
                        >
                            {plotTwist}
                        </span>
                    {/if}
                </div>
                {#if plotError}
                    <div class="name-generator-error">{plotError}</div>
                {/if}

                <!-- Combat Tactics Generator Section -->
                <p
                    style="color: black; font-weight: bold; font-size: 0.875rem; margin: 0.5rem 0 0.25rem;"
                >
                    Combat Tactics
                </p>
                <div
                    class="name-generator-controls"
                    style="flex-direction: column; gap: 0.5rem;"
                >
                    <!-- Solo Tactics -->
                    <div
                        style="display: flex; gap: 0.5rem; flex-wrap: wrap; justify-content: center;"
                    >
                        <button
                            class="name-generator-btn"
                            on:click={() => generateTactic("aggressive", false)}
                            disabled={isGeneratingTactic}
                        >
                            Aggressive Solo
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateTactic("cautious", false)}
                            disabled={isGeneratingTactic}
                        >
                            Cautious Solo
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateTactic("afraid", false)}
                            disabled={isGeneratingTactic}
                        >
                            Afraid Solo
                        </button>
                    </div>
                    <!-- Group Tactics -->
                    <div
                        style="display: flex; gap: 0.5rem; flex-wrap: wrap; justify-content: center;"
                    >
                        <button
                            class="name-generator-btn"
                            on:click={() => generateTactic("aggressive", true)}
                            disabled={isGeneratingTactic}
                        >
                            Aggressive Group
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateTactic("cautious", true)}
                            disabled={isGeneratingTactic}
                        >
                            Cautious Group
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateTactic("afraid", true)}
                            disabled={isGeneratingTactic}
                        >
                            Afraid Group
                        </button>
                    </div>
                    {#if tactic}
                        <span
                            class="name-generator-result"
                            style="width: 100%; box-sizing: border-box;"
                        >
                            {tactic}
                        </span>
                    {/if}
                </div>
                {#if tacticError}
                    <div class="name-generator-error">{tacticError}</div>
                {/if}

                <!-- Environment Effects Generator Section -->
                <p
                    style="color: black; font-weight: bold; font-size: 0.875rem; margin: 0.5rem 0 0.25rem;"
                >
                    Environment Effects
                </p>
                <div
                    class="name-generator-controls"
                    style="flex-direction: column; gap: 0.5rem;"
                >
                    <div
                        style="display: flex; gap: 0.5rem; flex-wrap: wrap; justify-content: center;"
                    >
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("dungeon")}
                            disabled={isGeneratingEffect}
                        >
                            Dungeon
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("forest")}
                            disabled={isGeneratingEffect}
                        >
                            Forest
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("coastal")}
                            disabled={isGeneratingEffect}
                        >
                            Coastal
                        </button>
                    </div>
                    <div
                        style="display: flex; gap: 0.5rem; flex-wrap: wrap; justify-content: center;"
                    >
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("mountain")}
                            disabled={isGeneratingEffect}
                        >
                            Mountain
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("urban")}
                            disabled={isGeneratingEffect}
                        >
                            Urban
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("swamp")}
                            disabled={isGeneratingEffect}
                        >
                            Swamp
                        </button>
                    </div>
                    <div
                        style="display: flex; gap: 0.5rem; flex-wrap: wrap; justify-content: center;"
                    >
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("desert")}
                            disabled={isGeneratingEffect}
                        >
                            Desert
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect("arctic")}
                            disabled={isGeneratingEffect}
                        >
                            Arctic
                        </button>
                        <button
                            class="name-generator-btn"
                            on:click={() => generateEffect()}
                            disabled={isGeneratingEffect}
                        >
                            Random
                        </button>
                    </div>
                    {#if environmentEffect}
                        <span
                            class="name-generator-result"
                            style="width: 100%; box-sizing: border-box;"
                        >
                            {environmentEffect}
                        </span>
                    {/if}
                </div>
                {#if effectError}
                    <div class="name-generator-error">{effectError}</div>
                {/if}
            {/if}
        </div>
    </div>
</div>
<!-- END microservices interface code -->

<!-- Search Dialog -->
{#if showSearch || showAllMonsters}
    <div class="search-dialog">
        {#if showSearch}
            <input
                class="search-input"
                type="text"
                placeholder="Name of monster..."
                bind:value={searchQuery}
                on:input={searchMonsters}
                use:autofocus
            />
        {/if}
        <div class="search-results">
            {#if showAllMonsters}
                <ul class="monster-list">
                    {#each monsters as monster}
                        <li class="monster-item">
                            <b>{monster.name}</b>
                            <div class="controls">
                                <button
                                    on:click={() =>
                                        decrementCount(monster.index)}>-</button
                                >
                                <input
                                    type="number"
                                    min="1"
                                    max="10"
                                    bind:value={monsterCounts[monster.index]}
                                />
                                <button
                                    on:click={() =>
                                        incrementCount(monster.index)}>+</button
                                >
                                <button
                                    on:click={() => addMonster(monster.index)}
                                    >Add</button
                                >
                            </div>
                        </li>
                    {/each}
                </ul>
            {:else if searchResults.length > 0}
                <ul class="monster-list">
                    {#each searchResults as monster}
                        <li class="monster-item">
                            <b>{monster.name}</b>
                            <div class="controls">
                                <button
                                    on:click={() =>
                                        decrementCount(monster.index)}>-</button
                                >
                                <input
                                    type="number"
                                    min="1"
                                    max="10"
                                    bind:value={monsterCounts[monster.index]}
                                />
                                <button
                                    on:click={() =>
                                        incrementCount(monster.index)}>+</button
                                >
                                <button
                                    on:click={() => addMonster(monster.index)}
                                    >Add</button
                                >
                            </div>
                        </li>
                    {/each}
                </ul>
            {:else if showSearch}
                <p class="monster-item">No results found</p>
            {/if}
        </div>

        <button class="search-dialog-close-btn" on:click={closeSearch}
            >Close</button
        >
    </div>
{/if}

<!-- Central Tray -->
<div class="tray">
    {#if monstersInTray.length > 0}
        <p>Monsters in tray:</p>
        {#each monstersInTray as monster, index}
            <div class="card">
                <div class="monster-name">
                    {monster.name}
                </div>
                <div class="hp-tracker">
                    <div class="hp-controls">
                        <button
                            class="hp-btn"
                            on:click={() => decrementHP(index)}
                            aria-label="Decrease HP"
                        >
                            -
                        </button>
                        <div class="hp-display">
                            <input
                                type="text"
                                inputmode="numeric"
                                pattern="[0-9]*"
                                class="hp-input"
                                value={monster.currentHP}
                                on:input={(e) => {
                                    const input = e.currentTarget;
                                    // Remove any non-numeric characters
                                    input.value = input.value.replace(
                                        /[^0-9]/g,
                                        ""
                                    );

                                    // Get the numeric value
                                    const numericValue =
                                        parseInt(input.value) || 0;

                                    // If the entered value exceeds maxHP, immediately update the input value
                                    if (numericValue > monster.maxHP) {
                                        input.value = monster.maxHP.toString();
                                    }

                                    // Update the monster's HP
                                    updateHP(index, input.value);
                                }}
                            />
                            <span class="hp-separator">/</span>
                            <span class="hp-max">{monster.maxHP}</span>
                        </div>
                        <button
                            class="hp-btn"
                            on:click={() => incrementHP(index)}
                            aria-label="Increase HP"
                        >
                            +
                        </button>
                    </div>
                    <div class="hp-bar">
                        <div
                            class="hp-bar-fill"
                            style="width: {(monster.currentHP / monster.maxHP) *
                                100}%; 
                               background-color: {monster.currentHP <=
                            monster.maxHP * 0.25
                                ? '#ff4444'
                                : monster.currentHP <= monster.maxHP * 0.5
                                  ? '#ffaa00'
                                  : '#44ff44'};"
                        />
                    </div>
                </div>
                <div class="card-body-text">
                    {#each selectedAttributes as attribute}
                        {@const isHidden = shouldHideAttribute(attribute.key)}
                        {#if !isHidden || expandedCards[index]}
                            <div class="attribute">
                                <strong>{attribute.label}:&nbsp;</strong>
                                {#if attribute.key === "actions"}
                                    {formatActions(monster[attribute.key])}
                                {:else}
                                    {monster[attribute.key]}
                                {/if}
                            </div>
                        {/if}
                    {/each}
                    <button
                        class="expander-btn"
                        on:click={() => toggleCard(index)}
                    >
                        {expandedCards[index] ? "▲ Show Less" : "▼ Show More"}
                    </button>
                </div>
                <button
                    class="remove-card-btn"
                    on:click={() => initiateRemove(index)}
                >
                    Remove
                </button>
            </div>
        {/each}
    {:else}
        <div class="empty-tray-container">
            <div class="empty-tray-msg">
                The monster tray is currently empty...
            </div>
        </div>
    {/if}
</div>

<!-- Instruction Dialog -->
{#if showInstructions}
    <div class="instruction-dialog">
        <h4>Welcome to the Battle Manager!</h4>
        <p>
            To add a monster to your battle, click the <b>Search</b>
            button and type the name of the monster you want to add, then press
            <b>Submit</b>.
        </p>
        <p>
            If you can’t think of a monster, try searching the word <b>goblin</b
            >.
        </p>
        <p>
            If you want to see a list of all available monsters, click <b
                >Show All</b
            >.
        </p>

        <button on:click={closeInstructions}>Okay</button>
    </div>
{/if}

<!-- Removal confirmation dialog -->
{#if showConfirmDialog}
    <div class="instruction-dialog">
        <div class="confirmation-content">
            <p>
                Are you sure you want to remove this monster card from the tray?
            </p>
            <p>
                Removing the monster will result in losing the data of any
                changed hit points, such as damage sustained in combat.
            </p>
            <div class="confirmation-buttons">
                <button class="confirm-btn" on:click={confirmRemove}
                    >Yes, Remove</button
                >
                <button class="cancel-btn" on:click={cancelRemove}
                    >Cancel</button
                >
            </div>
        </div>
    </div>
{/if}
