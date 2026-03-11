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

最后更新：2026-03-11 06:49:16 UTC（2026-03-11 14:49:16 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 860ms | ✓ 1047ms | ✓ 735ms | ✓ 884ms | ✓ 700ms | http |
| 45.136.130.175:8443 | ✓ 860ms | ✓ 1880ms | ✓ 737ms | ✓ 886ms | ✓ 693ms | http |
| 45.136.131.47:8443 | ✓ 867ms | ✓ 824ms | ✓ 951ms | ✓ 1669ms | ✓ 850ms | http |
| 35.225.22.61:80 | ✓ 852ms | 否 | ✓ 1077ms | 否 | ✓ 1163ms | http |
| 1.231.81.166:3128 | ✓ 1828ms | ✓ 1550ms | ✓ 1068ms | ✓ 1178ms | ✓ 894ms | http |
| 160.238.65.4:3128 | ✓ 1004ms | ✓ 1719ms | ✓ 1835ms | 否 | 否 | http |
| 14.225.212.37:7890 | 否 | ✓ 1550ms | 否 | ✓ 1460ms | ✓ 1985ms | http |
| 160.238.65.3:3128 | ✓ 1003ms | ✓ 1781ms | ✓ 1775ms | 否 | 否 | http |
| 160.238.65.2:3128 | ✓ 1006ms | ✓ 1873ms | ✓ 1679ms | 否 | 否 | http |
| 47.77.193.180:1080 | ✓ 314ms | ✓ 1963ms | ✓ 436ms | ✓ 899ms | ✓ 650ms | http |
| 120.92.212.16:7890 | ✓ 1059ms | ✓ 1369ms | ✓ 1202ms | 否 | ✓ 1067ms | http |
| 178.236.245.17:3128 | ✓ 647ms | 否 | ✓ 726ms | ✓ 1704ms | ✓ 1316ms | http |
| 178.236.245.59:3128 | ✓ 1614ms | 否 | ✓ 748ms | ✓ 1624ms | ✓ 1321ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1354ms | ✓ 1404ms | ✓ 1313ms | http |
| 158.69.185.37:3129 | ✓ 161ms | ✓ 1225ms | ✓ 436ms | ✓ 993ms | ✓ 748ms | http |
| 43.162.113.116:3128 | ✓ 498ms | ✓ 1726ms | ✓ 399ms | ✓ 830ms | ✓ 651ms | http |
| 45.136.130.191:8443 | 否 | ✓ 798ms | ✓ 421ms | ✓ 845ms | ✓ 668ms | http |
| 46.183.25.8:443 | ✓ 705ms | 否 | ✓ 564ms | ✓ 1158ms | 否 | http |
| 45.136.130.188:8443 | 否 | ✓ 1857ms | ✓ 230ms | ✓ 876ms | ✓ 669ms | http |
| 160.238.65.5:3128 | ✓ 990ms | ✓ 1384ms | ✓ 691ms | ✓ 1526ms | ✓ 1152ms | http |
| 160.238.65.6:3128 | ✓ 990ms | ✓ 1470ms | ✓ 611ms | ✓ 1523ms | ✓ 1160ms | http |
| 160.238.65.9:3128 | ✓ 990ms | ✓ 1399ms | ✓ 682ms | ✓ 1547ms | ✓ 1159ms | http |
| 160.238.65.7:3128 | ✓ 989ms | 否 | ✓ 685ms | ✓ 1475ms | ✓ 1049ms | http |
| 160.238.65.8:3128 | ✓ 990ms | ✓ 1997ms | ✓ 697ms | ✓ 1455ms | ✓ 1065ms | http |
| 95.3.9.78:3128 | ✓ 1027ms | 否 | ✓ 1114ms | ✓ 1619ms | ✓ 1261ms | http |
| 152.42.213.210:8080 | ✓ 1626ms | 否 | ✓ 1222ms | ✓ 1216ms | ✓ 1071ms | http |
| 152.70.98.46:8888 | ✓ 1858ms | 否 | ✓ 1740ms | ✓ 1091ms | ✓ 1298ms | http |
| 91.107.141.42:8081 | ✓ 992ms | 否 | ✓ 1165ms | 否 | ✓ 1521ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 225ms | ✓ 1124ms | ✓ 714ms | http |
| 38.35.247.157:999 | ✓ 1043ms | ✓ 1580ms | ✓ 1171ms | ✓ 1684ms | ✓ 1343ms | http |
| 95.3.9.78:8080 | ✓ 1567ms | 否 | 否 | ✓ 1680ms | ✓ 1294ms | http |
| 45.136.130.223:8443 | ✓ 752ms | ✓ 1716ms | ✓ 251ms | ✓ 971ms | ✓ 664ms | http |
| 202.155.12.161:443 | ✓ 1903ms | 否 | ✓ 1333ms | ✓ 1125ms | 否 | http |
| 42.96.16.158:1311 | ✓ 1752ms | 否 | ✓ 1201ms | ✓ 1378ms | 否 | http |
| 45.136.130.239:8443 | 否 | ✓ 855ms | 否 | ✓ 1036ms | ✓ 1482ms | http |
| 190.9.109.198:999 | ✓ 833ms | ✓ 1399ms | ✓ 1182ms | ✓ 1276ms | ✓ 1243ms | http |
| 194.213.18.200:443 | ✓ 1203ms | 否 | 否 | ✓ 1066ms | ✓ 1735ms | http |
| 101.43.255.96:80 | ✓ 1289ms | 否 | 否 | ✓ 1430ms | ✓ 1093ms | http |
| 5.252.33.13:2025 | ✓ 1441ms | 否 | ✓ 1241ms | 否 | ✓ 1697ms | http |
| 81.70.169.194:80 | ✓ 1168ms | ✓ 1417ms | ✓ 1073ms | ✓ 1420ms | 否 | http |
| 37.139.33.145:1080 | ✓ 1087ms | 否 | ✓ 1308ms | 否 | ✓ 1910ms | http |
| 115.231.181.40:8128 | ✓ 1387ms | ✓ 1276ms | ✓ 1174ms | 否 | 否 | http |
| 39.104.201.40:7890 | ✓ 1739ms | ✓ 1890ms | ✓ 1815ms | ✓ 1368ms | ✓ 1048ms | http |
| 107.172.125.217:3128 | ✓ 363ms | 否 | ✓ 746ms | ✓ 857ms | ✓ 677ms | http |
| 38.180.2.107:3128 | ✓ 823ms | ✓ 1806ms | ✓ 1555ms | 否 | ✓ 1855ms | http |
| 138.124.90.140:1080 | ✓ 1124ms | ✓ 1904ms | ✓ 1530ms | ✓ 1878ms | ✓ 1760ms | http |
| 59.46.216.131:30001 | ✓ 1101ms | ✓ 1564ms | 否 | ✓ 1533ms | 否 | http |
| 205.209.118.30:3138 | ✓ 1980ms | 否 | ✓ 1791ms | 否 | ✓ 1920ms | http |
| 121.138.61.193:8880 | ✓ 1674ms | ✓ 1254ms | ✓ 978ms | ✓ 1129ms | 否 | http |
| 121.230.8.213:1080 | ✓ 1069ms | ✓ 1488ms | ✓ 1113ms | ✓ 1430ms | ✓ 1963ms | http |
| 121.230.8.208:1080 | ✓ 1248ms | ✓ 1836ms | ✓ 1687ms | ✓ 1465ms | ✓ 1296ms | http |
| 45.136.198.40:3128 | ✓ 1357ms | 否 | ✓ 1227ms | 否 | ✓ 1627ms | http |
| 86.53.183.16:1080 | ✓ 722ms | 否 | ✓ 1001ms | 否 | ✓ 1412ms | http |
| 103.191.92.157:1009 | ✓ 1662ms | 否 | ✓ 1327ms | ✓ 1333ms | ✓ 1400ms | http |
| 150.249.255.91:3128 | ✓ 1227ms | 否 | ✓ 962ms | 否 | ✓ 1865ms | http |
| 121.230.8.211:1080 | ✓ 1448ms | ✓ 1390ms | 否 | ✓ 1341ms | ✓ 1154ms | http |
| 103.113.70.189:1081 | ✓ 217ms | ✓ 1270ms | 否 | ✓ 1478ms | ✓ 830ms | http |
| 103.82.23.118:5178 | ✓ 1518ms | 否 | ✓ 1439ms | ✓ 1645ms | ✓ 1358ms | http |
| 162.248.165.72:1080 | ✓ 1121ms | 否 | ✓ 1631ms | 否 | ✓ 1651ms | http |
| 45.140.147.155:1081 | ✓ 880ms | ✓ 1380ms | ✓ 1797ms | ✓ 1698ms | ✓ 1271ms | http |
| 103.39.51.190:8080 | ✓ 1777ms | 否 | 否 | ✓ 1669ms | ✓ 1543ms | http |
| 45.140.147.82:1081 | ✓ 905ms | 否 | ✓ 520ms | ✓ 1628ms | ✓ 1481ms | http |
| 192.71.213.85:9091 | ✓ 832ms | 否 | ✓ 678ms | ✓ 1720ms | 否 | http |
| 218.60.0.214:80 | 否 | 否 | ✓ 1959ms | ✓ 1689ms | ✓ 1301ms | http |
| 190.212.131.238:3128 | ✓ 1313ms | 否 | ✓ 1826ms | 否 | ✓ 1684ms | http |
| 113.177.131.2:3128 | ✓ 1086ms | 否 | ✓ 1105ms | 否 | ✓ 1179ms | http |
| 109.234.38.35:3128 | ✓ 1425ms | 否 | ✓ 1345ms | 否 | ✓ 1426ms | http |
| 121.230.9.198:1080 | ✓ 1351ms | ✓ 1563ms | ✓ 1153ms | 否 | ✓ 1164ms | http |
| 8.140.104.98:3128 | ✓ 980ms | ✓ 1335ms | ✓ 1061ms | ✓ 1407ms | ✓ 1080ms | http |
| 14.225.222.164:7890 | ✓ 1420ms | 否 | ✓ 1414ms | ✓ 1663ms | 否 | http |
| 168.235.110.63:3128 | ✓ 365ms | 否 | ✓ 1447ms | ✓ 1849ms | 否 | http |
| 172.212.68.37:3128 | ✓ 678ms | ✓ 1327ms | ✓ 761ms | ✓ 1720ms | ✓ 1087ms | http |
| 8.219.97.248:80 | ✓ 1427ms | 否 | 否 | ✓ 1947ms | ✓ 1700ms | http |

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
