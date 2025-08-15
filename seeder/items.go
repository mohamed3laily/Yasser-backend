package seeder

import (
	"log"
	"math/rand"

	"yasser-backend/internal/item-group/item"
	itemcategory "yasser-backend/internal/item-group/item-category"
	"yasser-backend/internal/vendor-group/category"
	"yasser-backend/internal/vendor-group/vendor"

	"gorm.io/gorm"
)

func SeedItemsAndCategories(db *gorm.DB) error {

	if err := db.Where("1 = 1").Delete(&item.ItemSize{}).Error; err != nil {
		log.Printf("❌ Failed to delete item sizes: %v", err)
		return err
	}
	if err := db.Where("1 = 1").Delete(&item.ItemVariant{}).Error; err != nil {
		log.Printf("❌ Failed to delete item variants: %v", err)
		return err
	}
	if err := db.Where("1 = 1").Delete(&item.ItemAddon{}).Error; err != nil {
		log.Printf("❌ Failed to delete item addons: %v", err)
		return err
	}
	if err := db.Where("1 = 1").Delete(&item.Item{}).Error; err != nil {
		log.Printf("❌ Failed to delete items: %v", err)
		return err
	}
	log.Println("🧹 Cleared old items, sizes, variants, and addons successfully")

	// ----------------------
	// ITEM CATEGORIES
	// ----------------------
	restaurantCategories := []itemcategory.ItemsCategory{
		{NameEn: "Appetizers", NameAr: "مقبلات"},
		{NameEn: "Main Courses", NameAr: "أطباق رئيسية"},
		{NameEn: "Burgers", NameAr: "برجر"},
		{NameEn: "Pizzas", NameAr: "بيتزا"},
		{NameEn: "Pasta", NameAr: "معكرونة"},
		{NameEn: "Seafood", NameAr: "مأكولات بحرية"},
		{NameEn: "Desserts", NameAr: "حلويات"},
		{NameEn: "Drinks", NameAr: "مشروبات"},
		{NameEn: "Breakfast", NameAr: "إفطار"},
		{NameEn: "Salads", NameAr: "سلطات"},
		{NameEn: "Sandwiches", NameAr: "ساندوتشات"},
		{NameEn: "Kids Menu", NameAr: "قائمة الأطفال"},
	}

	supermarketCategories := []itemcategory.ItemsCategory{
		{NameEn: "Fruits", NameAr: "فواكه"},
		{NameEn: "Vegetables", NameAr: "خضروات"},
		{NameEn: "Meat & Poultry", NameAr: "لحوم ودواجن"},
		{NameEn: "Seafood", NameAr: "مأكولات بحرية"},
		{NameEn: "Dairy", NameAr: "ألبان"},
		{NameEn: "Bakery", NameAr: "مخبوزات"},
		{NameEn: "Beverages", NameAr: "مشروبات"},
		{NameEn: "Snacks", NameAr: "وجبات خفيفة"},
		{NameEn: "Frozen Foods", NameAr: "مجمدات"},
		{NameEn: "Canned Goods", NameAr: "معلبات"},
		{NameEn: "Cleaning Supplies", NameAr: "منظفات"},
		{NameEn: "Personal Care", NameAr: "عناية شخصية"},
		{NameEn: "Baby Products", NameAr: "منتجات الأطفال"},
	}

	allCats := append(restaurantCategories, supermarketCategories...)
	for _, cat := range allCats {
		var existing itemcategory.ItemsCategory
		if err := db.Where("name_en = ?", cat.NameEn).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&cat).Error; err != nil {
				return err
			}
			log.Printf("✅ Inserted item category: %s", cat.NameEn)
		}
	}

	// ----------------------
	// VENDORS
	// ----------------------
	var vendors []vendor.Vendor
	if err := db.Find(&vendors).Error; err != nil {
		return err
	}
	if len(vendors) == 0 {
		log.Println("⚠️ No vendors found, skipping items seeding")
		return nil
	}

	// ----------------------
	// ITEMS
	// ----------------------

	pictureURLs := []string{
		"https://media02.stockfood.com/largepreviews/NDA4ODkwNzc1/13190025-Seefood-platter-with-seaweed-salsa-verde.jpg",
		"https://images.ctfassets.net/0dkgxhks0leg/4LaYoCoepR6ZwEkAmQFh2F/e82fa8e3c87f0e4cdb3e914b3e766fa0/blog-large-2020veg.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/0/0b/RedDot_Burger.jpg/960px-RedDot_Burger.jpg",
		"https://www.tutorialspoint.com/food_and_beverage_services/images/non_alcoholic_beverages.jpg",
		"https://images.ctfassets.net/721r5zsckzl2/4dJLc584T89nuDNVENFtaO/88f155cf5cf94bb7f6e60252f14524eb/PPZcanadian.jpg",
		"https://www.thecookingcollective.com.au/wp-content/uploads/2023/04/creamy-tomato-pasta-3.jpg",
	}
	itemsByCategory := map[string][]item.Item{
			// Restaurants
			"Appetizers": {
				{NameEn: "Mozzarella Sticks", NameAr: "رقائق جبنة الموزاريلا", DescriptionEn: "Crispy fried mozzarella sticks with marinara sauce", DescriptionAr: "رقائق جبنة الموزاريلا المقلية مع صلصة المارينارا", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Chicken Wings", NameAr: "أجنحة دجاج", DescriptionEn: "Spicy buffalo wings with blue cheese dip", DescriptionAr: "أجنحة دجاج حارة مع صلصة الجبن الأزرق", BasePrice: 45, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Nachos", NameAr: "ناتشوز", DescriptionEn: "Tortilla chips with cheese, jalapeños, and guacamole", DescriptionAr: "رقائق الذرة مع الجبن والجلابينو والأفوكادو", BasePrice: 30, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Main Courses": {
				{NameEn: "Grilled Salmon", NameAr: "سمك السلمون المشوي", DescriptionEn: "Fresh Atlantic salmon with lemon butter sauce", DescriptionAr: "سمك السلمون الطازج مع صلصة الزبدة والليمون", BasePrice: 120, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Beef Steak", NameAr: "شريحة لحم بقري", DescriptionEn: "8oz ribeye steak with mashed potatoes", DescriptionAr: "شريحة لحم بقري مع البطاطس المهروسة", BasePrice: 150, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Chicken Parmesan", NameAr: "دجاج بارميزان", DescriptionEn: "Breaded chicken breast with tomato sauce and mozzarella", DescriptionAr: "صدر دجاج مغطى بالفتات مع صلصة الطماطم والجبن", BasePrice: 85, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Burgers": {
				{NameEn: "Classic Cheeseburger", NameAr: "برجر الجبن الكلاسيكي", DescriptionEn: "Beef patty with cheddar cheese, lettuce, and tomato", DescriptionAr: "باتي لحم بقري مع جبن تشيدر والخس والطماطم", BasePrice: 55, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "BBQ Bacon Burger", NameAr: "برجر باربيكيو بيكون", DescriptionEn: "Beef burger with crispy bacon and BBQ sauce", DescriptionAr: "برجر لحم بقري مع بيكون مقرمش وصلصة الباربيكيو", BasePrice: 65, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Mushroom Swiss Burger", NameAr: "برجر فطر سويسري", DescriptionEn: "Beef patty with sautéed mushrooms and Swiss cheese", DescriptionAr: "باتي لحم بقري مع الفطر المقلي وجبن السويسري", BasePrice: 60, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Pizzas": {
				{NameEn: "Margherita Pizza", NameAr: "بيتزا مارغريتا", DescriptionEn: "Classic pizza with tomato sauce and fresh mozzarella", DescriptionAr: "بيتزا كلاسيكية مع صلصة الطماطم وجبن الموزاريلا الطازج", BasePrice: 70, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Pepperoni Pizza", NameAr: "بيتزا بيبروني", DescriptionEn: "Pizza topped with pepperoni slices and mozzarella cheese", DescriptionAr: "بيتزا مغطاة بشرائح البيبروني وجبن الموزاريلا", BasePrice: 85, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Vegetarian Pizza", NameAr: "بيتزا خضار", DescriptionEn: "Loaded with bell peppers, mushrooms, olives, and onions", DescriptionAr: "محملة بفلفل حلو والفطر والزيتون والبصل", BasePrice: 75, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Seafood Pizza", NameAr: "بيتزا مأكولات بحرية", DescriptionEn: "Shrimp and calamari with garlic sauce", DescriptionAr: "روبيان وحبار مع صلصة الثوم", BasePrice: 95, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Pasta": {
				{NameEn: "Spaghetti Carbonara", NameAr: "معكرونة كاربونارا", DescriptionEn: "Classic Italian pasta with eggs, cheese, and pancetta", DescriptionAr: "معكرونة إيطالية كلاسيكية مع البيض والجبن والبانسيتا", BasePrice: 65, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Fettuccine Alfredo", NameAr: "معكرونة الفيتوتشيني ألفريدو", DescriptionEn: "Creamy pasta with parmesan cheese and butter", DescriptionAr: "معكرونة كريمية مع جبن البارميزان والزبدة", BasePrice: 70, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Penne Arrabbiata", NameAr: "معكرونة بيني أرابياتا", DescriptionEn: "Penne pasta in spicy tomato sauce with garlic", DescriptionAr: "معكرونة بيني في صلصة طماطم حارة مع الثوم", BasePrice: 55, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Seafood": {
				{NameEn: "Fish & Chips", NameAr: "سمك ورقائق", DescriptionEn: "Beer-battered fish with french fries", DescriptionAr: "سمك مغطى بعجين البيرة مع شرائح البطاطس", BasePrice: 75, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Shrimp Scampi", NameAr: "جمبري سكامبي", DescriptionEn: "Shrimp in garlic butter sauce with pasta", DescriptionAr: "جمبري في صلصة زبدة الثوم مع المعكرونة", BasePrice: 85, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Desserts": {
				{NameEn: "Tiramisu", NameAr: "تيراميسو", DescriptionEn: "Classic Italian coffee-flavored dessert", DescriptionAr: "حلوى إيطالية كلاسيكية بنكهة القهوة", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Chocolate Lava Cake", NameAr: "كيك شوكولاتة ساخن", DescriptionEn: "Warm chocolate cake with molten center", DescriptionAr: "كيك شوكولاتة دافئ مع مركز سائل", BasePrice: 30, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Cheesecake", NameAr: "كيك الجبن", DescriptionEn: "New York style cheesecake with berry compote", DescriptionAr: "كيك جبن بأسلوب نيويورك مع مربى التوت", BasePrice: 32, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Drinks": {
				{NameEn: "Fresh Lemonade", NameAr: "ليمونادة طازجة", DescriptionEn: "Freshly squeezed lemons with mint", DescriptionAr: "ليمون معصور طازج مع النعناع", BasePrice: 15, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Iced Coffee", NameAr: "قهوة مثلجة", DescriptionEn: "Cold brew coffee with milk and ice", DescriptionAr: "قهوة باردة مع الحليب والثلج", BasePrice: 20, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Mango Smoothie", NameAr: "سموذي المانجو", DescriptionEn: "Fresh mango blended with yogurt", DescriptionAr: "مانجو طازج ممزوج مع الزبادي", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Craft Beer", NameAr: "بيرة حرفية", DescriptionEn: "Local craft beer selection", DescriptionAr: "اختيار من البيرة الحرفية المحلية", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Breakfast": {
				{NameEn: "Avocado Toast", NameAr: "خبز الأفوكادو", DescriptionEn: "Sourdough bread with smashed avocado and poached egg", DescriptionAr: "خبز السوردوغ مع الأفوكادو المهروس وبيض مسلوق", BasePrice: 30, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Pancakes", NameAr: "فطائر", DescriptionEn: "Fluffy pancakes with maple syrup and berries", DescriptionAr: "فطائر ناعمة مع شراب القيقب والتوت", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Salads": {
				{NameEn: "Caesar Salad", NameAr: "سلطة سيزر", DescriptionEn: "Romaine lettuce with Caesar dressing and croutons", DescriptionAr: "خس رومين مع صلصة سيزر والكروتون", BasePrice: 40, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Greek Salad", NameAr: "سلطة يونانية", DescriptionEn: "Tomatoes, cucumbers, olives, and feta cheese", DescriptionAr: "طماطم وخيار وزيتون وجبن الفيتا", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Sandwiches": {
				{NameEn: "Club Sandwich", NameAr: "ساندوتش كلوب", DescriptionEn: "Triple-layer sandwich with turkey, bacon, and lettuce", DescriptionAr: "ساندوتش ثلاثي الطبقات مع تركيا وبيكون وخس", BasePrice: 45, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Chicken Shawarma", NameAr: "شاورما دجاج", DescriptionEn: "Marinated chicken with garlic sauce in pita bread", DescriptionAr: "دجاج متبل مع صلصة الثوم في خبز البيتا", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Kids Menu": {
				{NameEn: "Chicken Nuggets", NameAr: "قطع دجاج مقلي", DescriptionEn: "6 pieces of breaded chicken nuggets with fries", DescriptionAr: "6 قطع من قطع الدجاج المغطاة بالفتات مع البطاطس", BasePrice: 30, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Mini Burger", NameAr: "برجر صغير", DescriptionEn: "Small beef burger with cheese and fries", DescriptionAr: "برجر لحم بقري صغير مع الجبن والبطاطس", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},

			// Supermarkets
			"Fruits": {
				{NameEn: "Organic Apples", NameAr: "تفاح عضوي", DescriptionEn: "Fresh organic red apples from local farms", DescriptionAr: "تفاح أحمر عضوي طازج من المزارع المحلية", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Bananas", NameAr: "موز", DescriptionEn: "Fresh yellow bananas", DescriptionAr: "موز أصفر طازج", BasePrice: 15, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Oranges", NameAr: "برتقال", DescriptionEn: "Juicy navel oranges", DescriptionAr: "برتقال حلو وعصيري", BasePrice: 20, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Grapes", NameAr: "عنب", DescriptionEn: "Sweet red grapes", DescriptionAr: "عنب أحمر حلو", BasePrice: 30, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Strawberries", NameAr: "فراولة", DescriptionEn: "Fresh strawberries", DescriptionAr: "فراولة طازجة", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Vegetables": {
				{NameEn: "Tomatoes", NameAr: "طماطم", DescriptionEn: "Fresh vine-ripened tomatoes", DescriptionAr: "طماطم طازجة من الكرمة", BasePrice: 12, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Cucumbers", NameAr: "خيار", DescriptionEn: "Crisp fresh cucumbers", DescriptionAr: "خيار طازج ومقرمش", BasePrice: 8, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Carrots", NameAr: "جزر", DescriptionEn: "Fresh orange carrots", DescriptionAr: "جزر برتقالي طازج", BasePrice: 10, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Bell Peppers", NameAr: "فلفل حلو", DescriptionEn: "Red, yellow, and green bell peppers", DescriptionAr: "فلفل أحمر وأصفر وأخضر حلو", BasePrice: 18, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Onions", NameAr: "بصل", DescriptionEn: "Yellow onions", DescriptionAr: "بصل أصفر", BasePrice: 6, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Meat & Poultry": {
				{NameEn: "Chicken Breast", NameAr: "صدر دجاج", DescriptionEn: "Boneless skinless chicken breasts", DescriptionAr: "صدور دجاج بدون عظام وجلد", BasePrice: 45, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Ground Beef", NameAr: "لحم بقري مفروم", DescriptionEn: "80% lean ground beef", DescriptionAr: "لحم بقري مفروم 80% خالي من الدهن", BasePrice: 55, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Salmon Fillet", NameAr: "فيليه سلمون", DescriptionEn: "Fresh Atlantic salmon fillet", DescriptionAr: "فيليه سلمون أطلسي طازج", BasePrice: 120, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Dairy": {
				{NameEn: "Whole Milk", NameAr: "حليب كامل الدسم", DescriptionEn: "1 liter fresh whole milk", DescriptionAr: "1 لتر حليب طازج كامل الدسم", BasePrice: 15, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Cheddar Cheese", NameAr: "جبن تشيدر", DescriptionEn: "200g block of aged cheddar cheese", DescriptionAr: "قطعة 200 جرام من جبن التشيدر المعمر", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Greek Yogurt", NameAr: "زبادي يوناني", DescriptionEn: "500g plain Greek yogurt", DescriptionAr: "500 جرام زبادي يوناني عادي", BasePrice: 18, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Butter", NameAr: "زبدة", DescriptionEn: "250g unsalted butter", DescriptionAr: "250 جرام زبدة غير مملحة", BasePrice: 20, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Bakery": {
				{NameEn: "Baguette", NameAr: "خبز باكيت", DescriptionEn: "Freshly baked French baguette", DescriptionAr: "خبز باكيت فرنسي طازج", BasePrice: 8, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Croissants", NameAr: "كرواسون", DescriptionEn: "Buttery flaky croissants", DescriptionAr: "كرواسون زبدة مقرمش", BasePrice: 12, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Whole Wheat Bread", NameAr: "خبز القمح الكامل", DescriptionEn: "1 loaf of whole wheat bread", DescriptionAr: "1 رغيف خبز قمح كامل", BasePrice: 10, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Beverages": {
				{NameEn: "Mineral Water", NameAr: "مياه معدنية", DescriptionEn: "500ml natural mineral water", DescriptionAr: "500 مل مياه معدنية طبيعية", BasePrice: 5, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Orange Juice", NameAr: "عصير برتقال", DescriptionEn: "1 liter fresh orange juice", DescriptionAr: "1 لتر عصير برتقال طازج", BasePrice: 20, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Green Tea", NameAr: "شاي أخضر", DescriptionEn: "Box of 20 green tea bags", DescriptionAr: "علبة تحتوي على 20 كيس شاي أخضر", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Coffee Beans", NameAr: "حبوب قهوة", DescriptionEn: "250g premium Arabica coffee beans", DescriptionAr: "250 جرام حبوب قهوة عربية ممتازة", BasePrice: 45, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Snacks": {
				{NameEn: "Potato Chips", NameAr: "رقائق بطاطس", DescriptionEn: "Classic salted potato chips", DescriptionAr: "رقائق بطاطس مملحة كلاسيكية", BasePrice: 7, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Chocolate Bar", NameAr: "شوكولاتة", DescriptionEn: "Milk chocolate bar", DescriptionAr: "شريط شوكولاتة حليب", BasePrice: 5, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Mixed Nuts", NameAr: "مكسرات مختلطة", DescriptionEn: "200g mixed nuts (almonds, walnuts, cashews)", DescriptionAr: "200 جرام مكسرات مختلطة (لوز وجوز وكاجو)", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Frozen Foods": {
				{NameEn: "Frozen Vegetables", NameAr: "خضروات مجمدة", DescriptionEn: "Mixed frozen vegetables", DescriptionAr: "خضروات مجمدة مختلطة", BasePrice: 15, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Ice Cream", NameAr: "آيس كريم", DescriptionEn: "Vanilla ice cream tub", DescriptionAr: "علبة آيس كريم فانيليا", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Canned Goods": {
				{NameEn: "Canned Tuna", NameAr: "تونة معلبة", DescriptionEn: "Chunk light tuna in water", DescriptionAr: "قطع تونة خفيفة في الماء", BasePrice: 12, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Tomato Soup", NameAr: "شوربة طماطم", DescriptionEn: "Condensed tomato soup", DescriptionAr: "شوربة طماطم مكثفة", BasePrice: 8, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Cleaning Supplies": {
				{NameEn: "Laundry Detergent", NameAr: "مسحوق غسيل", DescriptionEn: "3L concentrated laundry detergent", DescriptionAr: "3 لتر مسحوق غسيل مركز", BasePrice: 35, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Dish Soap", NameAr: "صابون أطباق", DescriptionEn: "Liquid dish soap with lemon scent", DescriptionAr: "صابون أطباق سائل برائحة الليمون", BasePrice: 15, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "All-Purpose Cleaner", NameAr: "منظف متعدد الأغراض", DescriptionEn: "500ml all-purpose surface cleaner", DescriptionAr: "500 مل منظف سطوح متعدد الأغراض", BasePrice: 20, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Personal Care": {
				{NameEn: "Shampoo", NameAr: "شامبو", DescriptionEn: "500ml anti-dandruff shampoo", DescriptionAr: "500 مل شامبو مضاد للقشرة", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Toothpaste", NameAr: "معجون أسنان", DescriptionEn: "Fresh mint toothpaste", DescriptionAr: "معجون أسنان بنكهة النعناع الطازج", BasePrice: 12, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
			"Baby Products": {
				{NameEn: "Baby Diapers", NameAr: "حفاضات أطفال", DescriptionEn: "Pack of 50 size 2 diapers", DescriptionAr: "علبة 50 حفاضة مقاس 2", BasePrice: 80, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
				{NameEn: "Baby Wipes", NameAr: "مناديل أطفال", DescriptionEn: "70 wipes with aloe vera", DescriptionAr: "70 منديل مع ألوة فيرا", BasePrice: 25, Picture: pictureURLs[rand.Intn(len(pictureURLs))]},
			},
		}

	// ----------------------
	// SEED ITEMS TO VENDORS
	// ----------------------
	var vendorCategories []category.VendorCategory
	db.Find(&vendorCategories)

	for _, v := range vendors {
		// detect vendor category
		var vc category.VendorCategory
		if err := db.First(&vc, v.CategoryID).Error; err != nil {
			continue
		}

		// pick categories depending on vendor type
		var cats []itemcategory.ItemsCategory
		if vc.NameEn == "Restaurants" {
			db.Where("name_en IN ?", []string{
				"Appetizers", "Main Courses", "Burgers", "Pizzas", "Pasta", "Seafood",
				"Desserts", "Drinks", "Breakfast", "Salads", "Sandwiches", "Kids Menu",
			}).Find(&cats)
		} else if vc.NameEn == "Supermarkets" {
			db.Where("name_en IN ?", []string{
				"Fruits", "Vegetables", "Meat & Poultry", "Seafood", "Dairy", "Bakery",
				"Beverages", "Snacks", "Frozen Foods", "Canned Goods", "Cleaning Supplies",
				"Personal Care", "Baby Products",
			}).Find(&cats)
		}

		for _, cat := range cats {
			for i, it := range itemsByCategory[cat.NameEn] {
				// Create a new item instance to avoid reference issues
				newItem := item.Item{
					NameEn:        it.NameEn,
					NameAr:        it.NameAr,
					DescriptionEn: it.DescriptionEn,
					DescriptionAr: it.DescriptionAr,
					BasePrice:     it.BasePrice,
					Picture:       it.Picture,
					CategoryID:    int64(cat.ID),
					VendorID:      int64(v.ID),
					DiscountPercent: 0, // default
					Stock:         100,
					IsActive:      true,
				}

				// Apply discount to some items (e.g., every even index)
				if i%2 == 0 {
					newItem.DiscountPercent = 10
				} else if i%3 == 0 {
					newItem.DiscountPercent = 20
				}

				var existing item.Item
				if err := db.Where("name_en = ? AND vendor_id = ?", newItem.NameEn, newItem.VendorID).First(&existing).Error; err == gorm.ErrRecordNotFound {
					if err := db.Create(&newItem).Error; err != nil {
						log.Printf("❌ Failed to insert item %s for vendor %s: %v", newItem.NameEn, v.NameEn, err)
						continue
					}

					db.First(&newItem, "name_en = ? AND vendor_id = ?", newItem.NameEn, newItem.VendorID)

					// Vary sizes based on category
					var sizes []item.ItemSize
					if cat.NameEn == "Burgers" {
						sizes = []item.ItemSize{
							{Name: "Single", Price: newItem.BasePrice - 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Double", Price: newItem.BasePrice, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Triple", Price: newItem.BasePrice + 20, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else if cat.NameEn == "Pizzas" {
						sizes = []item.ItemSize{
							{Name: "Small (8\")", Price: newItem.BasePrice - 20, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Medium (12\")", Price: newItem.BasePrice, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Large (16\")", Price: newItem.BasePrice + 30, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Extra Large (20\")", Price: newItem.BasePrice + 50, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else if cat.NameEn == "Drinks" || cat.NameEn == "Beverages" {
						sizes = []item.ItemSize{
							{Name: "Small (250ml)", Price: newItem.BasePrice - 5, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Medium (500ml)", Price: newItem.BasePrice, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Large (1L)", Price: newItem.BasePrice + 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else if cat.NameEn == "Fruits" || cat.NameEn == "Vegetables" {
						sizes = []item.ItemSize{
							{Name: "500g", Price: newItem.BasePrice - 5, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "1kg", Price: newItem.BasePrice, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "2kg", Price: newItem.BasePrice + 15, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else {
						sizes = []item.ItemSize{
							{Name: "Small", Price: newItem.BasePrice - 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Medium", Price: newItem.BasePrice, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{Name: "Large", Price: newItem.BasePrice + 15, ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					}
					db.Create(&sizes)

					// Vary variants based on category
					var variants []item.ItemVariant
					if cat.NameEn == "Burgers" || cat.NameEn == "Pizzas" {
						variants = []item.ItemVariant{
							{NameEn: "Spicy", NameAr: "حار", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Vegetarian", NameAr: "نباتي", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Gluten-Free", NameAr: "خالي من الغلوتين", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else if cat.NameEn == "Seafood" {
						variants = []item.ItemVariant{
							{NameEn: "Grilled", NameAr: "مشوي", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Fried", NameAr: "مقلي", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Steamed", NameAr: "مبخر", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else if cat.NameEn == "Drinks" || cat.NameEn == "Beverages" {
						variants = []item.ItemVariant{
							{NameEn: "Sweetened", NameAr: "محلى", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Unsweetened", NameAr: "غير محلى", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Decaf", NameAr: "خالي من الكافيين", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else if cat.NameEn == "Fruits" || cat.NameEn == "Vegetables" {
						variants = []item.ItemVariant{
							{NameEn: "Organic", NameAr: "عضوي", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Conventional", NameAr: "تقليدي", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					} else {
						variants = []item.ItemVariant{
							{NameEn: "Spicy", NameAr: "حار", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Extra", NameAr: "إضافي", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
							{NameEn: "Lite", NameAr: "خفيف", ItemID: int64(newItem.ID), VendorID: newItem.VendorID},
						}
					}
					db.Create(&variants)

					// Vary addons based on category, including removals
					var addons []item.ItemAddon
					if cat.NameEn == "Burgers" {
						addons = []item.ItemAddon{
							{NameEn: "Extra Patty", NameAr: "باتي إضافي", Price: 20, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "Bacon Strips", NameAr: "شرائح بيكون", Price: 15, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "No Pickles", NameAr: "بدون مخلل", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: true},
							{NameEn: "Avocado", NameAr: "أفوكادو", Price: 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "No Onion", NameAr: "بدون بصل", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: true},
						}
					} else if cat.NameEn == "Pizzas" {
						addons = []item.ItemAddon{
							{NameEn: "Extra Cheese", NameAr: "جبن إضافي", Price: 15, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "Pepperoni", NameAr: "بيبروني", Price: 20, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "No Olives", NameAr: "بدون زيتون", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: true},
							{NameEn: "Mushrooms", NameAr: "فطر", Price: 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
						}
					} else if cat.NameEn == "Seafood" {
						addons = []item.ItemAddon{
							{NameEn: "Lemon Wedge", NameAr: "شريحة ليمون", Price: 5, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "Tartar Sauce", NameAr: "صلصة التارتار", Price: 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "No Garlic", NameAr: "بدون ثوم", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: true},
						}
					} else if cat.NameEn == "Drinks" || cat.NameEn == "Beverages" {
						addons = []item.ItemAddon{
							{NameEn: "Extra Ice", NameAr: "ثلج إضافي", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "Lemon Slice", NameAr: "شريحة ليمون", Price: 5, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "No Sugar", NameAr: "بدون سكر", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: true},
						}
					} else if cat.NameEn == "Fruits" || cat.NameEn == "Vegetables" {
						addons = []item.ItemAddon{
							{NameEn: "Pre-Washed", NameAr: "مغسول مسبقاً", Price: 5, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "Sliced", NameAr: "مقطع", Price: 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
						}
					} else {
						addons = []item.ItemAddon{
							{NameEn: "Extra Cheese", NameAr: "جبن إضافي", Price: 10, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "Extra Sauce", NameAr: "صلصة إضافية", Price: 5, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
							{NameEn: "No Onion", NameAr: "بدون بصل", Price: 0, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: true},
							{NameEn: "Extra Toppings", NameAr: "إضافات إضافية", Price: 15, ItemID: int64(newItem.ID), VendorID: newItem.VendorID, IsRemoval: false},
						}
					}
					db.Create(&addons)

					log.Printf("🍴 Inserted %s under %s for vendor %s", newItem.NameEn, cat.NameEn, v.NameEn)
				}
			}
		}
	}

	log.Println("✅ Seeded items for restaurants & supermarkets successfully.")
	return nil
}