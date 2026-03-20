**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-20 16:38:02 UTC（2026-03-21 00:38:02 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 20.27.15.111:8561 | ✓ 1324ms | ✓ 1076ms | ✓ 677ms | ✓ 850ms | ✓ 686ms | http |
| 20.27.13.35:8561 | ✓ 1325ms | ✓ 1327ms | ✓ 499ms | ✓ 829ms | ✓ 650ms | http |
| 20.78.26.206:8561 | ✓ 1325ms | ✓ 814ms | ✓ 560ms | ✓ 1233ms | ✓ 824ms | http |
| 178.156.187.185:10001 | ✓ 906ms | 否 | ✓ 910ms | ✓ 1478ms | ✓ 1544ms | http |
| 147.161.210.140:8800 | ✓ 1325ms | 否 | ✓ 967ms | ✓ 957ms | ✓ 1085ms | http |
| 113.160.132.26:8080 | ✓ 1493ms | 否 | ✓ 1215ms | ✓ 1324ms | ✓ 1044ms | http |
| 174.138.24.77:1080 | ✓ 774ms | 否 | 否 | ✓ 1632ms | ✓ 1370ms | http |
| 167.103.34.108:8800 | ✓ 1433ms | 否 | ✓ 1700ms | ✓ 1840ms | ✓ 1606ms | http |
| 46.101.190.71:3128 | 否 | 否 | ✓ 1830ms | ✓ 1998ms | ✓ 1571ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 810ms | ✓ 1140ms | ✓ 1615ms | http |
| 20.27.14.220:8561 | ✓ 498ms | ✓ 777ms | ✓ 552ms | ✓ 901ms | ✓ 681ms | http |
| 20.210.39.153:8561 | ✓ 550ms | ✓ 1040ms | ✓ 488ms | ✓ 800ms | ✓ 625ms | http |
| 20.27.11.248:8561 | ✓ 544ms | ✓ 1532ms | ✓ 518ms | ✓ 884ms | ✓ 644ms | http |
| 20.78.118.91:8561 | ✓ 1531ms | ✓ 812ms | ✓ 488ms | ✓ 801ms | ✓ 636ms | http |
| 38.145.220.11:8445 | 否 | ✓ 1650ms | ✓ 209ms | ✓ 742ms | ✓ 568ms | http |
| 120.92.212.16:7890 | ✓ 1055ms | ✓ 1236ms | ✓ 969ms | ✓ 1261ms | ✓ 999ms | http |
| 38.34.183.234:8450 | 否 | 否 | ✓ 333ms | ✓ 871ms | ✓ 1201ms | http |
| 137.220.150.22:6005 | ✓ 1885ms | 否 | ✓ 1855ms | ✓ 1369ms | ✓ 1250ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1237ms | ✓ 1472ms | ✓ 1249ms | http |
| 45.167.124.52:8080 | ✓ 1737ms | 否 | ✓ 646ms | ✓ 1632ms | ✓ 1385ms | http |
| 116.80.65.79:3172 | 否 | 否 | ✓ 1526ms | ✓ 1841ms | ✓ 1682ms | http |
| 186.148.180.46:999 | ✓ 1033ms | 否 | ✓ 1679ms | ✓ 1885ms | ✓ 1435ms | http |
| 157.245.194.13:8888 | ✓ 740ms | 否 | ✓ 1054ms | ✓ 1067ms | ✓ 865ms | http |
| 219.117.204.211:7799 | ✓ 1752ms | 否 | ✓ 1065ms | 否 | ✓ 1969ms | http |
| 167.103.31.122:8800 | ✓ 1386ms | 否 | ✓ 1557ms | 否 | ✓ 1595ms | http |
| 38.145.208.244:8448 | ✓ 1683ms | ✓ 1451ms | ✓ 1488ms | ✓ 1215ms | ✓ 618ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1575ms | ✓ 1526ms | ✓ 843ms | http |
| 38.34.179.162:8451 | ✓ 1663ms | ✓ 1522ms | ✓ 1460ms | ✓ 1766ms | ✓ 1551ms | http |
| 38.34.183.233:8448 | ✓ 1671ms | ✓ 1228ms | 否 | ✓ 1403ms | ✓ 1088ms | http |
| 137.220.150.152:6005 | ✓ 1883ms | 否 | ✓ 759ms | ✓ 1277ms | 否 | http |
| 38.34.179.150:8449 | ✓ 1674ms | 否 | ✓ 1641ms | ✓ 1959ms | 否 | http |
| 38.34.179.16:8451 | ✓ 780ms | 否 | ✓ 460ms | ✓ 969ms | ✓ 633ms | http |
| 154.64.243.50:7890 | ✓ 777ms | 否 | ✓ 882ms | 否 | ✓ 673ms | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 648ms | ✓ 1686ms | ✓ 1138ms | http |
| 147.161.239.240:8800 | ✓ 1257ms | 否 | ✓ 1704ms | 否 | ✓ 1578ms | http |
| 91.238.105.64:2024 | ✓ 1284ms | 否 | ✓ 1598ms | 否 | ✓ 1970ms | http |
| 101.43.127.100:8877 | ✓ 1742ms | ✓ 1769ms | 否 | ✓ 1769ms | ✓ 939ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 1046ms | ✓ 998ms | ✓ 1607ms | http |
| 35.225.22.61:80 | ✓ 711ms | 否 | ✓ 352ms | ✓ 1066ms | ✓ 1117ms | http |
| 45.136.131.60:8452 | ✓ 484ms | 否 | ✓ 820ms | ✓ 1323ms | ✓ 1310ms | http |
| 162.240.154.26:3128 | ✓ 1257ms | 否 | ✓ 488ms | ✓ 1014ms | ✓ 1125ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1356ms | 否 | ✓ 1640ms | ✓ 1317ms | http |
| 38.34.179.14:8450 | ✓ 1015ms | ✓ 1718ms | ✓ 1637ms | 否 | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1019ms | 否 | ✓ 1062ms | ✓ 868ms | http |
| 133.242.138.34:8100 | 否 | 否 | ✓ 1468ms | ✓ 1004ms | ✓ 813ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1045ms | ✓ 1565ms | ✓ 961ms | http |
| 194.5.212.40:8080 | ✓ 1734ms | 否 | ✓ 1133ms | ✓ 1704ms | ✓ 1631ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1153ms | ✓ 1486ms | ✓ 1212ms | http |
| 101.32.244.83:8080 | ✓ 1496ms | ✓ 1651ms | ✓ 1131ms | ✓ 1200ms | ✓ 1225ms | http |
| 121.43.196.210:8222 | ✓ 955ms | ✓ 1109ms | ✓ 904ms | ✓ 1117ms | ✓ 904ms | http |
| 121.43.196.213:8222 | ✓ 954ms | ✓ 1109ms | ✓ 902ms | ✓ 1144ms | ✓ 904ms | http |
| 114.55.226.123:10086 | ✓ 1037ms | 否 | ✓ 1048ms | ✓ 1290ms | ✓ 1039ms | http |
| 114.237.77.245:1080 | ✓ 894ms | ✓ 1789ms | ✓ 944ms | ✓ 1481ms | ✓ 929ms | http |
| 38.55.104.182:6005 | ✓ 861ms | 否 | ✓ 933ms | ✓ 1223ms | ✓ 693ms | http |
| 45.136.131.38:8449 | 否 | 否 | ✓ 1848ms | ✓ 820ms | ✓ 1012ms | http |
| 38.145.203.97:8448 | 否 | ✓ 1291ms | 否 | ✓ 1223ms | ✓ 1516ms | http |
| 38.145.203.106:8448 | 否 | ✓ 1453ms | 否 | ✓ 1124ms | ✓ 1724ms | http |
| 8.219.97.248:80 | ✓ 1256ms | 否 | ✓ 1353ms | 否 | ✓ 1318ms | http |
| 116.80.65.80:3172 | ✓ 508ms | 否 | ✓ 919ms | ✓ 1642ms | ✓ 815ms | http |
| 139.159.99.242:8080 | 否 | ✓ 1052ms | ✓ 866ms | ✓ 1097ms | ✓ 846ms | http |
| 38.34.178.154:8445 | 否 | 否 | ✓ 647ms | ✓ 766ms | ✓ 582ms | http |
| 111.79.111.126:3128 | ✓ 1052ms | ✓ 1381ms | ✓ 1096ms | 否 | 否 | http |
| 137.220.151.110:6005 | ✓ 1520ms | 否 | ✓ 781ms | ✓ 1151ms | ✓ 866ms | http |
| 142.171.224.229:7890 | ✓ 1569ms | 否 | ✓ 804ms | ✓ 768ms | ✓ 590ms | http |
| 47.77.193.180:1080 | ✓ 173ms | ✓ 1313ms | ✓ 493ms | ✓ 759ms | ✓ 561ms | http |
| 101.47.73.135:3128 | ✓ 1309ms | 否 | ✓ 1652ms | ✓ 1310ms | 否 | http |
| 115.231.181.40:8128 | ✓ 945ms | ✓ 1963ms | ✓ 1962ms | 否 | 否 | http |
| 45.129.141.143:3128 | 否 | 否 | ✓ 1869ms | ✓ 1877ms | ✓ 1633ms | http |
| 38.145.208.227:8451 | ✓ 445ms | ✓ 1269ms | ✓ 790ms | ✓ 776ms | ✓ 611ms | http |
| 38.145.203.34:8451 | ✓ 702ms | 否 | ✓ 149ms | ✓ 774ms | ✓ 853ms | http |
| 38.145.203.96:8452 | ✓ 433ms | ✓ 1136ms | ✓ 938ms | ✓ 1752ms | ✓ 630ms | http |
| 38.145.208.160:8447 | ✓ 1441ms | ✓ 1736ms | ✓ 423ms | ✓ 1544ms | ✓ 698ms | http |
| 8.222.175.80:6128 | ✓ 803ms | ✓ 1860ms | ✓ 860ms | ✓ 1055ms | ✓ 863ms | http |
| 38.34.179.61:8445 | 否 | ✓ 1431ms | ✓ 525ms | ✓ 982ms | 否 | http |
| 211.217.231.234:8080 | 否 | ✓ 1317ms | ✓ 985ms | ✓ 1017ms | ✓ 850ms | http |
| 150.249.255.91:3128 | ✓ 1395ms | 否 | ✓ 569ms | ✓ 842ms | ✓ 1720ms | http |
| 194.150.220.97:3128 | 否 | 否 | ✓ 868ms | ✓ 1590ms | ✓ 1189ms | http |
| 62.113.119.14:8080 | 否 | ✓ 1677ms | 否 | ✓ 1694ms | ✓ 1190ms | http |
| 38.34.179.60:8450 | 否 | ✓ 1849ms | ✓ 758ms | ✓ 1813ms | ✓ 1381ms | http |
| 62.234.206.73:3128 | 否 | 否 | ✓ 986ms | ✓ 1820ms | ✓ 1039ms | http |
| 185.191.236.162:3128 | ✓ 1337ms | 否 | ✓ 1490ms | ✓ 1735ms | ✓ 1181ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1507ms | ✓ 1384ms | 否 | ✓ 1517ms | http |
| 38.34.179.96:8451 | ✓ 697ms | ✓ 1710ms | ✓ 196ms | ✓ 752ms | ✓ 807ms | http |
| 8.212.172.106:8080 | ✓ 1947ms | 否 | 否 | ✓ 1349ms | ✓ 1044ms | http |
| 152.69.229.220:3128 | 否 | 否 | ✓ 1455ms | ✓ 1702ms | ✓ 1738ms | http |
| 206.81.27.105:3128 | ✓ 1801ms | 否 | ✓ 1543ms | ✓ 1783ms | ✓ 1399ms | http |
| 103.39.51.190:8080 | ✓ 1873ms | 否 | ✓ 1237ms | ✓ 1540ms | 否 | http |
| 38.145.208.166:8451 | ✓ 833ms | ✓ 1302ms | ✓ 920ms | ✓ 810ms | 否 | http |
| 106.75.15.167:7890 | ✓ 1348ms | ✓ 1410ms | ✓ 918ms | 否 | ✓ 1594ms | http |
| 106.117.208.101:7890 | ✓ 1352ms | 否 | ✓ 1451ms | 否 | ✓ 1086ms | http |
| 45.140.147.155:1081 | ✓ 1049ms | 否 | ✓ 1057ms | ✓ 1832ms | ✓ 1225ms | http |
| 134.209.153.66:3128 | ✓ 1446ms | 否 | ✓ 1871ms | ✓ 1774ms | 否 | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
