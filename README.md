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

最后更新：2026-03-05 21:44:51 UTC（2026-03-06 05:44:51 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1006ms | ✓ 1743ms | ✓ 583ms | ✓ 1057ms | ✓ 992ms | http |
| 185.191.236.162:3128 | ✓ 487ms | 否 | ✓ 1532ms | ✓ 1503ms | ✓ 1084ms | http |
| 103.84.95.54:7890 | ✓ 1641ms | 否 | ✓ 910ms | 否 | ✓ 1668ms | http |
| 61.72.221.194:3128 | ✓ 1623ms | 否 | ✓ 1686ms | ✓ 1984ms | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1287ms | ✓ 1352ms | ✓ 845ms | http |
| 125.128.12.14:3128 | ✓ 828ms | ✓ 1735ms | ✓ 1207ms | ✓ 1493ms | 否 | http |
| 61.72.221.94:3128 | ✓ 1514ms | 否 | ✓ 815ms | ✓ 1674ms | ✓ 1066ms | http |
| 61.72.110.54:3128 | ✓ 802ms | 否 | ✓ 1025ms | ✓ 1812ms | ✓ 1291ms | http |
| 125.128.12.144:3128 | ✓ 1649ms | ✓ 1694ms | 否 | 否 | ✓ 1471ms | http |
| 45.140.147.82:1081 | ✓ 443ms | ✓ 1009ms | ✓ 384ms | 否 | ✓ 1147ms | http |
| 115.231.181.40:8128 | ✓ 1029ms | ✓ 1269ms | ✓ 1901ms | ✓ 1656ms | ✓ 1037ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1662ms | ✓ 1255ms | ✓ 1548ms | 否 | http |
| 5.75.196.26:40000 | ✓ 1002ms | ✓ 1973ms | ✓ 1930ms | ✓ 1941ms | ✓ 1692ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1778ms | ✓ 1103ms | ✓ 1320ms | ✓ 1307ms | http |
| 69.48.179.20:3128 | ✓ 119ms | 否 | ✓ 125ms | ✓ 950ms | 否 | http |
| 61.72.110.94:3128 | ✓ 807ms | 否 | ✓ 1332ms | 否 | ✓ 1342ms | http |
| 188.132.141.249:443 | ✓ 990ms | 否 | ✓ 1346ms | 否 | ✓ 1947ms | http |
| 121.128.121.54:3128 | ✓ 861ms | ✓ 1229ms | ✓ 851ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1085ms | ✓ 1515ms | ✓ 1127ms | ✓ 1430ms | ✓ 1130ms | http |
| 46.249.103.192:443 | ✓ 1089ms | 否 | ✓ 1362ms | ✓ 1757ms | 否 | http |
| 107.174.80.186:3128 | ✓ 418ms | ✓ 1702ms | ✓ 655ms | ✓ 1221ms | 否 | http |
| 180.127.149.244:1080 | ✓ 1039ms | ✓ 1459ms | ✓ 1045ms | ✓ 1506ms | ✓ 1096ms | http |
| 101.43.255.96:80 | ✓ 1166ms | ✓ 1444ms | ✓ 1275ms | ✓ 1559ms | ✓ 1281ms | http |
| 91.193.240.157:9877 | ✓ 903ms | 否 | ✓ 1620ms | 否 | ✓ 1399ms | http |
| 61.72.221.234:3128 | ✓ 1754ms | 否 | ✓ 1038ms | 否 | ✓ 994ms | http |
| 103.82.23.118:5221 | ✓ 1734ms | 否 | ✓ 1327ms | ✓ 1833ms | ✓ 1445ms | http |
| 122.248.45.54:8080 | ✓ 1994ms | 否 | ✓ 1913ms | ✓ 1822ms | 否 | http |
| 121.126.185.63:25152 | ✓ 1182ms | ✓ 1093ms | ✓ 984ms | ✓ 1184ms | ✓ 949ms | http |
| 154.90.48.209:9090 | ✓ 1029ms | 否 | ✓ 1193ms | 否 | ✓ 1293ms | http |
| 5.252.33.13:2025 | ✓ 1573ms | 否 | ✓ 1164ms | 否 | ✓ 1787ms | http |
| 120.92.212.16:7890 | ✓ 1150ms | ✓ 1710ms | ✓ 1130ms | ✓ 1495ms | 否 | http |
| 103.82.23.118:5207 | ✓ 1426ms | ✓ 1880ms | ✓ 1309ms | ✓ 1726ms | ✓ 1461ms | http |
| 103.82.23.118:5171 | ✓ 1350ms | 否 | ✓ 1365ms | ✓ 1946ms | ✓ 1409ms | http |
| 45.140.147.82:1082 | 否 | ✓ 1243ms | ✓ 526ms | 否 | ✓ 1053ms | http |
| 116.58.162.45:3128 | ✓ 1246ms | ✓ 1988ms | ✓ 1550ms | 否 | 否 | http |
| 103.18.78.250:1111 | ✓ 1801ms | 否 | 否 | ✓ 1603ms | ✓ 1533ms | http |
| 51.81.6.158:3128 | ✓ 835ms | 否 | ✓ 1997ms | ✓ 1834ms | 否 | http |
| 51.79.207.21:8080 | ✓ 1558ms | ✓ 1937ms | ✓ 897ms | 否 | ✓ 996ms | http |
| 175.215.54.252:3040 | ✓ 1809ms | ✓ 1431ms | ✓ 1098ms | 否 | 否 | http |
| 180.76.115.231:3128 | ✓ 1604ms | ✓ 1550ms | ✓ 1130ms | ✓ 1636ms | ✓ 1771ms | http |
| 14.56.177.44:3128 | ✓ 1434ms | 否 | ✓ 1156ms | ✓ 1359ms | ✓ 1080ms | http |
| 165.227.5.10:8888 | ✓ 1614ms | 否 | 否 | ✓ 1145ms | ✓ 769ms | http |
| 38.180.2.107:3128 | ✓ 800ms | ✓ 1919ms | ✓ 1776ms | 否 | ✓ 1733ms | http |
| 61.109.216.213:8080 | ✓ 1403ms | 否 | ✓ 1597ms | ✓ 1682ms | ✓ 1186ms | http |
| 210.223.44.230:3128 | ✓ 749ms | ✓ 1175ms | ✓ 726ms | ✓ 1274ms | ✓ 942ms | http |
| 103.74.192.243:7890 | 否 | 否 | ✓ 999ms | ✓ 1018ms | ✓ 845ms | http |
| 103.215.36.88:17013 | ✓ 1080ms | ✓ 1461ms | ✓ 1268ms | ✓ 1561ms | ✓ 1120ms | http |
| 18.100.143.47:6550 | ✓ 1381ms | 否 | ✓ 1481ms | 否 | ✓ 1898ms | http |
| 168.235.110.63:3128 | ✓ 218ms | ✓ 1904ms | ✓ 857ms | ✓ 947ms | ✓ 747ms | http |
| 120.92.212.16:8890 | ✓ 1135ms | ✓ 1401ms | 否 | ✓ 1491ms | ✓ 1113ms | http |
| 185.191.236.162:8080 | ✓ 525ms | ✓ 1507ms | ✓ 1754ms | ✓ 1961ms | ✓ 1774ms | http |
| 91.107.175.112:10801 | ✓ 402ms | 否 | ✓ 1565ms | ✓ 1713ms | ✓ 1729ms | http |
| 121.230.8.34:1080 | 否 | ✓ 1651ms | ✓ 1253ms | ✓ 1760ms | ✓ 1269ms | http |
| 121.230.9.75:1080 | ✓ 1344ms | ✓ 1740ms | ✓ 1284ms | ✓ 1811ms | ✓ 1856ms | http |
| 103.135.102.161:8080 | 否 | 否 | ✓ 1901ms | ✓ 1753ms | ✓ 1174ms | http |
| 45.136.198.40:3128 | ✓ 1159ms | ✓ 1716ms | ✓ 1959ms | 否 | ✓ 1985ms | http |
| 88.80.150.82:8080 | ✓ 1485ms | ✓ 1899ms | 否 | ✓ 1997ms | ✓ 1688ms | https |
| 172.212.68.37:3128 | ✓ 725ms | ✓ 1355ms | 否 | ✓ 1074ms | ✓ 944ms | http |
| 121.230.8.135:1080 | ✓ 1175ms | ✓ 1430ms | ✓ 1092ms | ✓ 1645ms | ✓ 1158ms | http |
| 103.215.36.88:16474 | ✓ 1114ms | ✓ 1480ms | ✓ 1222ms | ✓ 1583ms | ✓ 1183ms | http |
| 103.215.36.88:19195 | ✓ 1234ms | ✓ 1692ms | ✓ 1308ms | ✓ 1847ms | ✓ 1211ms | http |
| 144.31.69.170:1080 | ✓ 891ms | 否 | ✓ 1256ms | 否 | ✓ 1521ms | http |
| 138.124.53.25:7443 | ✓ 935ms | 否 | ✓ 1619ms | 否 | ✓ 1979ms | http |
| 120.79.99.232:8099 | ✓ 1435ms | ✓ 1727ms | ✓ 1453ms | ✓ 1685ms | ✓ 1415ms | http |
| 192.166.82.55:1080 | ✓ 1898ms | 否 | ✓ 1711ms | ✓ 1532ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1652ms | 否 | ✓ 1493ms | ✓ 1862ms | ✓ 1581ms | http |
| 121.230.8.136:1080 | ✓ 1647ms | 否 | ✓ 1649ms | 否 | ✓ 1235ms | http |
| 201.150.116.32:999 | ✓ 1169ms | ✓ 1880ms | ✓ 1053ms | ✓ 1191ms | ✓ 1021ms | http |
| 103.215.36.88:16316 | ✓ 1376ms | ✓ 1870ms | ✓ 1065ms | 否 | 否 | http |
| 147.161.182.46:10810 | ✓ 1031ms | ✓ 1445ms | ✓ 1882ms | ✓ 1671ms | ✓ 1116ms | http |
| 147.161.182.46:9443 | ✓ 1030ms | ✓ 1443ms | ✓ 1885ms | ✓ 1641ms | ✓ 1347ms | http |
| 45.186.6.104:3128 | ✓ 1806ms | ✓ 1809ms | ✓ 1649ms | 否 | 否 | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1761ms | ✓ 1896ms | ✓ 1919ms | http |
| 103.215.36.88:17100 | ✓ 1148ms | ✓ 1596ms | ✓ 1262ms | ✓ 1582ms | ✓ 1223ms | http |
| 103.215.36.88:16777 | ✓ 1576ms | ✓ 1819ms | ✓ 1374ms | ✓ 1981ms | ✓ 1375ms | http |
| 183.237.195.130:3128 | ✓ 1427ms | ✓ 1526ms | ✓ 1492ms | ✓ 1391ms | ✓ 1113ms | http |
| 212.175.29.184:8080 | ✓ 1238ms | ✓ 1917ms | ✓ 1684ms | 否 | ✓ 1635ms | http |

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
