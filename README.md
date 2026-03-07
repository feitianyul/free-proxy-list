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

最后更新：2026-03-07 19:30:19 UTC（2026-03-08 03:30:19 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 438ms | ✓ 1122ms | ✓ 1167ms | ✓ 1299ms | ✓ 970ms | http |
| 1.231.81.166:3128 | ✓ 1522ms | ✓ 1035ms | ✓ 898ms | ✓ 943ms | ✓ 742ms | http |
| 152.42.213.210:8080 | ✓ 1355ms | 否 | ✓ 1700ms | ✓ 1042ms | ✓ 869ms | http |
| 61.72.221.94:3128 | 否 | 否 | ✓ 1603ms | ✓ 1364ms | ✓ 1023ms | http |
| 178.236.245.17:3128 | ✓ 868ms | 否 | ✓ 1924ms | ✓ 1794ms | ✓ 1781ms | http |
| 61.72.110.94:3128 | ✓ 1534ms | ✓ 1930ms | 否 | 否 | ✓ 1655ms | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 1166ms | ✓ 1009ms | ✓ 798ms | http |
| 61.72.110.54:3128 | ✓ 716ms | ✓ 1008ms | ✓ 895ms | 否 | ✓ 793ms | http |
| 35.225.22.61:80 | ✓ 374ms | ✓ 1461ms | ✓ 1021ms | ✓ 1003ms | ✓ 829ms | http |
| 165.227.5.10:8888 | ✓ 441ms | 否 | ✓ 810ms | 否 | ✓ 593ms | http |
| 85.9.195.140:1080 | 否 | 否 | ✓ 1753ms | ✓ 1824ms | ✓ 1531ms | http |
| 116.80.82.232:3172 | ✓ 1617ms | 否 | 否 | ✓ 1946ms | ✓ 1907ms | http |
| 167.172.69.123:8080 | ✓ 738ms | 否 | ✓ 1686ms | ✓ 1208ms | ✓ 912ms | http |
| 67.169.98.211:443 | ✓ 1113ms | 否 | ✓ 478ms | ✓ 1093ms | 否 | http |
| 167.172.69.123:80 | ✓ 761ms | 否 | ✓ 1218ms | ✓ 1089ms | ✓ 845ms | http |
| 81.70.169.194:80 | ✓ 1001ms | ✓ 1318ms | ✓ 1021ms | ✓ 1330ms | ✓ 1036ms | http |
| 125.128.12.14:3128 | ✓ 1457ms | 否 | ✓ 972ms | ✓ 1048ms | ✓ 1386ms | http |
| 14.56.177.44:3128 | ✓ 1508ms | ✓ 1833ms | ✓ 1072ms | ✓ 1389ms | ✓ 941ms | http |
| 121.128.121.54:3128 | ✓ 1453ms | ✓ 1037ms | ✓ 1107ms | 否 | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1222ms | ✓ 914ms | ✓ 1320ms | 否 | http |
| 14.225.222.164:7890 | ✓ 1965ms | 否 | ✓ 1202ms | ✓ 1532ms | 否 | http |
| 185.115.74.185:8080 | ✓ 1414ms | ✓ 1858ms | ✓ 1625ms | 否 | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1274ms | ✓ 945ms | ✓ 1207ms | ✓ 1101ms | http |
| 121.40.231.103:7890 | ✓ 857ms | 否 | ✓ 1143ms | ✓ 1370ms | ✓ 1891ms | http |
| 103.139.138.194:3128 | ✓ 1769ms | 否 | ✓ 1763ms | ✓ 1404ms | ✓ 1264ms | http |
| 42.115.72.27:2039 | ✓ 1704ms | 否 | ✓ 1519ms | ✓ 1725ms | 否 | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 960ms | ✓ 1309ms | ✓ 1115ms | http |
| 14.225.217.30:7890 | ✓ 1452ms | 否 | 否 | ✓ 1284ms | ✓ 853ms | http |
| 162.248.165.72:1080 | ✓ 770ms | 否 | ✓ 1265ms | 否 | ✓ 1979ms | http |
| 61.72.221.234:3128 | ✓ 1094ms | ✓ 1522ms | ✓ 1536ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1003ms | 否 | 否 | ✓ 1822ms | ✓ 841ms | http |
| 14.56.107.244:3128 | ✓ 646ms | ✓ 1461ms | ✓ 588ms | ✓ 1032ms | ✓ 804ms | http |
| 159.223.42.219:3128 | ✓ 768ms | 否 | ✓ 1339ms | ✓ 1072ms | ✓ 948ms | http |
| 115.231.181.40:8128 | ✓ 953ms | ✓ 1096ms | ✓ 1004ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1187ms | ✓ 1214ms | 否 | ✓ 1457ms | ✓ 965ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1775ms | ✓ 1854ms | ✓ 1720ms | http |
| 193.168.173.136:443 | ✓ 663ms | ✓ 1753ms | ✓ 679ms | 否 | 否 | http |
| 168.235.110.63:3128 | ✓ 1625ms | ✓ 1384ms | ✓ 1247ms | ✓ 1146ms | ✓ 1254ms | http |
| 194.59.204.87:9080 | ✓ 705ms | ✓ 1756ms | ✓ 1523ms | 否 | 否 | http |
| 103.82.23.118:5185 | ✓ 1777ms | 否 | ✓ 1603ms | 否 | ✓ 1426ms | http |
| 121.230.9.64:1080 | ✓ 1207ms | ✓ 1415ms | ✓ 1287ms | ✓ 1634ms | ✓ 1239ms | http |
| 91.193.240.157:9877 | ✓ 1497ms | 否 | ✓ 1029ms | 否 | ✓ 1535ms | http |
| 106.14.203.63:3333 | ✓ 896ms | ✓ 1970ms | 否 | ✓ 1928ms | 否 | http |
| 180.125.216.109:8118 | ✓ 1961ms | ✓ 1204ms | ✓ 883ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 989ms | 否 | ✓ 1064ms | ✓ 1454ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1623ms | 否 | ✓ 1883ms | ✓ 1762ms | ✓ 1868ms | http |
| 173.212.222.244:8888 | ✓ 1095ms | ✓ 1696ms | ✓ 999ms | 否 | ✓ 1274ms | http |
| 103.236.89.228:7890 | ✓ 1056ms | ✓ 1335ms | ✓ 993ms | ✓ 1759ms | 否 | http |
| 129.213.162.27:17777 | ✓ 837ms | ✓ 1597ms | 否 | ✓ 1828ms | ✓ 1378ms | http |
| 103.84.95.54:7890 | ✓ 1309ms | 否 | ✓ 882ms | ✓ 1992ms | 否 | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1513ms | ✓ 1886ms | ✓ 1562ms | http |
| 103.215.36.88:19475 | ✓ 903ms | ✓ 1245ms | ✓ 982ms | ✓ 1255ms | ✓ 995ms | http |
| 103.215.36.88:13763 | 否 | ✓ 1256ms | 否 | ✓ 1346ms | ✓ 1102ms | http |
| 113.177.131.2:3128 | ✓ 1441ms | ✓ 1610ms | ✓ 1059ms | ✓ 1586ms | ✓ 1062ms | http |
| 202.155.12.161:443 | ✓ 1571ms | ✓ 1293ms | ✓ 1070ms | ✓ 977ms | 否 | http |
| 178.236.245.59:3128 | 否 | 否 | ✓ 820ms | ✓ 1971ms | ✓ 1381ms | http |
| 42.115.72.27:2038 | ✓ 1492ms | 否 | ✓ 1677ms | ✓ 1694ms | ✓ 1650ms | http |
| 45.136.198.40:3128 | ✓ 1924ms | 否 | ✓ 1430ms | 否 | ✓ 1847ms | http |
| 42.115.72.27:2049 | 否 | 否 | ✓ 1517ms | ✓ 1681ms | ✓ 1480ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1539ms | 否 | ✓ 1820ms | ✓ 1102ms | http |
| 193.228.139.78:8888 | ✓ 770ms | ✓ 1858ms | ✓ 1986ms | 否 | 否 | http |
| 103.215.36.88:17421 | ✓ 1027ms | ✓ 1302ms | ✓ 972ms | ✓ 1377ms | 否 | http |
| 39.104.201.40:7890 | ✓ 976ms | ✓ 1233ms | ✓ 1992ms | 否 | ✓ 1012ms | http |
| 192.166.82.55:1080 | ✓ 1245ms | 否 | ✓ 1719ms | ✓ 1627ms | 否 | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1315ms | ✓ 1438ms | ✓ 1385ms | http |
| 62.113.119.14:8080 | ✓ 762ms | 否 | ✓ 1130ms | ✓ 1555ms | ✓ 1239ms | http |
| 42.115.72.27:2102 | 否 | 否 | ✓ 1611ms | ✓ 1725ms | ✓ 1501ms | http |
| 88.80.150.82:8080 | ✓ 922ms | ✓ 1793ms | ✓ 876ms | ✓ 1793ms | ✓ 1489ms | https |
| 103.215.36.88:18147 | ✓ 1318ms | ✓ 1642ms | ✓ 1310ms | ✓ 1389ms | ✓ 1157ms | http |
| 103.190.108.97:3128 | ✓ 1573ms | 否 | ✓ 1393ms | ✓ 1486ms | 否 | http |

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
