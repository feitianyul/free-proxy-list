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

最后更新：2026-03-06 10:37:08 UTC（2026-03-06 18:37:08 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 946ms | 否 | ✓ 795ms | ✓ 1205ms | ✓ 878ms | http |
| 130.36.36.29:443 | ✓ 949ms | 否 | ✓ 1820ms | ✓ 1171ms | ✓ 826ms | http |
| 125.128.12.144:3128 | ✓ 1624ms | 否 | ✓ 1952ms | ✓ 1212ms | ✓ 1060ms | http |
| 120.92.212.16:7890 | ✓ 1502ms | 否 | ✓ 1304ms | ✓ 1603ms | ✓ 1397ms | http |
| 14.56.107.244:3128 | ✓ 1031ms | 否 | ✓ 722ms | 否 | ✓ 927ms | http |
| 217.76.245.80:999 | ✓ 730ms | 否 | ✓ 1156ms | ✓ 1722ms | ✓ 1602ms | http |
| 121.128.121.54:3128 | 否 | ✓ 1252ms | ✓ 726ms | ✓ 1326ms | ✓ 926ms | http |
| 185.191.236.162:3128 | ✓ 993ms | 否 | ✓ 992ms | ✓ 1762ms | ✓ 1107ms | http |
| 91.107.175.112:10801 | ✓ 513ms | 否 | ✓ 1764ms | 否 | ✓ 1670ms | http |
| 35.225.22.61:80 | ✓ 402ms | ✓ 1227ms | ✓ 264ms | ✓ 1078ms | 否 | http |
| 168.235.110.63:3128 | ✓ 649ms | ✓ 1220ms | ✓ 821ms | ✓ 1263ms | ✓ 785ms | http |
| 91.193.240.157:9877 | ✓ 872ms | 否 | ✓ 1356ms | 否 | ✓ 1823ms | http |
| 81.70.169.194:80 | ✓ 1286ms | ✓ 1511ms | ✓ 1132ms | ✓ 1550ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1385ms | ✓ 1705ms | ✓ 1160ms | 否 | 否 | http |
| 154.37.208.132:30000 | ✓ 657ms | 否 | 否 | ✓ 1556ms | ✓ 1678ms | http |
| 14.56.177.44:3128 | ✓ 857ms | ✓ 1622ms | ✓ 868ms | ✓ 1204ms | ✓ 989ms | http |
| 178.236.245.59:3128 | ✓ 719ms | ✓ 1895ms | ✓ 554ms | ✓ 1662ms | ✓ 1280ms | http |
| 178.236.245.17:3128 | ✓ 702ms | 否 | ✓ 549ms | ✓ 1590ms | ✓ 1271ms | http |
| 1.231.81.166:3128 | ✓ 844ms | 否 | ✓ 1018ms | ✓ 1123ms | ✓ 844ms | http |
| 152.42.195.165:8888 | ✓ 861ms | 否 | ✓ 1016ms | ✓ 1341ms | ✓ 1039ms | http |
| 172.105.212.216:8888 | 否 | 否 | ✓ 1361ms | ✓ 1140ms | ✓ 1062ms | http |
| 106.14.203.63:3333 | ✓ 963ms | ✓ 1390ms | ✓ 1032ms | ✓ 1386ms | ✓ 999ms | http |
| 159.223.42.219:3128 | ✓ 854ms | 否 | ✓ 1334ms | ✓ 1287ms | ✓ 1994ms | http |
| 101.43.255.96:80 | ✓ 1272ms | ✓ 1467ms | ✓ 1480ms | 否 | 否 | http |
| 68.183.92.28:3128 | ✓ 1179ms | 否 | ✓ 1941ms | ✓ 1855ms | ✓ 1260ms | http |
| 103.215.36.88:16299 | ✓ 1448ms | 否 | ✓ 1467ms | ✓ 1721ms | ✓ 1320ms | http |
| 5.252.33.13:2025 | ✓ 1512ms | 否 | ✓ 1751ms | 否 | ✓ 1837ms | http |
| 38.183.146.57:8080 | ✓ 1593ms | 否 | ✓ 1852ms | ✓ 1812ms | ✓ 1610ms | http |
| 171.6.81.113:8080 | ✓ 1931ms | 否 | ✓ 1455ms | ✓ 1816ms | ✓ 1679ms | http |
| 42.115.72.27:2038 | ✓ 1595ms | 否 | ✓ 1672ms | ✓ 1951ms | ✓ 1741ms | http |
| 125.25.28.17:8080 | ✓ 1943ms | 否 | ✓ 1587ms | ✓ 1751ms | ✓ 1680ms | http |
| 1.4.165.203:8080 | ✓ 1990ms | 否 | ✓ 1932ms | ✓ 1733ms | ✓ 1557ms | http |
| 101.108.16.22:8080 | ✓ 1939ms | 否 | ✓ 1881ms | ✓ 1749ms | ✓ 1655ms | http |
| 125.27.155.78:8080 | ✓ 1925ms | 否 | 否 | ✓ 1798ms | ✓ 1765ms | http |
| 171.7.24.240:8080 | ✓ 1949ms | 否 | ✓ 1854ms | 否 | ✓ 1790ms | http |
| 120.92.212.16:8890 | ✓ 1202ms | 否 | ✓ 1355ms | 否 | ✓ 1765ms | http |
| 14.225.222.185:7890 | ✓ 1307ms | ✓ 1907ms | ✓ 1853ms | ✓ 1343ms | ✓ 1230ms | http |
| 212.175.29.184:8080 | ✓ 1409ms | 否 | ✓ 1669ms | 否 | ✓ 1521ms | http |
| 125.128.12.14:3128 | ✓ 1863ms | ✓ 1880ms | ✓ 1419ms | 否 | ✓ 1314ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1042ms | 否 | ✓ 1182ms | ✓ 957ms | http |
| 61.72.221.234:3128 | ✓ 1856ms | 否 | ✓ 1977ms | ✓ 1219ms | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1082ms | 否 | ✓ 1766ms | ✓ 1186ms | http |
| 161.97.115.10:3128 | ✓ 1020ms | 否 | ✓ 867ms | 否 | ✓ 1809ms | http |
| 121.230.9.168:1080 | 否 | ✓ 1606ms | ✓ 1232ms | ✓ 1710ms | ✓ 1260ms | http |
| 171.7.71.21:8080 | ✓ 1504ms | 否 | ✓ 1862ms | ✓ 1837ms | ✓ 1907ms | http |
| 124.122.241.10:8080 | ✓ 1675ms | 否 | 否 | ✓ 1761ms | ✓ 1711ms | http |
| 103.82.23.118:5253 | ✓ 1993ms | 否 | ✓ 1413ms | ✓ 1626ms | ✓ 1589ms | http |
| 116.80.82.229:3172 | ✓ 1764ms | 否 | 否 | ✓ 1933ms | ✓ 1754ms | http |
| 116.80.82.227:3172 | ✓ 1766ms | 否 | ✓ 1960ms | 否 | ✓ 1799ms | http |
| 104.129.203.244:10785 | ✓ 619ms | 否 | ✓ 1091ms | ✓ 1098ms | ✓ 750ms | http |
| 39.104.201.40:7890 | ✓ 1179ms | ✓ 1331ms | ✓ 1165ms | ✓ 1488ms | ✓ 1077ms | http |
| 192.166.82.55:1080 | ✓ 1685ms | 否 | ✓ 1681ms | ✓ 1649ms | ✓ 1552ms | http |
| 121.230.8.208:1080 | 否 | ✓ 1868ms | 否 | ✓ 1707ms | ✓ 1633ms | http |
| 101.47.73.135:3128 | ✓ 1883ms | 否 | 否 | ✓ 1593ms | ✓ 1110ms | http |
| 101.32.244.83:8080 | ✓ 1086ms | 否 | ✓ 1083ms | ✓ 1776ms | ✓ 1545ms | http |
| 121.43.196.210:8222 | ✓ 1124ms | ✓ 1306ms | ✓ 1091ms | ✓ 1444ms | ✓ 1130ms | http |
| 121.43.196.213:8222 | 否 | ✓ 1278ms | ✓ 1058ms | ✓ 1261ms | ✓ 1095ms | http |
| 114.55.226.123:10086 | ✓ 1188ms | ✓ 1635ms | ✓ 1233ms | ✓ 1417ms | ✓ 1199ms | http |
| 47.97.40.4:8222 | ✓ 1552ms | ✓ 1221ms | ✓ 1181ms | ✓ 1226ms | ✓ 1141ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1189ms | ✓ 1196ms | 否 | ✓ 1111ms | http |
| 172.212.68.37:3128 | ✓ 1644ms | 否 | ✓ 1308ms | 否 | ✓ 1095ms | http |
| 104.129.203.244:11465 | ✓ 838ms | 否 | ✓ 224ms | ✓ 927ms | ✓ 755ms | http |
| 45.136.198.40:3128 | ✓ 704ms | 否 | ✓ 1322ms | ✓ 1908ms | ✓ 1720ms | http |
| 107.174.80.186:3128 | ✓ 889ms | ✓ 1608ms | 否 | ✓ 1033ms | ✓ 797ms | http |
| 74.48.78.224:2080 | ✓ 696ms | ✓ 1478ms | ✓ 1062ms | 否 | 否 | http |
| 61.72.221.194:3128 | ✓ 1620ms | 否 | ✓ 1903ms | 否 | ✓ 1630ms | http |
| 45.140.147.155:1081 | ✓ 1407ms | ✓ 1821ms | ✓ 1592ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 792ms | 否 | 否 | ✓ 1300ms | ✓ 770ms | http |
| 103.215.36.88:15852 | 否 | ✓ 1623ms | ✓ 1431ms | ✓ 1589ms | 否 | http |
| 194.59.204.87:9080 | ✓ 987ms | 否 | ✓ 1434ms | 否 | ✓ 1558ms | http |
| 45.129.141.143:3128 | ✓ 667ms | ✓ 1949ms | ✓ 1968ms | 否 | ✓ 1641ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1823ms | ✓ 1754ms | ✓ 1207ms | ✓ 1018ms | http |
| 103.139.138.194:3128 | ✓ 1390ms | 否 | ✓ 1363ms | ✓ 1670ms | ✓ 1262ms | http |
| 160.250.4.13:1 | 否 | 否 | ✓ 1772ms | ✓ 1808ms | ✓ 1486ms | http |
| 47.95.231.180:8084 | ✓ 995ms | ✓ 1556ms | ✓ 1094ms | 否 | 否 | http |
| 211.171.114.154:3128 | ✓ 1673ms | ✓ 1620ms | ✓ 1163ms | ✓ 1710ms | ✓ 1497ms | http |
| 61.72.110.94:3128 | ✓ 1660ms | 否 | ✓ 1022ms | 否 | ✓ 1523ms | http |
| 61.72.221.94:3128 | ✓ 1663ms | 否 | ✓ 976ms | ✓ 1906ms | 否 | http |
| 46.249.103.192:443 | ✓ 969ms | 否 | ✓ 1654ms | ✓ 1982ms | 否 | http |
| 167.172.69.123:8080 | ✓ 851ms | 否 | ✓ 861ms | ✓ 1202ms | ✓ 1113ms | http |
| 151.245.137.203:8085 | ✓ 480ms | 否 | ✓ 1623ms | ✓ 1829ms | ✓ 1716ms | http |
| 14.225.222.164:7890 | ✓ 1888ms | 否 | ✓ 1160ms | ✓ 1182ms | ✓ 958ms | http |
| 103.100.159.145:80 | ✓ 974ms | ✓ 1685ms | ✓ 1321ms | ✓ 1666ms | 否 | http |

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
