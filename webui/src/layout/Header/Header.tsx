import React from "react";
import styles from "./Header.module.scss";
import { ReactComponent as BacalhauLogo} from "../../images/bacalhau.svg";
import profile from "../../images/profile.png";

interface HeaderProps {
  collapsed: boolean;
  pageTitle: string;
}

const Header: React.FC<HeaderProps> = ({ pageTitle, collapsed }) => {
  return (
    <header className={styles.header} data-collapsed={collapsed}>
      <div className={styles.left}>
        <BacalhauLogo className={styles.logo} height="24"/>
        <div className={styles.pageTitle}>{pageTitle}</div>
        <div className={styles.searchBar}>
          {/* Placeholder for search bar */}
          {/* <input type="text" placeholder="Search..." /> */}
        </div>
      </div>
      <div className={styles.right}>
        {/* Profile section */}
        <div className={styles.profile}>
          <img src={profile} alt="Logo" width={150} height={150} />
        </div>
      </div>
    </header>
  );
};

export default Header;
