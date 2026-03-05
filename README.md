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

最后更新：2026-03-05 10:44:38 UTC（2026-03-05 18:44:38 UTC+8）

**代理总数：63**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 271ms | ✓ 948ms | ✓ 1012ms | ✓ 1070ms | ✓ 866ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1718ms | ✓ 1861ms | ✓ 1329ms | http |
| 211.171.114.154:3128 | ✓ 1019ms | ✓ 1246ms | ✓ 1907ms | ✓ 1394ms | ✓ 1536ms | http |
| 121.230.9.108:1080 | ✓ 1642ms | 否 | ✓ 1608ms | ✓ 1607ms | 否 | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1852ms | ✓ 1685ms | ✓ 1496ms | http |
| 35.225.22.61:80 | 否 | ✓ 1913ms | ✓ 1177ms | 否 | ✓ 852ms | http |
| 103.215.36.88:19239 | ✓ 1189ms | ✓ 1623ms | ✓ 1748ms | ✓ 1451ms | ✓ 1138ms | http |
| 103.84.95.54:7890 | ✓ 1663ms | 否 | ✓ 1102ms | 否 | ✓ 1097ms | http |
| 91.193.240.157:9877 | ✓ 1141ms | 否 | ✓ 1268ms | 否 | ✓ 1735ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1295ms | ✓ 1424ms | ✓ 1444ms | 否 | http |
| 81.70.169.194:80 | ✓ 1154ms | 否 | ✓ 1105ms | ✓ 1440ms | ✓ 1973ms | http |
| 62.113.119.14:8080 | ✓ 1030ms | 否 | ✓ 931ms | ✓ 1538ms | ✓ 1801ms | http |
| 101.43.255.96:80 | ✓ 1131ms | ✓ 1834ms | ✓ 1336ms | ✓ 1546ms | ✓ 1538ms | http |
| 116.80.82.224:3172 | ✓ 1891ms | 否 | ✓ 1727ms | 否 | ✓ 1840ms | http |
| 116.80.82.232:3172 | ✓ 1892ms | 否 | ✓ 1728ms | 否 | ✓ 1811ms | http |
| 116.80.82.225:3172 | ✓ 1898ms | 否 | ✓ 1721ms | 否 | ✓ 1817ms | http |
| 116.80.82.227:3172 | ✓ 1891ms | 否 | ✓ 1727ms | 否 | ✓ 1824ms | http |
| 2.56.178.131:443 | ✓ 962ms | 否 | ✓ 925ms | 否 | ✓ 1604ms | http |
| 185.243.218.43:49153 | ✓ 596ms | ✓ 1753ms | ✓ 1506ms | ✓ 1871ms | 否 | http |
| 188.132.141.249:443 | ✓ 687ms | 否 | ✓ 1588ms | 否 | ✓ 1620ms | http |
| 160.250.4.13:1 | ✓ 1072ms | 否 | ✓ 1606ms | ✓ 1648ms | ✓ 1548ms | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1257ms | ✓ 1543ms | ✓ 1599ms | http |
| 220.197.44.36:3128 | ✓ 1386ms | ✓ 1478ms | 否 | 否 | ✓ 1735ms | http |
| 35.234.17.221:8080 | ✓ 1232ms | ✓ 1743ms | 否 | ✓ 1806ms | 否 | http |
| 5.75.196.26:40000 | ✓ 416ms | ✓ 1364ms | 否 | 否 | ✓ 1919ms | http |
| 210.223.44.230:3128 | ✓ 1933ms | ✓ 1113ms | ✓ 1238ms | ✓ 1676ms | 否 | http |
| 103.215.36.88:16894 | 否 | ✓ 1602ms | ✓ 1385ms | 否 | ✓ 1306ms | http |
| 83.219.250.8:62920 | ✓ 661ms | 否 | ✓ 1831ms | ✓ 1866ms | 否 | http |
| 38.180.2.107:3128 | ✓ 950ms | ✓ 1839ms | ✓ 1756ms | 否 | 否 | http |
| 45.10.70.247:8888 | ✓ 772ms | ✓ 1818ms | ✓ 1112ms | ✓ 1130ms | 否 | http |
| 107.173.111.110:7890 | ✓ 1214ms | 否 | ✓ 1952ms | ✓ 1850ms | ✓ 1210ms | http |
| 74.48.78.224:2080 | ✓ 343ms | ✓ 1213ms | ✓ 435ms | ✓ 1044ms | ✓ 1186ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1290ms | ✓ 1810ms | ✓ 1118ms | http |
| 103.215.36.88:18867 | ✓ 1133ms | ✓ 1508ms | 否 | ✓ 1549ms | ✓ 1974ms | http |
| 94.176.3.43:7443 | ✓ 1245ms | 否 | ✓ 1756ms | 否 | ✓ 1885ms | http |
| 121.204.158.249:3128 | ✓ 1159ms | ✓ 1558ms | ✓ 1480ms | 否 | ✓ 1208ms | http |
| 185.40.7.206:3128 | ✓ 1220ms | ✓ 1858ms | ✓ 1703ms | 否 | ✓ 1720ms | http |
| 178.213.25.221:7890 | ✓ 1020ms | 否 | ✓ 1914ms | 否 | ✓ 1664ms | http |
| 180.127.149.244:1080 | ✓ 1120ms | 否 | ✓ 1216ms | ✓ 1507ms | ✓ 1174ms | http |
| 121.230.8.45:1080 | ✓ 1687ms | ✓ 1908ms | ✓ 1468ms | 否 | ✓ 1467ms | http |
| 180.103.19.151:1080 | ✓ 1165ms | 否 | ✓ 1434ms | 否 | ✓ 1232ms | http |
| 46.249.103.192:443 | ✓ 795ms | 否 | ✓ 1186ms | ✓ 1822ms | 否 | http |
| 183.237.195.130:3128 | ✓ 1121ms | ✓ 1661ms | ✓ 1970ms | 否 | 否 | http |
| 103.178.86.178:8080 | ✓ 1970ms | 否 | ✓ 1848ms | ✓ 1705ms | ✓ 1700ms | http |
| 61.72.221.94:3128 | 否 | ✓ 1914ms | ✓ 1890ms | 否 | ✓ 1795ms | http |
| 45.136.198.40:3128 | ✓ 841ms | 否 | ✓ 702ms | 否 | ✓ 1543ms | http |
| 88.80.150.82:8080 | ✓ 1207ms | ✓ 1921ms | 否 | 否 | ✓ 1907ms | https |
| 121.230.9.148:1080 | ✓ 1236ms | 否 | ✓ 1407ms | ✓ 1638ms | ✓ 1139ms | http |
| 192.166.82.55:1080 | ✓ 596ms | 否 | ✓ 1680ms | ✓ 1115ms | ✓ 1397ms | http |
| 114.214.162.73:26001 | ✓ 1341ms | 否 | ✓ 1667ms | ✓ 1702ms | 否 | http |
| 181.78.79.155:999 | ✓ 1080ms | ✓ 1891ms | ✓ 1529ms | 否 | 否 | http |
| 103.215.36.88:16345 | ✓ 1161ms | ✓ 1426ms | ✓ 1154ms | 否 | ✓ 1173ms | http |
| 103.215.36.88:16198 | ✓ 1206ms | ✓ 1504ms | 否 | ✓ 1589ms | ✓ 1551ms | http |
| 89.185.85.138:1080 | ✓ 1105ms | 否 | ✓ 1744ms | ✓ 1921ms | ✓ 1403ms | http |
| 45.140.147.82:1082 | ✓ 1401ms | 否 | ✓ 939ms | ✓ 1571ms | ✓ 1051ms | http |
| 45.140.147.82:1081 | ✓ 530ms | ✓ 1984ms | ✓ 1798ms | ✓ 1585ms | ✓ 1990ms | http |
| 45.140.147.155:1081 | ✓ 802ms | ✓ 1954ms | ✓ 1188ms | ✓ 1219ms | ✓ 1035ms | http |
| 212.175.29.184:8080 | ✓ 919ms | 否 | ✓ 1802ms | 否 | ✓ 1695ms | http |
| 64.181.240.152:3128 | ✓ 401ms | ✓ 1530ms | ✓ 704ms | ✓ 922ms | ✓ 812ms | http |
| 121.230.8.136:1080 | ✓ 1335ms | 否 | ✓ 1416ms | 否 | ✓ 1382ms | http |
| 90.84.188.97:8000 | 否 | 否 | ✓ 965ms | ✓ 1745ms | ✓ 1568ms | http |
| 39.104.201.40:7890 | ✓ 1138ms | ✓ 1456ms | ✓ 1369ms | ✓ 1443ms | ✓ 1149ms | http |
| 160.238.65.7:3128 | ✓ 772ms | 否 | ✓ 1061ms | 否 | ✓ 1653ms | http |

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
