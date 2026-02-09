import type { App } from 'vue'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import {
  faAddressCard as farAddressCard,
  faEye as farEye,
  faEyeSlash as farEyeSlash,
  faSquare as farSquare,
} from '@fortawesome/free-regular-svg-icons'
import {
  faAddressBook as fasAddressBook,
  faArrowRight as fasArrowRight,
  faArrowRightFromBracket as fasArrowRightFromBracket,
  faBars as fasBars,
  faChevronLeft as fasChevronLeft,
  faChevronRight as fasChevronRight,
  faCircleMinus as fasCircleMinus,
  faCirclePlus as fasCirclePlus,
  faDesktop as fasDesktop,
  faHouse as fasHouse,
  faSquareCheck as fasSquareCheck,
  faRotate as fasRotate,
  faUserPlus as fasUserPlus,
} from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(
  farAddressCard,
  farEye,
  farEyeSlash,
  farSquare,
  fasAddressBook,
  fasArrowRight,
  fasArrowRightFromBracket,
  fasBars,
  fasChevronLeft,
  fasChevronRight,
  fasCircleMinus,
  fasCirclePlus,
  fasDesktop,
  fasHouse,
  fasSquareCheck,
  fasRotate,
  fasUserPlus
)

export function registerVueComponent(app: App<Element>) {
  app.component('FontAwesomeIcon', FontAwesomeIcon)
}
