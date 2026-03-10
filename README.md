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

最后更新：2026-03-10 21:41:19 UTC（2026-03-11 05:41:19 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.130.188:8443 | ✓ 647ms | ✓ 974ms | ✓ 182ms | ✓ 753ms | ✓ 554ms | http |
| 45.136.131.47:8443 | ✓ 647ms | ✓ 850ms | ✓ 307ms | ✓ 793ms | ✓ 568ms | http |
| 35.225.22.61:80 | ✓ 677ms | 否 | ✓ 871ms | 否 | ✓ 1053ms | http |
| 45.136.131.63:8443 | ✓ 646ms | ✓ 937ms | ✓ 777ms | ✓ 822ms | ✓ 591ms | http |
| 45.136.130.175:8443 | ✓ 647ms | ✓ 1044ms | ✓ 776ms | ✓ 887ms | ✓ 717ms | http |
| 205.209.118.30:3138 | ✓ 372ms | ✓ 1098ms | ✓ 854ms | ✓ 1293ms | ✓ 1169ms | http |
| 1.231.81.166:3128 | ✓ 1445ms | ✓ 968ms | ✓ 1281ms | ✓ 949ms | ✓ 737ms | http |
| 162.240.154.26:3128 | ✓ 856ms | 否 | ✓ 1705ms | ✓ 1677ms | ✓ 1108ms | http |
| 178.236.245.59:3128 | ✓ 1311ms | 否 | ✓ 1394ms | 否 | ✓ 1728ms | http |
| 168.235.110.63:3128 | ✓ 1450ms | 否 | ✓ 956ms | ✓ 1421ms | ✓ 1077ms | http |
| 178.236.245.17:3128 | ✓ 730ms | ✓ 1933ms | ✓ 683ms | ✓ 1808ms | ✓ 1432ms | http |
| 14.225.222.164:7890 | ✓ 1378ms | 否 | ✓ 1512ms | ✓ 1536ms | 否 | http |
| 45.136.130.191:8443 | ✓ 657ms | ✓ 710ms | ✓ 134ms | ✓ 758ms | ✓ 562ms | http |
| 152.70.98.46:8888 | ✓ 1084ms | ✓ 1245ms | ✓ 1278ms | ✓ 974ms | ✓ 740ms | http |
| 115.231.181.40:8128 | ✓ 959ms | ✓ 1161ms | ✓ 921ms | 否 | ✓ 865ms | http |
| 210.223.44.230:3128 | ✓ 1360ms | ✓ 1392ms | ✓ 1034ms | 否 | ✓ 705ms | http |
| 113.177.131.2:3128 | ✓ 907ms | ✓ 1770ms | ✓ 1075ms | 否 | ✓ 962ms | http |
| 190.212.131.238:3128 | ✓ 1913ms | 否 | ✓ 1194ms | 否 | ✓ 1742ms | http |
| 91.107.141.42:8081 | ✓ 1147ms | ✓ 1759ms | 否 | 否 | ✓ 1617ms | http |
| 190.9.109.198:999 | ✓ 822ms | ✓ 1587ms | ✓ 1233ms | ✓ 1589ms | ✓ 1270ms | http |
| 165.227.5.10:8888 | ✓ 727ms | ✓ 1174ms | 否 | ✓ 1299ms | ✓ 852ms | http |
| 95.3.9.78:8080 | ✓ 1257ms | ✓ 1819ms | ✓ 979ms | ✓ 1737ms | ✓ 1433ms | http |
| 202.155.12.161:443 | ✓ 1810ms | 否 | ✓ 1911ms | ✓ 1303ms | 否 | http |
| 81.70.169.194:80 | ✓ 970ms | ✓ 1245ms | ✓ 1078ms | ✓ 1294ms | ✓ 970ms | http |
| 101.43.255.96:80 | ✓ 956ms | ✓ 1315ms | ✓ 998ms | ✓ 1205ms | 否 | http |
| 39.104.201.40:7890 | ✓ 1593ms | 否 | ✓ 1205ms | ✓ 1446ms | ✓ 950ms | http |
| 138.124.53.25:7443 | ✓ 1603ms | 否 | ✓ 1902ms | ✓ 1684ms | 否 | http |
| 67.169.98.211:443 | ✓ 803ms | 否 | 否 | ✓ 1560ms | ✓ 1833ms | http |
| 120.92.212.16:8890 | ✓ 1436ms | ✓ 1703ms | 否 | ✓ 1261ms | 否 | http |
| 46.183.25.8:443 | ✓ 1013ms | 否 | ✓ 355ms | ✓ 942ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1635ms | 否 | ✓ 1690ms | 否 | ✓ 1226ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1059ms | 否 | ✓ 1118ms | ✓ 1706ms | http |
| 158.69.185.37:3129 | ✓ 473ms | ✓ 1725ms | ✓ 959ms | ✓ 1081ms | ✓ 837ms | http |
| 89.185.85.138:1080 | ✓ 1551ms | ✓ 1828ms | 否 | ✓ 1879ms | ✓ 1212ms | http |
| 45.136.130.223:8443 | ✓ 957ms | ✓ 1358ms | ✓ 1745ms | ✓ 1605ms | ✓ 1722ms | http |
| 194.213.18.200:443 | ✓ 1682ms | 否 | ✓ 1668ms | 否 | ✓ 1748ms | http |
| 195.158.8.123:3128 | ✓ 1884ms | 否 | ✓ 1871ms | 否 | ✓ 1944ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1422ms | ✓ 1689ms | ✓ 1721ms | 否 | http |
| 103.139.138.194:3128 | ✓ 1891ms | 否 | ✓ 1289ms | ✓ 1565ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1020ms | ✓ 1418ms | ✓ 1160ms | 否 | 否 | http |
| 185.191.236.162:3128 | ✓ 971ms | 否 | ✓ 1375ms | ✓ 1756ms | ✓ 1258ms | http |
| 180.103.19.56:1080 | ✓ 1324ms | ✓ 1525ms | ✓ 1388ms | ✓ 1653ms | ✓ 1277ms | http |
| 5.252.33.13:2025 | ✓ 1667ms | 否 | ✓ 1377ms | 否 | ✓ 1910ms | http |
| 45.136.198.40:3128 | ✓ 770ms | ✓ 1592ms | ✓ 763ms | ✓ 1675ms | ✓ 1262ms | http |
| 150.249.255.91:3128 | ✓ 1846ms | ✓ 1384ms | ✓ 1908ms | ✓ 969ms | 否 | http |
| 45.136.130.239:8443 | 否 | ✓ 1683ms | 否 | ✓ 891ms | ✓ 601ms | http |
| 103.35.188.243:3128 | ✓ 1536ms | ✓ 1212ms | 否 | ✓ 1416ms | ✓ 1076ms | http |
| 152.42.213.210:8080 | ✓ 1082ms | 否 | ✓ 1517ms | ✓ 1305ms | ✓ 948ms | http |
| 95.3.9.78:3128 | ✓ 1544ms | 否 | ✓ 1463ms | 否 | ✓ 1454ms | http |
| 103.82.23.118:5242 | ✓ 1708ms | ✓ 1845ms | ✓ 1435ms | 否 | ✓ 1364ms | http |
| 111.79.111.126:3128 | ✓ 1725ms | 否 | ✓ 1731ms | ✓ 1674ms | ✓ 1377ms | http |
| 45.140.147.82:1081 | ✓ 1078ms | ✓ 1703ms | ✓ 1482ms | ✓ 1826ms | ✓ 1244ms | http |
| 217.77.102.18:3128 | ✓ 1686ms | 否 | 否 | ✓ 1948ms | ✓ 1345ms | http |
| 103.39.51.190:8080 | ✓ 1791ms | 否 | ✓ 1710ms | ✓ 1399ms | ✓ 1345ms | http |
| 200.174.198.32:8888 | ✓ 891ms | 否 | ✓ 1441ms | 否 | ✓ 1814ms | http |
| 61.52.131.172:8443 | ✓ 886ms | ✓ 1122ms | ✓ 959ms | ✓ 1876ms | ✓ 912ms | http |
| 103.113.70.189:1081 | ✓ 554ms | ✓ 1606ms | 否 | ✓ 1649ms | ✓ 1030ms | http |

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
