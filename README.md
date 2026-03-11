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

最后更新：2026-03-11 15:42:07 UTC（2026-03-11 23:42:07 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.63:8443 | ✓ 1024ms | ✓ 1035ms | ✓ 387ms | ✓ 841ms | ✓ 641ms | http |
| 45.136.130.175:8443 | ✓ 1026ms | ✓ 1377ms | ✓ 202ms | ✓ 832ms | ✓ 628ms | http |
| 45.136.131.47:8443 | ✓ 1026ms | ✓ 793ms | ✓ 765ms | ✓ 931ms | ✓ 1007ms | http |
| 101.47.73.135:3128 | ✓ 1742ms | 否 | ✓ 1945ms | ✓ 1350ms | ✓ 1185ms | http |
| 205.209.118.30:3138 | ✓ 772ms | 否 | ✓ 869ms | ✓ 1182ms | ✓ 921ms | http |
| 211.171.114.154:3128 | ✓ 974ms | ✓ 1703ms | 否 | ✓ 1587ms | ✓ 1550ms | http |
| 94.176.3.43:7443 | ✓ 1173ms | 否 | ✓ 1996ms | 否 | ✓ 1978ms | http |
| 115.231.181.40:8128 | ✓ 1352ms | ✓ 1265ms | 否 | ✓ 1340ms | ✓ 1957ms | http |
| 103.84.95.54:7890 | ✓ 728ms | 否 | ✓ 715ms | ✓ 1385ms | ✓ 847ms | http |
| 59.46.216.131:30001 | ✓ 1193ms | 否 | ✓ 1949ms | 否 | ✓ 1109ms | http |
| 111.48.191.1:7890 | ✓ 803ms | ✓ 1049ms | ✓ 809ms | ✓ 1090ms | ✓ 902ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1326ms | ✓ 1405ms | ✓ 1344ms | ✓ 1265ms | http |
| 39.104.201.40:7890 | ✓ 1013ms | ✓ 1803ms | ✓ 1247ms | 否 | ✓ 1561ms | http |
| 202.155.12.161:443 | ✓ 1944ms | 否 | 否 | ✓ 1719ms | ✓ 1961ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1648ms | ✓ 1669ms | ✓ 1599ms | http |
| 101.43.255.96:80 | ✓ 1221ms | ✓ 1371ms | 否 | ✓ 1471ms | ✓ 1755ms | http |
| 190.9.109.198:999 | ✓ 1064ms | 否 | 否 | ✓ 1790ms | ✓ 1337ms | http |
| 194.213.18.200:443 | ✓ 1301ms | 否 | ✓ 249ms | 否 | ✓ 927ms | http |
| 91.107.141.42:8081 | ✓ 599ms | 否 | ✓ 1429ms | 否 | ✓ 1592ms | http |
| 45.136.130.191:8443 | ✓ 194ms | ✓ 995ms | ✓ 194ms | ✓ 822ms | ✓ 697ms | http |
| 45.136.130.188:8443 | ✓ 735ms | ✓ 828ms | ✓ 624ms | ✓ 799ms | ✓ 631ms | http |
| 45.136.130.223:8443 | ✓ 751ms | ✓ 805ms | ✓ 787ms | ✓ 833ms | ✓ 786ms | http |
| 172.212.68.37:3128 | ✓ 260ms | 否 | ✓ 1002ms | ✓ 1819ms | ✓ 1418ms | http |
| 158.69.185.37:3129 | ✓ 469ms | 否 | ✓ 1654ms | ✓ 1123ms | ✓ 1960ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1470ms | 否 | ✓ 958ms | ✓ 944ms | http |
| 5.252.33.13:2025 | ✓ 1439ms | 否 | ✓ 1279ms | 否 | ✓ 1685ms | http |
| 95.3.9.78:3128 | ✓ 1885ms | 否 | 否 | ✓ 1674ms | ✓ 1380ms | http |
| 120.92.212.16:7890 | ✓ 1050ms | 否 | ✓ 1380ms | ✓ 1311ms | 否 | http |
| 95.3.9.78:8080 | 否 | 否 | ✓ 1954ms | ✓ 1921ms | ✓ 1981ms | http |
| 171.251.172.78:5108 | 否 | 否 | ✓ 1671ms | ✓ 1652ms | ✓ 1507ms | http |
| 171.251.172.78:5110 | 否 | 否 | ✓ 1949ms | ✓ 1618ms | ✓ 1529ms | http |
| 121.29.195.167:7890 | ✓ 1127ms | 否 | ✓ 1008ms | ✓ 1290ms | 否 | http |
| 46.183.25.8:443 | ✓ 1304ms | 否 | ✓ 1310ms | ✓ 1806ms | 否 | http |
| 107.173.0.178:1080 | ✓ 1237ms | 否 | ✓ 918ms | ✓ 1127ms | ✓ 820ms | http |
| 107.172.125.217:3128 | ✓ 895ms | 否 | ✓ 1876ms | ✓ 832ms | ✓ 681ms | http |
| 168.235.110.63:3128 | ✓ 1571ms | ✓ 1614ms | ✓ 1824ms | 否 | ✓ 835ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1549ms | ✓ 1327ms | ✓ 1426ms | http |
| 103.139.138.194:3128 | ✓ 1886ms | 否 | ✓ 1808ms | ✓ 1657ms | ✓ 1679ms | http |
| 152.42.213.210:8080 | ✓ 805ms | 否 | ✓ 1085ms | 否 | ✓ 1265ms | http |
| 81.70.169.194:80 | ✓ 1183ms | 否 | 否 | ✓ 1425ms | ✓ 1915ms | http |
| 178.236.245.59:3128 | ✓ 1286ms | 否 | ✓ 1015ms | ✓ 1755ms | ✓ 1521ms | http |
| 152.42.213.210:443 | ✓ 819ms | 否 | ✓ 1523ms | ✓ 1150ms | 否 | http |
| 178.236.245.17:3128 | 否 | 否 | ✓ 949ms | ✓ 1748ms | ✓ 1324ms | http |
| 45.136.198.40:3128 | ✓ 808ms | ✓ 1694ms | 否 | ✓ 1968ms | ✓ 1635ms | http |
| 190.212.131.238:3128 | ✓ 1110ms | 否 | ✓ 1630ms | 否 | ✓ 1726ms | http |
| 116.80.47.62:3172 | 否 | 否 | ✓ 1743ms | ✓ 1910ms | ✓ 1758ms | http |
| 35.225.22.61:80 | 否 | ✓ 1952ms | ✓ 401ms | ✓ 1325ms | 否 | http |
| 165.227.5.10:8888 | ✓ 372ms | 否 | ✓ 1116ms | ✓ 876ms | ✓ 796ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1525ms | ✓ 1376ms | ✓ 1717ms | ✓ 1055ms | http |
| 107.173.52.58:7890 | ✓ 402ms | 否 | ✓ 980ms | ✓ 1224ms | ✓ 1036ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1715ms | ✓ 1518ms | ✓ 1311ms | ✓ 1066ms | http |
| 47.101.149.27:9010 | ✓ 1437ms | 否 | ✓ 1459ms | ✓ 1368ms | 否 | http |
| 152.70.98.46:8888 | 否 | ✓ 1771ms | ✓ 1834ms | ✓ 1007ms | ✓ 712ms | http |
| 103.39.51.190:8080 | ✓ 1362ms | 否 | ✓ 1288ms | ✓ 1370ms | ✓ 1440ms | http |
| 171.251.172.78:5102 | 否 | 否 | ✓ 1895ms | ✓ 1808ms | ✓ 1522ms | http |
| 103.82.23.118:5221 | ✓ 1493ms | 否 | ✓ 1460ms | 否 | ✓ 1530ms | http |
| 61.52.131.172:8443 | ✓ 950ms | ✓ 1221ms | ✓ 1067ms | 否 | ✓ 983ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 977ms | ✓ 1172ms | ✓ 865ms | http |
| 86.53.183.16:1080 | ✓ 1408ms | 否 | ✓ 1609ms | 否 | ✓ 1748ms | http |
| 171.251.172.78:5105 | 否 | 否 | ✓ 1648ms | ✓ 1710ms | ✓ 1463ms | http |
| 185.191.236.162:3128 | ✓ 657ms | 否 | ✓ 873ms | ✓ 1642ms | ✓ 1368ms | http |

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
